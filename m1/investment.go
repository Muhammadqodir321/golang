package main

import (
    "fmt"
    "math"
)

func main() {
    const inflationRate = 2.5
    var investmentAmount float64
    var years float64
    expectedReturnRate := 5.5

    fmt.Print("Investment Amount: ")
    fmt.Scan(&investmentAmount)

    fmt.Print("Expected Return Rate: ")
    fmt.Scan(&expectedReturnRate)

    fmt.Print("Years: ")
    fmt.Scan(&years)

 
	
	
	futureValue := float64(investmentAmount) * math.Pow(1 + expectedReturnRate / 100, float64(years))
	futureRealvalue := futureValue / math.Pow(1 + inflationRate / 100, years)

	formattedFV := fmt.Sprintf(" Future Value: %.1f\n ",futureValue,)
	formattedRFV :=fmt.Sprintf("Future Value(adjusted for inflation)\n: %.1f",futureRealvalue)

	//fmt.Printf(`Future Value: %.1f\n 
	//Future Value(adjusted for inflation): %.1f`,futureValue,futureRealvalue)
	//fmt.Println("Future Value(adjusted for inflation): ",futureRealvalue)

fmt.Print(formattedFV,formattedRFV )



	} 
