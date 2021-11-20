package main

import (
	"fmt"

	"Github.com/SehonKwon/learnGo/accounts"
)

func main() {
	account := accounts.NewAccount("Sehoon")
	account.Deposit(10)
	fmt.Println((account.Balacne()))
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println((account.Balacne()))
	}
}
