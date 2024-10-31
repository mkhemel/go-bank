package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

//var balance float64 = 2000

func main() {
	fmt.Println("Welcome to Go Bank!")
	starter()
}

func writeBalanceToFiles(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func getBalanceFormFile() (float64, error) {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		return 1000, errors.New("failed to read Balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}
	return balance, nil
}

func starter() {
	fmt.Println("what do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	getUserInput()
}

func getUserInput() {
	var choice int
	fmt.Print("Your Choice : ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		checkBalance()
	case 2:
		deposit()
	case 3:
		withdraw()
	case 4:
		fmt.Println("Thank you for using Go Bank!")
		return
	default:
		fmt.Println("Your input is invalid.")
		getUserInput()
	}
}

func checkBalance() {
	balance, err := getBalanceFormFile()
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("--------------")
		return
	}
	fmt.Println("Your Balance is : ", balance)
	starter()
}

func deposit() {
	var depositAmount float64
	balance, err := getBalanceFormFile()
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("--------------")
		return
	}
	fmt.Print("Deposit Amount : ")
	fmt.Scan(&depositAmount)
	fmt.Println("Deposit Success!")
	balance += depositAmount
	fmt.Println("Your new balance is : ", balance)
	writeBalanceToFiles(balance)
	starter()
}

func withdraw() {
	var withdrawAmount float64
	balance, err := getBalanceFormFile()
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("--------------")
		return
	}
	fmt.Print("Withdraw Amount : ")
	fmt.Scan(&withdrawAmount)
	if balance >= withdrawAmount { // Updated to >= to include exact balance
		fmt.Println("Withdraw Success!")
		balance -= withdrawAmount // Directly subtract from balance
		writeBalanceToFiles(balance)
		fmt.Println("Your new balance is : ", balance)
	} else {
		fmt.Println("You don't have sufficient balance")
	}
	starter()
}
