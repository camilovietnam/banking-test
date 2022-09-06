package main

// Accountable provides a simple interface for keeping track of a bank account.
// TODO: Implement the interface so that it can be used inside http handlers to perform account operations
// TODO: Implement the endpoints provided in the main function
type Accountable interface {
	Deposit(int)
	Withdraw(int)
	Balance() int
	Reset()
}

func main() {
	// TODO: implement the following endpoints. Use the HTTP verbs you consider adequate.
	// /account/reset
	// /account/deposit/{amount}
	// /account/withdraw/{amount}
	// /account/balance
}
