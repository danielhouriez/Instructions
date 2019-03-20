package main

import (
    "reflect"
    "testing"
)

func TestDailyAmountsAreCalculatedCorrectly(t *testing.T) {
    instructions := []Instruction{
        buildTestInstruction("", "", 0.50, "", "", "02 Jan 2016", 200, 100.25),
        buildTestInstruction("", "", 0.22, "", "", "07 Jan 2016", 450, 150.5),
    }

    amounts := GetDailyAmounts(instructions)
    expected := []DailyAmount{
        {Date: instructions[0].SettlementDate, Value: 10025},
        {Date: instructions[1].SettlementDate, Value: 14899.5},
    }

    if !reflect.DeepEqual(amounts, expected) {
        t.Errorf("Expected: %v; Got: %v", expected, amounts)
    }
}

func TestDailyAmountsOfTheSameDayAreAccumulatedCorrectly(t *testing.T) {
    instructions := []Instruction{
        buildTestInstruction("", "", 0.50, "", "", "07 Jan 2016", 200, 100.25),
        buildTestInstruction("", "", 0.22, "", "", "07 Jan 2016", 450, 150.5),
    }

    amounts := GetDailyAmounts(instructions)
    expected := []DailyAmount{
        {Date: instructions[0].SettlementDate, Value: 24924.5},
    }

    if !reflect.DeepEqual(amounts, expected) {
        t.Errorf("Expected: %v; Got: %v", expected, amounts)
    }
}

func TestDailyAmountsAreSortedByAscendingDate(t *testing.T) {
    instructions := []Instruction{
        buildTestInstruction("", "", 0.22, "", "", "07 Jan 2016", 450, 150.5),
        buildTestInstruction("", "", 0.50, "", "", "02 Jan 2016", 200, 100.25),
    }

    amounts := GetDailyAmounts(instructions)

    if amounts[0].Date.After(amounts[1].Date) {
        t.Errorf("Daily amounts are not sorted by date.")
    }
}

func TestRankingEntities(t *testing.T) {
    instructions := []Instruction{
        buildTestInstruction("foo", "", 1.0, "", "", "", 100, 1.0),
        buildTestInstruction("bar", "", 1.0, "", "", "", 300, 1.0),
        buildTestInstruction("zzz", "", 1.0, "", "", "", 200, 1.0),
    }

    ranking := RankEntities(instructions)

    if (len(ranking) != 3 ||
        ranking[0].Entity != "bar" ||
        ranking[1].Entity != "zzz" ||
        ranking[2].Entity != "foo") {
        t.Errorf("Expected order: bar,zzz,foo; Got: %v", ranking)
    }
}

func TestRankingEntitiesAccumulatesAmountsCorrectly(t *testing.T) {
    instructions := []Instruction{
        buildTestInstruction("foo", "", 1.0, "", "", "", 100, 1.0),
        buildTestInstruction("foo", "", 1.0, "", "", "", 200, 1.0),
    }

    ranking := RankEntities(instructions)
    expected := float32(300)

    if ranking[0].Amount != expected {
        t.Errorf("Expected: %v; Got: %v", expected, ranking[0].Amount)
    }
}
