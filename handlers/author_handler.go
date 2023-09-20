package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/washington-shoji/gobare/collections/author"
	"github.com/washington-shoji/gobare/helpers"
)

func HandleAuthor(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return handleCreateAuthor(w, r)
	}

	if r.Method == "GET" {
		return HandleGetAuthor(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func handleCreateAuthor(w http.ResponseWriter, r *http.Request) error {
	createAuthorReq := author.CreateAuthorRequest{}
	if err := json.NewDecoder(r.Body).Decode(&createAuthorReq); err != nil {
		return err
	}

	newAuthor, err := author.NewAuthor(
		createAuthorReq.FullName,
	)
	if err != nil {
		return err
	}

	if err := author.CreateAuthor(newAuthor); err != nil {
		return err
	}

	return helpers.WriteJson(w, http.StatusOK, newAuthor)
}

func HandleGetAuthor(w http.ResponseWriter, r *http.Request) error {
	authors, err := author.GetAuthors()
	if err != nil {
		return err
	}

	return helpers.WriteJson(w, http.StatusOK, authors)
}

func HandleAuthorByID(w http.ResponseWriter, r *http.Request) error {
	id, err := helpers.GetUuiID(r)
	if err != nil {
		return err
	}

	if r.Method == "GET" {
		author, err := author.GetAuthorByID(id)
		if err != nil {
			return err
		}
		return helpers.WriteJson(w, http.StatusOK, author)
	}

	if r.Method == "DELETE" {
		aut, err := author.DeleteAuthor()
		if err != nil {
			return err
		}

		if err := author.DeleteAuthorByID(id, aut); err != nil {
			return err
		}

		res := author.DeleteAuthorResponse{
			Message: "Deleted successfully",
		}

		return helpers.WriteJson(w, http.StatusOK, res)
	}

	if r.Method == "PUT" {
		updateAutReq := author.UpdateAuthorRequest{}
		if err := json.NewDecoder(r.Body).Decode(&updateAutReq); err != nil {
			return err
		}

		aut, err := author.UpdateAuthor(
			updateAutReq.FullName,
		)
		if err != nil {
			return err
		}

		if err := author.UpdateAuthorByID(id, aut); err != nil {
			return err
		}

		return helpers.WriteJson(w, http.StatusOK, aut)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}
