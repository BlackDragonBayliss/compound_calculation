package main

import (
	"fmt"
)

var costMitigationNeeded float64
var gainRate float64
var investment float64
var costOfLiving float64
var taxBuffer float64

var daysTilEquilibrium int
var daysTilNominal int

type AmountCalculationMatrix struct {
	ListAmountAtMonthInterval   []float64
	AmountOfMonthsInCalculation int
	FinalAmount                 float64
}

func main() {
	investment = 5030.00
	costOfLiving = 750.00
	taxBuffer = .0463
	gainRate = .02
	amountCalculationMatrix := compoundCalculation(investment, gainRate, 5, 12)
	// compoundCalculation
	for i, v := range amountCalculationMatrix.ListAmountAtMonthInterval {
		fmt.Println(v)
		i++
	}
	fmt.Scanln()
}

func compoundCalculation(startingValue float64, rate float64, amountGainPeriods int, amountmonths int) AmountCalculationMatrix {
	updatedValue := 0.0
	indexAmountMonths := 0
	indexAmountGainPeriods := 0

	amountCalculationMatrix := AmountCalculationMatrix{}
	for indexAmountMonths < amountmonths {
		indexAmountGainPeriods = 0
		if indexAmountMonths == 0 {
			updatedValue = startingValue
			amountCalculationMatrix.ListAmountAtMonthInterval = append(amountCalculationMatrix.ListAmountAtMonthInterval, startingValue)
			amountCalculationMatrix.AmountOfMonthsInCalculation = amountmonths
		}
		for indexAmountGainPeriods < amountGainPeriods {
			updatedValue = updatedValue + (updatedValue * rate)
			indexAmountGainPeriods++
		}
		amountCalculationMatrix.ListAmountAtMonthInterval = append(amountCalculationMatrix.ListAmountAtMonthInterval, updatedValue)
		indexAmountMonths++
		if indexAmountMonths == amountmonths {
			amountCalculationMatrix.FinalAmount = updatedValue
		}
	}
	return amountCalculationMatrix
}
