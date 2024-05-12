package main

import (
	"fmt"
)

func main() {
	exit := false
	const failureAttemptsLimit = 3
	failureAttempts := 0
	accountBalance := 100.0

	choices := map[int]string{
		1: "Check my balance",
		2: "Withdraw",
		3: "Deposit",
		4: "Exit",
	}

	fmt.Println("Welcome to Go ATM!")

	for !exit {
		fmt.Println("What do you want to do?")
		for i := 1; i <= len(choices); i++ {
			fmt.Printf("%d. %s\n", i, choices[i])
		}
		var selectedChoice int
		fmt.Scan(&selectedChoice)

		switch selectedChoice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Withdraw amount: ")
			var amount float64
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Withdraw amount must be greater than 0.")
				continue
			}
			if amount > accountBalance {
				fmt.Println("You cannot withdraw more than you account balance.")
				continue
			}
			accountBalance -= amount
			fmt.Println("Your balance is now: ", accountBalance)
		case 3:
			fmt.Print("Deposit amount: ")
			var amount float64
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Withdraw amount must be greater than 0.")
				continue
			}
			accountBalance += amount
			fmt.Println("Your balance is now: ", accountBalance)
		case 4:
			fmt.Println("Goodbye.")
			exit = true
		default:
			failureAttempts++
			if failureAttempts < failureAttemptsLimit {
				fmt.Printf("Invalid choise, please try again, %d attempt left.\n", failureAttemptsLimit-failureAttempts)
			} else {
				exit = true
				fmt.Println("You have exceeded the amount of failure attempts, please try again later.")
			}
		}
	}

	fmt.Println("Thanks for choosing GO ATM.")
}
