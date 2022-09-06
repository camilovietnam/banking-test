package main

import (
	"github.com/camilovietnam/banking-test/types"
	"sync"
	"testing"
)

var (
	totalRequests = 500
)

func TestOneHundredDeposits(t *testing.T) {
	var (
		account = types.NewAccount()
		wg      sync.WaitGroup
	)

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			account.Deposit(1)
		}()
	}

	wg.Wait()

	balance := account.Balance()

	if balance != totalRequests {
		t.Fatalf("Incorrect balance - Got: %d, Expected: %d", balance, totalRequests)
	}
}
