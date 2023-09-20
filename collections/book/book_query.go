package book

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/washington-shoji/gobare/databases"
)

func CreateBook(bk *Book) error {
	query := `
	INSERT INTO book (id, title, description, created_at)
	VALUES ($1, $2, $3, $4)
	`

	_, err := databases.DB.Database.Query(
		query,
		bk.ID,
		bk.Title,
		bk.Description,
		bk.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func GetBooks() ([]*Book, error) {
	rows, err := databases.DB.Database.Query(`SELECT * FROM book WHERE deleted_at IS NULL`)
	if err != nil {
		return nil, err
	}
	bks := []*Book{}
	for rows.Next() {
		bk, err := scanIntoBook(rows)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}

	return bks, nil
}

func GetBookByID(id uuid.UUID) (*Book, error) {
	rows, err := databases.DB.Database.Query(`SELECT * FROM book WHERE id = $1 AND deleted_at IS NULL`, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoBook(rows)
	}

	return nil, fmt.Errorf("book %d not found", id)
}

func GetBookByTitle(title string) (*Book, error) {
	rows, err := databases.DB.Database.Query(`SELECT * FROM book WHERE title = $1 AND deleted_at IS NULL`, title)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoBook(rows)
	}

	return nil, fmt.Errorf("book with title [%s] not found", title)
}

func UpdateBookByID(id uuid.UUID, bk *Book) error {
	query := `UPDATE book
	SET 
	title = $2,
	description = $3,
	updated_at = $4
	WHERE id = $1
	`

	_, err := databases.DB.Database.Query(
		query,
		id,
		bk.Title,
		bk.Description,
		bk.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func DeleteBookByID(id uuid.UUID, bk *Book) error {
	_, err := databases.DB.Database.Query(`UPDATE book SET deleted_at = $1 WHERE id = $2`, bk.DeletedAt, id)
	return err
}

func scanIntoBook(rows *sql.Rows) (*Book, error) {
	book := &Book{}
	err := rows.Scan(
		&book.ID,
		&book.Title,
		&book.Description,
		&book.CreatedAt,
		&book.UpdatedAt,
		&book.DeletedAt,
	)

	return book, err
}
