package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalnce, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")
	}

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalnce)
		case 2:
			fmt.Println("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. The programm will be end")
				continue
			}

			accountBalnce += depositAmount
			fileops.WriteFloatToFile(accountBalnce, accountBalanceFile)
			fmt.Println("Balance updated! New amount:", accountBalnce)
		case 3:
			fmt.Println("How much would you withdraw?")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amound. Should be more then 0")
				continue
			}

			if withdrawAmount > accountBalnce {
				fmt.Println("You can't withdraw more then you have")
				continue
			}

			accountBalnce -= withdrawAmount
			fileops.WriteFloatToFile(accountBalnce, accountBalanceFile)
			fmt.Println("This sum left on your account:", accountBalnce)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using our application!")
			return
		}
	}

}
