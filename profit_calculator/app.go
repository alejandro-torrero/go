package main

import (
	"fmt"
	"strings"
)

func requestValue(name string, pointer *float64) {

	query := "Insert {{name}}: "

	query = strings.Replace(query, "{{name}}", name, 1)

	fmt.Print(query)
	fmt.Scan(pointer)
}

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	requestValue("Revenue", &revenue)
	requestValue("Expenses", &expenses)
	requestValue("taxRate", &taxRate)

	// Make calculations

	earningsBefTax := revenue - expenses

	earningsAftTax := earningsBefTax - (1 - (taxRate / 100))

	ratio := earningsBefTax / earningsAftTax

	fmt.Println("== Result ==")
	fmt.Println("EBT", earningsBefTax)
	fmt.Println("Profit", earningsAftTax)
	fmt.Println("Ratio", ratio)
}
