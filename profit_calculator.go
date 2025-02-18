package main

import (
	"errors"
	"fmt"
	"os"
)

// GOALS OF THE TASK:
// 1) Get revenue, expenses, and tax rate from user
// 2) calculate earnings before tax (EBT)
// 3) calculate earnings after tax (profit)
// 4) calculate ratio (EBT / profit)
// 5) output EBT - PROFIT - RATIO to user

// EXTRA BONUS:
// 1) Validate the user input:
//    ==> Show error message & exit if invalid input is provided
//    - No negative numbers
//    - Not 0
// 2) Store calculated results into file

func main() {
	fmt.Println("----- Welcome to Our Profit Calculator ------")

	revenue, revenueError := getValueFromUser("Revenue")
	if revenueError != nil {
		fmt.Println(revenueError)
		panic("Sorry, something happened and the process went wrong!")
	}

	expenses, expensesError := getValueFromUser("Expenses")
	if expensesError != nil {
		fmt.Println(expensesError)
		panic("Sorry, something happened and the process went wrong!")
	}

	taxRate, taxRateError := getValueFromUser("Tax Rate")
	if taxRateError != nil {
		fmt.Println(taxRateError)
		panic("Sorry, something happened and the process went wrong!")
	}

	// calculate values
	ebt := calculateEBT(revenue, expenses)
	profitValue := calculateProfit(ebt, taxRate)
	ratio := calculateRatio(ebt, profitValue)

	// print the results to the user
	fmt.Printf(`EBT value is: %.2f, Profit value is: %.2f Ratio value is: %.2f`, ebt, profitValue, ratio)

	// store data
	writeToFile(ebt, profitValue, ratio)
}

func getValueFromUser(label string) (float64, error) {
	var inputValue float64

	fmt.Printf(`Please provide the %v: `, label)
	fmt.Scan(&inputValue)

	// do input validations
	if inputValue <= 0 {
		return 0, errors.New("Invalid input! Please provide a positive number.")
	}

	return inputValue, nil
}

// formula: revenue - expenses
func calculateEBT(revenue float64, expenses float64) float64 {
	var result = revenue - expenses
	return result
}

// formula: ebt * ( 1 - taxRate / 100)
func calculateProfit(ebtValue float64, taxRate float64) float64 {
	var result = ebtValue * (1 - taxRate/100)
	return result
}

func calculateRatio(ebtValue float64, profitValue float64) float64 {
	var result = ebtValue / profitValue
	return result
}

func writeToFile(ebtValue, profitValue, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebtValue, profitValue, ratio)
	os.WriteFile(`result.txt`, []byte(results), 0644)
}
