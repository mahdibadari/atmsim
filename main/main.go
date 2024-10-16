package main

import (
	"atmsim/monetarymodule"
	"atmsim/textmodule"
	"fmt"
)

func main() {
	monetarymodule.InitialATM("123456", 500)
	textmodule.Greeting()

	isPinGood := false
	for !isPinGood {
		textmodule.InputPin()
		isPinGood = monetarymodule.CheckPIN()
	}

	textmodule.ShowMenu()

	var choose int
	for {
		fmt.Scan(&choose)
		if choose == 5 {
			fmt.Println("Thank You, exiting.....")
			break
		} else {
			switch choose {
			case 1:
				textmodule.WithdrawDisplay()
				var withdrawAmount int
				fmt.Scan(&withdrawAmount)
				monetarymodule.Withdraw(withdrawAmount)
				textmodule.ShowMenu()

			case 2:
				textmodule.StoreDisplay()
				var storeAmount int
				fmt.Scan(&storeAmount)
				monetarymodule.Store(storeAmount)
				textmodule.ShowMenu()

			case 3:
				monetarymodule.CheckBalance()
				textmodule.ShowMenu()

			case 4:
				textmodule.ChangePINDisplay()
				var newPin string
				fmt.Scan(&newPin)
				monetarymodule.ChangePIN(newPin)
				textmodule.ShowMenu()

			default:
				fmt.Println("Re-enter your choice!")
				textmodule.ShowMenu()
			}
		}
	}
}
