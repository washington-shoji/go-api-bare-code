package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/washington-shoji/gobare/databases"
	"github.com/washington-shoji/gobare/helpers"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

	var req databases.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}

	acc, err := databases.GetAccountByUserName(req.UserName)
	if err != nil {
		return err
	}

	if !helpers.ValidatePassword(acc.EncryptedPassword, req.Password) {
		return fmt.Errorf("not authenticated")
	}

	token, err := helpers.CreateJWT(acc.UserName)
	if err != nil {
		return err
	}

	resp := databases.LoginResponse{
		Token:    token,
		UserName: acc.UserName,
	}

	return helpers.WriteJson(w, http.StatusOK, resp)

}
