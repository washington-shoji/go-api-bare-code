package helpers

import "net/http"

type APIFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func MakeHTTPHandleFunc(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			//handle error
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
