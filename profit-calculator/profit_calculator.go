package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// fmt.Print()
	// fmt.Scan(&expenses)

	// fmt.Print("Tax Rate: ")
	// fmt.Scan(&taxRate)

	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		fmt.Println(err)
		return
		// panic(err)
	}

	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Tax Rate: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	// if err1 != nil || err2 != nil || err3 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }

	ebt, profit, ratio := calclateFacals(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.1f\n", ratio)
	storeResult(ebt, profit, ratio)
}

func storeResult(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func calclateFacals(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value most be positive")
	}

	return userInput, nil
}
