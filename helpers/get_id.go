package helpers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetID(r *http.Request) (int, error) {
	idString := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idString)
	}

	return id, nil
}

func GetUuiID(r *http.Request) (uuid.UUID, error) {
	idString := mux.Vars(r)["id"]

	id, err := uuid.Parse(idString)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idString)
	}

	return id, nil
}
