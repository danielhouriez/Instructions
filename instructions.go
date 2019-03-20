package main

import (
    "flag"
    "fmt"
    "github.com/leekchan/accounting"
    "github.com/yunabe/easycsv"
    "log"
    "sort"
    "time"
)

var (
    inputDateFormat = "02 Jan 2006"
    outputDateFormat = "Mon 02 Jan 2006"
    instructionsFilename = "sample.csv"
    acc = accounting.Accounting{Symbol: "$", Precision: 2}
)

type DailyAmount struct {
    Date time.Time
    Value float32
}

type EntityRanking struct {
    Entity string
    Amount float32
}

func init() {
    flag.StringVar(&instructionsFilename, "file", instructionsFilename, "Name of the CSV file to be processed, located in the same directory.")
}

func main() {
    flag.Parse()
    incoming, outgoing, err := loadInstructions()

    if err != nil {
        log.Fatalf("Failed to read instructions file: %s", err)
    }

    fmt.Println("")
    fmt.Println("Incoming settled amounts")
    printDailyAmountsReport(GetDailyAmounts(incoming))
    fmt.Println("")
    fmt.Println("Outgoing settled amounts")
    printDailyAmountsReport(GetDailyAmounts(outgoing))
    fmt.Println("")
    fmt.Println("Incoming entity ranking")
    printEntityRankingReport(RankEntities(incoming))
    fmt.Println("")
    fmt.Println("Outgoing entity ranking")
    printEntityRankingReport(RankEntities(outgoing))
    fmt.Println("")
}

// loadInstructions parses instructions contained within the CSV file provided, returning incoming and outgoing ones separately.
func loadInstructions() ([]Instruction, []Instruction, error) {
    var (
        incoming []Instruction
        outgoing []Instruction
        instruction Instruction
    )

    csvDecoders := map[string]interface{}{
        "date": func(date string) (time.Time, error) {
            return time.Parse(inputDateFormat, date)
        },
    }

    options := easycsv.Option{Decoders: csvDecoders}
    r := easycsv.NewReaderFile(instructionsFilename, options)

    for r.Read(&instruction) {
        instruction.processSettlementDate()
        if instruction.Type == "B" {
            outgoing = append(outgoing, instruction)
        } else if instruction.Type == "S" {
            incoming = append(incoming, instruction)
        } else {
            log.Printf("Invalid instruction type %s at line %d", instruction.Type, r.LineNumber())
        }
    }

    return incoming, outgoing, r.Done()
}

// GetDailyAmounts processes a slice of instructions and returns the total amount for each day, ordering them by day
func GetDailyAmounts(instructions []Instruction) (amounts []DailyAmount) {
    // Ensure instructions are sorted by date before we start
    sort.Slice(instructions, func(a, b int) bool {
        return instructions[a].SettlementDate.Before(instructions[b].SettlementDate)
    })

    var amount DailyAmount

    for _, instruction := range instructions {
        if !instruction.SettlementDate.Equal(amount.Date) {
            amount = DailyAmount{
                Date: instruction.SettlementDate,
                Value: instruction.getUsdAmount(),
            }
            amounts = append(amounts, amount)
        } else {
            amounts[len(amounts) - 1].Value += instruction.getUsdAmount()
        }
    }

    return
}

// RankEntities processes a slice of instructions and returns the total amount for each entity, ordering them by the total amount, highest amount first
func RankEntities(instructions []Instruction) (ranking []EntityRanking) {
    entities := make(map[string]*EntityRanking)

    for _, instruction := range instructions {
        if _, ok := entities[instruction.Entity]; !ok {
            entities[instruction.Entity] = &EntityRanking{
                Entity: instruction.Entity,
                Amount: instruction.getUsdAmount(),
            }
        } else {
            entities[instruction.Entity].Amount += instruction.getUsdAmount()
        }
    }

    // Maps are unordered so we need to generate a slice we can sort to generate the ranking
    for _, v := range entities {
        ranking = append(ranking, *v)
    }

    sort.Slice(ranking, func(a, b int) bool {
        return ranking[a].Amount > ranking[b].Amount
    })

    return
}

// printDailyAmountsReport prints a report to the console of daily amounts
func printDailyAmountsReport(amounts []DailyAmount) {
    for _, amount := range amounts {
        fmt.Printf("%s %15v\n", amount.Date.Format(outputDateFormat), acc.FormatMoney(amount.Value))
    }
}

// printEntityRankingReport prints a report to the console of entities' ranking
func printEntityRankingReport(ranking []EntityRanking) {
    for k, rank := range ranking {
        fmt.Printf("%d. %-12s %15v\n", k + 1, rank.Entity, acc.FormatMoney(rank.Amount))
    }
}
