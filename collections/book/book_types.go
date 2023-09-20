package book

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
}

type CreateBookRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateBookRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type DeleteBookResponse struct {
	Message string `json:"message"`
}

func NewBook(title string, description string) (*Book, error) {
	id := uuid.New()
	time := time.Now().UTC()
	return &Book{
		ID:          id,
		Title:       title,
		Description: description,
		CreatedAt:   time,
	}, nil
}

func UpdateBook(title string, description string) (*Book, error) {
	time := time.Now().UTC()
	return &Book{
		Title:       title,
		Description: description,
		UpdatedAt:   &time,
	}, nil
}

func DeleteBook() (*Book, error) {
	time := time.Now().UTC()
	return &Book{
		DeletedAt: &time,
	}, nil
}
