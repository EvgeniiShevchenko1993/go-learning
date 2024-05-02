package main

import (
	"errors"
	"fmt"
	"os"
)

func storeResults(EBT, profit, ratio float64) {
	results := fmt.Sprintf("EBT=%v\nprofit=%v\nratio=%v\n", EBT, profit, ratio)
	os.WriteFile("Results.txt", []byte(results), 0644)
}

func main() {

	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("TaxRate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	EBT, profit, ratio := calculateEbtProfitRatio(revenue, expenses, taxRate)

	storeResults(EBT, profit, ratio)

	outputText(EBT, profit, ratio)

}

func getUserInput(someText string) (float64, error) {
	var userInput float64
	fmt.Print(someText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Must be more then 0")
	}
	return userInput, nil
}

func calculateEbtProfitRatio(revenue, expenses, taxRate float64) (EBT float64, profit float64, ratio float64) {
	EBT = revenue - expenses
	profit = EBT * (1 - taxRate/100)
	ratio = EBT / profit
	return EBT, profit, ratio
}

func outputText(EBT float64, profit float64, ratio float64) {
	fmt.Printf("%.1f\n", EBT)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}
