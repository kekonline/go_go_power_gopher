package main

import (
	"fmt"
	"math"
)

func main () {

const inflationRate = 3.5

// var investment = 1000
var investment float64 = 1000

// var investment, rate, years = 1000, 5.5, 10
// var investment, rate, years float64 = 1000, 5.5, 10
// investment, rate, years := 1000, 5.5, 10

// var rate = 5.5
// short hand declaration and infer the type
rate := 5.5
// var years = 10
var years float64 = 10

fmt.Println("Enter investment amount: ")
fmt.Scan(&investment)


// var futureValue = float64(investment) * math.Pow( 1 + rate / 100, float64(years)) 
var futureValue = investment * math.Pow( 1 + rate / 100, years) 

futureRealValue := futureValue / math.Pow( 1 + inflationRate / 100, years)
	


fmt.Println("futureValue: ", futureValue)
fmt.Println("futureRealValue: ", futureRealValue)


}