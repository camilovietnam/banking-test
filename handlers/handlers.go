package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/camilovietnam/banking-test/types"
	"github.com/go-chi/chi/v5"
)

type HttpResponse struct {
	Message string `json:"message"`
	Balance int    `json:"balance"`
}

func Balance(account types.Accountable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := HttpResponse{
			Message: "This is your current balance",
			Balance: account.Balance(),
		}

		body, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}

func Reset(account types.Accountable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		account.Reset()

		response := HttpResponse{
			Message: "Balance was reset to zero",
			Balance: account.Balance(),
		}

		body, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}

func Deposit(account types.Accountable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		amount, _ := strconv.Atoi(chi.URLParam(r, "amount"))
		account.Deposit(amount)

		response := HttpResponse{
			Message: "Successfully deposited into the account",
			Balance: account.Balance(),
		}

		body, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}

func Withdraw(account types.Accountable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var response HttpResponse
		amount, _ := strconv.Atoi(chi.URLParam(r, "amount"))
		balance := account.Balance()

		if amount > balance {
			response = HttpResponse{
				Message: "Account balance is not enough",
				Balance: balance,
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
		} else {
			account.Withdraw(amount)
			balance = account.Balance()

			response = HttpResponse{
				Message: "Successfully withdrawn from your account",
				Balance: balance,
			}
		}

		body, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}
