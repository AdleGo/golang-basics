package main

import (
	"fmt"
	"math"
)

func main() {
	const BMIPower float64 = 2
	var userHeight, userWeight float64

	fmt.Println("__Body Index Mass Calculator__")

	userHeight, userWeight = getUserInput()
	BMI := calculateBMI(userWeight, userHeight, BMIPower)
	outputResult(BMI)
}

func outputResult(BMI float64) {
	result := fmt.Sprintf("BMI: %.0f", BMI)
	fmt.Print(result)
}

func calculateBMI(weight, height, power float64) float64 {
	return weight / math.Pow(height/100, power)
}

func getUserInput() (float64, float64) {
	var height, weight float64

	fmt.Print("Enter user height: ")
	fmt.Scan(&height)

	fmt.Print("Enter user weight: ")
	fmt.Scan(&weight)

	return height, weight
}
