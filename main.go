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
	ListAmountAtMonthInterval   []Month
	AmountOfMonthsInCalculation int
	StartingAmount              float64
	FinalAmount                 float64
}
type Month struct {
	AmountTotal  float64
	AmountGained float64
}

type EquilibriumModel struct {
	ListAmountAtMonthInterval  []float64
	EquilibriumValue           float64
	AmountNeededForEquilibrium float64
}

type NominalModel struct {
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

	calculateEquilibriumModel(amountCalculationMatrix)

	// calculateNominalModel()
	fmt.Scanln()
}

func calculateEquilibriumModel(amountCalculationMatrix AmountCalculationMatrix) EquilibriumModel {

	//Given all metrics, at which day will equilibrium be met if any
	equilibriumModel := EquilibriumModel{}

	// costOfLiving
	return equilibriumModel
}

// func calculateMonthGross(amountCalculationMatrix AmountCalculationMatrix){
// 	for
// }

// func calculateNominal() {
// }

func compoundCalculation(startingValue float64, rate float64, amountGainPeriods int, amountmonths int) AmountCalculationMatrix {
	updatedValue := 0.0
	indexAmountMonths := 0
	indexAmountGainPeriods := 0
	amountCalculationMatrix := AmountCalculationMatrix{}
	for indexAmountMonths < amountmonths {
		month := Month{}
		indexAmountGainPeriods = 0
		if indexAmountMonths == 0 {
			updatedValue = startingValue
			amountCalculationMatrix.StartingAmount = startingValue
			amountCalculationMatrix.AmountOfMonthsInCalculation = amountmonths
		}
		for indexAmountGainPeriods < amountGainPeriods {
			updatedValue = updatedValue + (updatedValue * rate)
			indexAmountGainPeriods++
		}
		month.AmountTotal = updatedValue
		amountCalculationMatrix.ListAmountAtMonthInterval = append(amountCalculationMatrix.ListAmountAtMonthInterval, month)
		indexAmountMonths++
		if indexAmountMonths == amountmonths {
			amountCalculationMatrix.FinalAmount = updatedValue
		}
	}
	//calculate amount of difference per month
	previousMonthAmount := 0.0
	for i, v := range amountCalculationMatrix.ListAmountAtMonthInterval {
		if i == 0 {
			previousMonthAmount = amountCalculationMatrix.StartingAmount
		}
		amountDifference := v.AmountTotal - previousMonthAmount
		amountCalculationMatrix.ListAmountAtMonthInterval[i].AmountGained = amountDifference
		previousMonthAmount = v.AmountTotal
	}
	return amountCalculationMatrix
}
