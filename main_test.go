package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"testing"
)

var (
	totalRequests = 100

	baseURL    = "http://localhost:8080"
	resetURL   = fmt.Sprintf("%s/account/reset", baseURL)
	balanceURL = fmt.Sprintf("%s/account/balance", baseURL)
	depositURL = fmt.Sprintf("%s/account/deposit/1", baseURL)
)

func TestOneHundredDeposits(t *testing.T) {
	http.Get(resetURL)

	var wg sync.WaitGroup
	wg.Add(totalRequests)

	for i := 0; i < totalRequests; i++ {
		go func() {
			_, _ = http.Post(depositURL, "application/json", nil)
			wg.Done()
		}()
	}

	wg.Wait()

	res, err := http.Get(balanceURL)
	if err != nil {
		t.Fatal("http error: ", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal("read error: ", err)
	}

	fmt.Printf("Message: %s", body)
}
