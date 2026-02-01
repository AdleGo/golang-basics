package main

import "fmt"

func main() {
	transactions := []float64{}
	var total float64

	for {
		var singleTransaction float64
		fmt.Print("Please enter transaction (stop to exit): ")
		fmt.Scan(&singleTransaction)

		if singleTransaction == 0 {
			break
		}

		transactions = append(transactions, singleTransaction)
	}

	fmt.Println("transactions:", transactions)

	for _, val := range transactions {
		total += val
	}

	fmt.Println("Balance:", total)
}
