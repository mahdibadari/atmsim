package textmodule

import (
	"fmt"
)

func Greeting() {
	fmt.Println("Welcome to ATM")
}

func InputPin() {
	fmt.Println("Please Input your card PIN")
}

func ShowMenu() {
	fmt.Println("Please select which activity you want to do")
	fmt.Println("1. Withdraw Money")
	fmt.Println("2. Store Money")
	fmt.Println("3. Check Balance")
	fmt.Println("4. Change PIN")
	fmt.Println("5. Finish ATM usage")
}

func CheckPin(isPINValid bool) bool {
	if isPINValid {
		ShowMenu()
		return true
	} else {
		fmt.Println("PIN is not matched")
		return false
	}
}

func ChangePIN(success bool) {
	if success {
		fmt.Println("PIN has been changed")
	} else {
		fmt.Println("PIN has been used. Please reselect the menu and input another PIN")
	}
}

func Withdraw(success bool, amount int) {
	if !success {
		fmt.Println("Remaining Balance is not enough")
	} else {
		fmt.Println("Money withdrawn. Remaining balance: ", amount)
	}
}

func CheckBalance(remainingBalance int) {
	fmt.Println("Money stored. Remaining balance: ", remainingBalance)
}

func StoreDisplay() {
	fmt.Println("Please input how much money you want to store")
}

func WithdrawDisplay() {
	fmt.Println("Please input how much money you want to withdraw")
}

func ChangePINDisplay() {
	fmt.Println("Please input your new PIN")
}

func Store(success bool, amount int) {
	if !success {
		fmt.Println("Invalid Amount")
	} else {
		fmt.Println("Money stored. Remaining balance: ", amount)
	}
}
