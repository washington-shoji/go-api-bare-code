package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/washington-shoji/gobare/collections/account"
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
	createAccReq := account.CreateAccountRequest{}
	if err := json.NewDecoder(r.Body).Decode(&createAccReq); err != nil {
		return err
	}
	acc, err := account.NewAccount(
		createAccReq.UserName,
		createAccReq.Email,
		createAccReq.FirstName,
		createAccReq.LastName,
		createAccReq.Password,
	)
	if err != nil {
		return err
	}

	if err := account.CreateAccount(acc); err != nil {
		return err
	}

	return helpers.WriteJson(w, http.StatusOK, acc)
}

func HandleGetAccount(w http.ResponseWriter, r *http.Request) error {
	accounts, err := account.GetAccounts()
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
		account, err := account.GetAccountByID(id)
		if err != nil {
			return err
		}
		return helpers.WriteJson(w, http.StatusOK, account)
	}

	if r.Method == "DELETE" {
		acc, err := account.DeleteAccount()
		if err != nil {
			return err
		}

		if err := account.DeleteAccountByID(id, acc); err != nil {
			return err
		}

		res := account.DeleteAccountResponse{
			Message: "Deleted successfully",
		}

		return helpers.WriteJson(w, http.StatusOK, res)
	}

	if r.Method == "PUT" {
		updateAccReq := account.UpdateAccountRequest{}
		if err := json.NewDecoder(r.Body).Decode(&updateAccReq); err != nil {
			return err
		}
		acc, err := account.UpdateAccount(
			updateAccReq.UserName,
			updateAccReq.Email,
			updateAccReq.FirstName,
			updateAccReq.LastName,
			updateAccReq.Password,
		)
		if err != nil {
			return err
		}

		if err := account.UpdateAccountByID(id, acc); err != nil {
			return err
		}

		return helpers.WriteJson(w, http.StatusOK, acc)

	}

	return fmt.Errorf("method not allowed %s", r.Method)
}
