package types

// Accountable provides a simple interface for keeping track of a bank account.
type Accountable interface {
	Balance() int
	Deposit(int)
	Withdraw(int)
	Reset()
}
