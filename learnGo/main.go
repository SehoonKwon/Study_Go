package main

import (
	"fmt"

	"Github.com/SehonKwon/learnGo/accounts"
)

func main() {
	account := accounts.NewAccount("Sehoon")
	fmt.Println(account)
}
