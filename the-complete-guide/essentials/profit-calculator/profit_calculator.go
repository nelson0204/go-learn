package main

import (
	"errors"
	"fmt"
	"os"
)

/*
Validate user input
=> show error message & exit if invalid input is provided
	-> No negative numbers
	-> Not 0
Store calculated results int file
*/

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
	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	storeCalculationToFile(ebt, profit, ratio)
	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func storeCalculationToFile(ebt, profit, ratio float64) {
	calculationText := fmt.Sprintf("ebt:%.2f\t profit:%.2f\t ratio:%.2f", ebt, profit, ratio)
	os.WriteFile("calculationFinalcials", []byte(calculationText), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number.")
	}

	return userInput, nil
}
