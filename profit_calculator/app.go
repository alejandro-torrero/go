package main

import (
	"fmt"
)

func getUserInput(name string, pointer *float64) {
	fmt.Printf("Insert %v: ", name)
	fmt.Scan(pointer)
}

func calcRatios(revenue, expenses, taxRate float64) (earningsBefTax float64, earningsAftTax float64, ratio float64) {
	earningsBefTax = revenue - expenses

	earningsAftTax = earningsBefTax * (1 - (taxRate / 100))

	ratio = earningsBefTax / earningsAftTax
	return
}

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	getUserInput("Revenue", &revenue)
	getUserInput("Expenses", &expenses)
	getUserInput("taxRate", &taxRate)

	// Make calculations
	earningsBefTax, earningsAftTax, ratio := calcRatios(revenue, expenses, taxRate)

	fmt.Println("== Result ==")
	fmt.Println("EBT", earningsBefTax)
	fmt.Println("Profit", earningsAftTax)
	fmt.Printf("Ratio %.2f \n", ratio)
}
