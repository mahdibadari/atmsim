package monetarymodule

import (
	"atmsim/textmodule"
	"fmt"
)

type atm struct {
	id      int
	pin     string
	balance int
}

var defaultATM atm

func InitialATM(pin string, balance int) {
	defaultATM = atm{1, pin, balance}
}

func CheckPIN() bool {
	var pin string
	fmt.Scan(&pin)
	isPinValid := defaultATM.pin == pin
	return textmodule.CheckPin(isPinValid)
}

func ChangePIN(newPIN string) {
	success := defaultATM.pin != newPIN
	if success {
		defaultATM.pin = newPIN
	}
	textmodule.ChangePIN(success)
}

func Withdraw(amount int) {
	var remainingBalance = defaultATM.balance
	success := false
	if amount <= remainingBalance {
		remainingBalance -= amount
		defaultATM.balance = remainingBalance
		success = true
	}
	textmodule.Withdraw(success, remainingBalance)
}

func Store(amount int) {
	var remainingBalance = defaultATM.balance
	success := false
	if amount > 0 {
		remainingBalance += amount
		defaultATM.balance = remainingBalance
		success = true
	}
	textmodule.Store(success, remainingBalance)
}

func CheckBalance() {
	textmodule.CheckBalance(defaultATM.balance)
}
