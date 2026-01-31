package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	var wantContinue string

	for {
		const BMIPower float64 = 2
		var userHeight, userWeight float64

		fmt.Println("__Body Index Mass Calculator__")

		userHeight, userWeight = getUserInput()
		BMI, err := calculateBMI(userWeight, userHeight, BMIPower)

		if err != nil {
			panic(err)
		}

		outputResult(BMI)

		fmt.Print("Do you want to continue (yes or no): ")
		fmt.Scan(&wantContinue)

		if wantContinue == "yes" {
			continue
		} else {
			break
		}
	}
}

func outputResult(BMI float64) {
	result := fmt.Sprintf("Your BMI is: %.0f", BMI)

	switch {
	case BMI < 16:
		fmt.Println("Severe Thinness")
	case BMI < 18.5:
		fmt.Println("Moderate Thinness")
	case BMI < 25:
		fmt.Println("Normal")
	case BMI < 30:
		fmt.Println("Overweight")
	default:
		fmt.Println("Obese")
	}

	fmt.Println(result)
}

func calculateBMI(weight, height, power float64) (float64, error) {
	if weight <= 0 || height <= 0 {
		return 0, errors.New("weight or height is not specified")
	}

	return weight / math.Pow(height/100, power), nil
}

func getUserInput() (float64, float64) {
	var height, weight float64

	fmt.Print("Enter user height: ")
	fmt.Scan(&height)

	fmt.Print("Enter user weight: ")
	fmt.Scan(&weight)

	return height, weight
}
