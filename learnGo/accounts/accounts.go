package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balacne int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balacne: 0}
	return &account
}

// Deposit x amount on your account
// a는 receiver, 규칙은 struct의 첫글자를 딴 소문자
// pointer를 활용해 접근해야 call by ref.(실제 데이터 변경)
func (a *Account) Deposit(amount int) {
	a.balacne += amount
}

// Balance of your account
func (a Account) Balacne() int {
	return a.balacne
}

var errNoMoney = errors.New("Can't withdraw you are poor")

// Withdraw from your account
func (a *Account) Withdraw(amount int) error {
	if a.balacne < amount {
		return errNoMoney
	}
	a.balacne -= amount
	return nil // nil == null
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

//
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.owner, "'s account. \nHas: ", a.Balacne())
}
