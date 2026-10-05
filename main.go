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
			break
		}

		transactions = append(transactions, newTransaction)
	}

	balance := calcBalance(transactions)

	fmt.Println("Your transactions are:", transactions)
	fmt.Printf("Your balance is %.2f", balance)
}

func checkForTransactions() (float64, error) {
	var newTransaction float64

	fmt.Println("Enter your transaction:")
	fmt.Scan(&newTransaction)

	if newTransaction == 0 {
		return 0, errors.New("Stop")
	}

	return newTransaction, nil
}

func calcBalance(transactions []float64) float64 {
	balance := 0.0

	for _, value := range transactions {
		balance += value
	}

	return balance
}
