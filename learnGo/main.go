package main

import (
	"fmt"

	"Github.com/SehonKwon/learnGo/accounts"
)

func main() {
	account := accounts.NewAccount("Sehoon")
	account.Deposit(10)
	fmt.Println(account.String())
}
