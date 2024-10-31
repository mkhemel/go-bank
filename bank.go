package main

import (
	"example.com/bank/fileops"
	"fmt"
)

//var balance float64 = 2000

var accountBalanceFile = "balance.txt"

func main() {
	fmt.Println("Welcome to Go Bank!")
	starter()
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
	balance, err := fileops.GetFloatFormFile(accountBalanceFile)
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
	balance, err := fileops.GetFloatFormFile(accountBalanceFile)
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
	fileops.WriteFloatToFiles(balance, accountBalanceFile)
	starter()
}

func withdraw() {
	var withdrawAmount float64
	balance, err := fileops.GetFloatFormFile(accountBalanceFile)
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
		fileops.WriteFloatToFiles(balance, accountBalanceFile)
		fmt.Println("Your new balance is : ", balance)
	} else {
		fmt.Println("You don't have sufficient balance")
	}
	starter()
}
