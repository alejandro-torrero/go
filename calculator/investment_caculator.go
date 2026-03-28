package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Insert investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Insert expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Insert years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1+(expectedReturnRate/100)), years)
	futureRealvalue := futureValue / math.Pow(1+(inflationRate/100), years)

	fmt.Println("Future Value: ", futureValue)
	fmt.Println("Future real value: ", futureRealvalue)
}
