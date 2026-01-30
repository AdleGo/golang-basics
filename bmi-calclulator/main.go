package main

import (
	"fmt"
	"math"
)

func main() {
	const BMIPower float64 = 2
	var userHeight, userWeight float64

	fmt.Println("__Body Index Mass Calculator__")

	fmt.Print("Enter user height: ")
	fmt.Scan(&userHeight)

	fmt.Print("Enter user weight: ")
	fmt.Scan(&userWeight)

	BMI := userWeight / math.Pow(userHeight, BMIPower)
	fmt.Printf("BMI: %v", BMI)
}
