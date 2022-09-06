package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/camilovietnam/banking-test/handlers"
	"github.com/camilovietnam/banking-test/types"
	"github.com/go-chi/chi/v5"
)

func main() {

	account := types.NewAccount()
	r := chi.NewRouter()

	r.Get("/account/balance", handlers.Balance(account))
	r.Post("/account/reset", handlers.Reset(account))
	r.Post("/account/deposit/{amount}", handlers.Deposit(account))
	r.Post("/account/withdraw/{amount}", handlers.Withdraw(account))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}

	fmt.Println("[👉] Server started in port 8080 ")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("server could not start: ", err)
	}
}
