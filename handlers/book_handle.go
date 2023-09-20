package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/washington-shoji/gobare/collections/book"
	"github.com/washington-shoji/gobare/helpers"
)

func HandleBook(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return HandleGetBook(w, r)
	case "POST":
		return HandleCreateBook(w, r)

	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}

}

func HandleCreateBook(w http.ResponseWriter, r *http.Request) error {
	createBkReq := book.CreateBookRequest{}
	if err := json.NewDecoder(r.Body).Decode(&createBkReq); err != nil {
		return err
	}

	bk, err := book.NewBook(
		createBkReq.Title,
		createBkReq.Description,
	)
	if err != nil {
		return err
	}

	if err := book.CreateBook(bk); err != nil {
		return err
	}
	return helpers.WriteJson(w, http.StatusOK, bk)
}

func HandleGetBook(w http.ResponseWriter, r *http.Request) error {
	bks, err := book.GetBooks()
	if err != nil {
		return err
	}

	return helpers.WriteJson(w, http.StatusOK, bks)
}

func HandleBookByID(w http.ResponseWriter, r *http.Request) error {
	id, err := helpers.GetUuiID(r)
	if err != nil {
		return err
	}

	switch r.Method {
	case "GET":
		bk, err := book.GetBookByID(id)
		if err != nil {
			return err
		}
		return helpers.WriteJson(w, http.StatusOK, bk)
	case "PUT":
		updateBkReq := book.UpdateBookRequest{}
		if err := json.NewDecoder(r.Body).Decode(&updateBkReq); err != nil {
			return err
		}
		bk, err := book.UpdateBook(
			updateBkReq.Title,
			updateBkReq.Description,
		)
		if err != nil {
			return err
		}

		if err := book.UpdateBookByID(id, bk); err != nil {
			return err
		}
		return helpers.WriteJson(w, http.StatusOK, bk)
	case "DELETE":
		bk, err := book.DeleteBook()
		if err != nil {
			return err
		}
		if err := book.DeleteBookByID(id, bk); err != nil {
			return err
		}
		res := book.DeleteBookResponse{
			Message: "Deleted successfully",
		}
		return helpers.WriteJson(w, http.StatusOK, res)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}
