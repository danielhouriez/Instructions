package main

import (
    "time"
)

var (
    /**
     * Each number corresponds to a day of the week, from Sunday to Saturday.
     * They represent the number of days we need to add to settlement dates
     * falling on that week day, in case it is not a work day according to
     * the instruction's currency, so settlement dates are changed to be the
     * next available working day instead.
     */
    aedSarDaysDiff = []int{0, 0, 0, 0, 0, 2, 1} // For AED and SAR currencies
    othersDaysDiff = []int{1, 0, 0, 0, 0, 0, 2} // For any other currencies
)

type Instruction struct {
    Entity              string      `name:"Entity"`
    Type                string      `name:"Buy/Sell"`
    AgreedFx            float32     `name:"AgreedFx"`
    Currency            string      `name:"Currency"`
    InstructionDate     time.Time   `name:"InstructionDate" enc:"date"`
    SettlementDate      time.Time   `name:"SettlementDate" enc:"date"`
    Units               uint16      `name:"Units"`
    PricePerUnit        float32     `name:"PricePerUnit"`
}

// processSettlementDate ensures the instruction's settlement date always falls on a work day, setting it to the next available one if needs be
func (i *Instruction) processSettlementDate() {
    var daysDiff []int

    if i.Currency == "AED" || i.Currency == "SAR" {
        daysDiff = aedSarDaysDiff
    } else {
        daysDiff = othersDaysDiff
    }

    weekday := int(i.SettlementDate.Weekday())

    if daysDiff[weekday] != 0 {
        i.SettlementDate = i.SettlementDate.AddDate(0, 0, daysDiff[weekday])
    }
}

// getUsdAmount returns the USD amount of the instruction
func (i *Instruction) getUsdAmount() float32 {
    return i.PricePerUnit * float32(i.Units) * i.AgreedFx
}
