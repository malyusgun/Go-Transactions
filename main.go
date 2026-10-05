package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Transactions")

	transactions := []float64{}

	for {
		newTransaction, err := checkForTransactions()

		if err != nil {
			fmt.Println("Your transactions are:", transactions)
			break
		}

		transactions = append(transactions, newTransaction)
	}
}

func checkForTransactions() (float64, error) {
	var input float64

	fmt.Println("Enter your transaction:")
	fmt.Scan(&input)

	if input == 0 {
		return 0, errors.New("Stop")
	}

	return input, nil
}
