package main

import (
    "fmt"
    "reflect"
    "testing"
    "time"
)

func TestInstructionsAreAlwaysSettledOnAWorkingDay(t *testing.T) {
    var tests = []struct{
        Currency string
        DaysDiff []int // Number of days to add to get to the next working day for each day of the week
    }{
        {"AED", aedSarDaysDiff},
        {"SAR", aedSarDaysDiff},
        {"GBP", othersDaysDiff},
    }

    for _, test := range tests {
        day := time.Date(2016, time.January, 3, 0, 0, 0, 0, time.UTC) // January 3rd 2016 was a Sunday
        end := day.AddDate(0, 0, 7) // A week later

        for day.Before(end) {
            subTestName := fmt.Sprintf("%s:%s", test.Currency, day.Weekday())
            t.Run(subTestName, func(t *testing.T) {
                instruction := buildTestInstruction("", "", 1.0, test.Currency, "", day.Format(inputDateFormat), 1, 1.0)
                instruction.processSettlementDate()
                days := test.DaysDiff[int(day.Weekday())]
                expected := day.AddDate(0, 0, days)

                if !reflect.DeepEqual(instruction.SettlementDate, expected) {
                    t.Errorf("Expected: %s; Got: %s", expected, instruction.SettlementDate)
                }
            })
            day = day.AddDate(0, 0, 1)
        }
    }
}

func TestGettingInstructionsUsdAMount(t *testing.T) {
    instruction := buildTestInstruction("", "", 2.0, "", "", "", 100, 1.5)
    amount := instruction.getUsdAmount()
    expected := float32(300)

    if amount != expected {
        t.Errorf("Expected: %v; Got: %v", expected, amount)
    }
}
