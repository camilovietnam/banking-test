package main

// Accountable provides a simple interface for keeping track of a bank account.
// TODO: Implement the interface so that it can be used inside http handlers to perform account operations
// TODO: Implement the endpoints provided in the main function
type Accountable interface {
	// Deposit adds to the balance of the account
	Deposit(int)
	// Withdraw removes from the balance of the account
	Withdraw(int)
	// Balance returns the current account balance
	Balance() int
	// Reset sets the account balance to zero
	Reset()
}

func main() {
	// TODO: implement the following endpoints. Use the HTTP verbs you consider adequate.
	// /account/reset - Should set the account balance to zero
	// /account/deposit/{amount} - should add the {amount} to the balance of the account
	// /account/withdraw/{amount} - should remove the {amount} from the balance of the account
	// /account/balance - should return the current account balance
}
