package main

import "fmt"

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

	// get the user
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getValueFromUser("Revenue")
	expenses = getValueFromUser("Expenses")
	taxRate = getValueFromUser("Tax Rate")

	// calculate values
	ebt := calculateEBT(revenue, expenses)
	profitValue := calculateProfit(ebt, taxRate)
	ratio := calculateRatio(ebt, profitValue)

	// print the results to the user
	fmt.Printf(`EBT value is: %.2f, Profit value is: %.2f Ratio value is: %.2f`, ebt, profitValue, ratio)
}

func getValueFromUser(label string) float64 {
	var inputValue float64

	fmt.Printf(`Please provide the %v: `, label)
	fmt.Scan(&inputValue)

	return inputValue
}

// formula: revenue - expenses
func calculateEBT(revenue float64, expenses float64) float64 {
	var result = revenue - expenses
	return result
}

// formula: ebt * ( 1 - taxRate / 100)
func calculateProfit(ebtValue float64, taxRate float64) float64 {
	var result = ebtValue * (1 - ebtValue/100)
	return result
}

func calculateRatio(ebtValue float64, profitValue float64) float64 {
	var result = ebtValue / profitValue
	return result
}
