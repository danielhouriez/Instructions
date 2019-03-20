package main

import (
    "time"
)

// buildTestInstruction builds an Instruction object for testing purposes and returns it
func buildTestInstruction(properties ...interface{}) Instruction {
    instructionDate, _ := time.Parse(inputDateFormat, properties[4].(string))
    settlementDate, _ := time.Parse(inputDateFormat, properties[5].(string))

    return Instruction{
        Entity: properties[0].(string),
        Type: properties[1].(string),
        AgreedFx: float32(properties[2].(float64)),
        Currency: properties[3].(string),
        InstructionDate: instructionDate,
        SettlementDate: settlementDate,
        Units: uint16(properties[6].(int)),
        PricePerUnit: float32(properties[7].(float64)),
    }
}
