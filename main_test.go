package main

import (
	"encoding/json"
	"fmt"
	"github.com/camilovietnam/banking-test/handlers"
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
	res, err := http.Post(resetURL, "application/json", nil)

	var wg sync.WaitGroup
	wg.Add(totalRequests)

	for i := 0; i < totalRequests; i++ {
		go func() {
			_, _ = http.Post(depositURL, "application/json", nil)
			wg.Done()
		}()
	}

	wg.Wait()

	res, err = http.Get(balanceURL)
	if err != nil {
		t.Fatal("http error: ", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal("read error: ", err)
	}

	var response handlers.HttpResponse
	if err = json.Unmarshal(body, &response); err != nil {
		t.Fatal("could not unmarshal response", err)
	}

	if response.Balance != totalRequests {
		t.Fatalf("Incorrect balance - Got: %d, Expected: %d", response.Balance, totalRequests)
	}
}
