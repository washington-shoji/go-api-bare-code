package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/washington-shoji/gobare/databases"
	"github.com/washington-shoji/gobare/helpers"
)

func HandleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return HandleCreateAccount(w, r)
	}

	if r.Method == "GET" {
		return HandleGetAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func HandleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccReq := databases.CreateAccountRequest{}
	if err := json.NewDecoder(r.Body).Decode(&createAccReq); err != nil {
		return err
	}
	account, err := databases.NewAccount(
		createAccReq.UserName,
		createAccReq.Email,
		createAccReq.FirstName,
		createAccReq.LastName,
		createAccReq.Password,
	)
	if err != nil {
		return err
	}

	if err := databases.CreateAccount(account); err != nil {
		return err
	}

	return helpers.WriteJson(w, http.StatusOK, account)
}

func HandleGetAccount(w http.ResponseWriter, r *http.Request) error {
	accounts, err := databases.GetAccounts()
	if err != nil {
		return err
	}

	return helpers.WriteJson(w, http.StatusOK, accounts)
}

func HandleAccountByID(w http.ResponseWriter, r *http.Request) error {
	id, err := helpers.GetID(r)
	if err != nil {
		return err
	}

	if r.Method == "GET" {
		account, err := databases.GetAccountByID(id)
		if err != nil {
			return err
		}
		return helpers.WriteJson(w, http.StatusOK, account)
	}

	if r.Method == "DELETE" {
		account, err := databases.DeleteAccount()
		if err != nil {
			return err
		}

		if err := databases.DeleteAccountByID(id, account); err != nil {
			return err
		}

		res := databases.DeleteAccountResponse{
			Message: "Deleted successfully",
		}

		return helpers.WriteJson(w, http.StatusOK, res)
	}

	if r.Method == "PUT" {
		updateAccReq := databases.UpdateAccountRequest{}
		if err := json.NewDecoder(r.Body).Decode(&updateAccReq); err != nil {
			return err
		}
		account, err := databases.UpdateAccount(
			updateAccReq.UserName,
			updateAccReq.Email,
			updateAccReq.FirstName,
			updateAccReq.LastName,
			updateAccReq.Password,
		)
		if err != nil {
			return err
		}

		if err := databases.UpdateAccountByID(id, account); err != nil {
			return err
		}

		return helpers.WriteJson(w, http.StatusOK, account)

	}

	return fmt.Errorf("method not allowed %s", r.Method)
}
