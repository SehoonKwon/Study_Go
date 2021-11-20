package accounts

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
