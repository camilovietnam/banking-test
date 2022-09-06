package types

type Account struct {
	balance int
}

func (a Account) Deposit(amount int) {
	a.balance += amount
}

func (a Account) Withdraw(amount int) {
	a.balance -= amount
	if a.balance < 0 {
		a.balance = 0
	}
}

func (a Account) Balance() int {
	balance := a.balance

	return balance
}

func (a Account) Reset() {
	a.balance = 0
}

func NewAccount() Account {
	return Account{
		balance: 0,
	}
}
