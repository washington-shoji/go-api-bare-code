package handlers

import (
	"fmt"
	"net/http"

	"github.com/washington-shoji/gobare/helpers"
)

type HealthCheck struct {
	Message string `json:"message"`
}

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

	res := HealthCheck{
		Message: "Server is up and running",
	}

	return helpers.WriteJson(w, http.StatusOK, res)
}
