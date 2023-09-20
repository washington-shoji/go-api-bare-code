package author

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/washington-shoji/gobare/databases"
)

func CreateAuthor(aut *Author) error {

	query := `
	INSERT INTO author (id, full_name, created_at)
	VALUES ($1, $2, $3)
	`
	if _, err := databases.DB.Database.Query(
		query,
		aut.ID,
		aut.FullName,
		aut.CreatedAt,
	); err != nil {
		return err
	}

	return nil
}

func GetAuthors() ([]*Author, error) {
	rows, err := databases.DB.Database.Query(`SELECT * FROM author WHERE deleted_at IS NULL`)
	if err != nil {
		return nil, err
	}

	authors := []*Author{}
	for rows.Next() {
		aut, err := scanIntoAuthor(rows)
		if err != nil {
			return nil, err
		}
		authors = append(authors, aut)

	}

	return authors, nil
}

func GetAuthorByID(id uuid.UUID) (*Author, error) {
	rows, err := databases.DB.Database.Query(`SELECT * FROM author WHERE id = $1 AND deleted_at IS NULL`, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoAuthor(rows)
	}

	return nil, fmt.Errorf("author %d not found", id)
}

func GetAuthorByFullName(fullName string) (*Author, error) {
	rows, err := databases.DB.Database.Query(`SELECT * FROM author WHERE full_name = $1`, fullName)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAuthor(rows)
	}

	return nil, fmt.Errorf("author with full-name [%s] not found", fullName)
}

func UpdateAuthorByID(id uuid.UUID, aut *Author) error {
	query := `UPDATE author
	SET full_name = $2, updated_at = $3
	WHERE id = $1
	`

	_, err := databases.DB.Database.Query(
		query,
		id,
		aut.FullName,
		aut.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAuthorByID(id uuid.UUID, aut *Author) error {
	_, err := databases.DB.Database.Query(`UPDATE author SET deleted_at = $1 WHERE id = $2`,
		aut.DeletedAt, id)
	if err != nil {
		return err
	}
	return nil
}

func scanIntoAuthor(rows *sql.Rows) (*Author, error) {
	aut := &Author{}
	err := rows.Scan(
		&aut.ID,
		&aut.FullName,
		&aut.CreatedAt,
		&aut.UpdatedAt,
		&aut.DeletedAt,
	)

	return aut, err
}
