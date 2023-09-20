package account

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/washington-shoji/gobare/databases"
)

func CreateAccount(acc *Account) error {
	query := `
	INSERT INTO account (user_name, email, first_name, last_name, encrypted_password, created_at)
	VALUES ($1, $2, $3, $4, $5, $6) 
	`

	_, err := databases.DB.Database.Query(
		query,
		acc.UserName,
		acc.Email,
		acc.FirstName,
		acc.LastName,
		acc.EncryptedPassword,
		acc.CreatedAt,
	)
	if err != nil {
		return err
	}

	// fmt.Printf("%+v\n", resp)

	return nil
}

func GetAccounts() ([]*Account, error) {
	rows, err := databases.DB.Database.Query(`SELECT * FROM account WHERE deleted_at IS NULL`)
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		acc, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, acc)
	}

	return accounts, nil
}

func GetAccountByID(id int) (*Account, error) {
	rows, err := databases.DB.Database.Query(`SELECT * FROM account WHERE id = $1 AND deleted_at IS NULL`, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func GetAccountByUserName(userName string) (*Account, error) {
	rows, err := databases.DB.Database.Query(`SELECT * FROM account WHERE user_name = $1`, userName)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account with number [%s] not found", userName)
}

func UpdateAccountByID(id int, acc *Account) error {
	query := `UPDATE account 
	SET 
	user_name = $2, 
	email = $3, 
	first_name = $4, 
	last_name = $5, 
	encrypted_password = $6, 
	updated_at = $7
	WHERE id = $1`

	_, err := databases.DB.Database.Query(
		query,
		id,
		acc.UserName,
		acc.Email,
		acc.FirstName,
		acc.LastName,
		acc.EncryptedPassword,
		acc.UpdatedAt,
	)
	if err != nil {
		return err
	}

	// fmt.Printf("%+v\n", resp)

	return nil
}

func DeleteAccountByID(id int, acc *Account) error {
	_, err := databases.DB.Database.Query(`UPDATE account SET deleted_at = $1 WHERE id = $2`, acc.DeletedAt, id)
	return err
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := &Account{}
	err := rows.Scan(
		&account.ID,
		&account.UserName,
		&account.Email,
		&account.FirstName,
		&account.LastName,
		&account.EncryptedPassword,
		&account.CreatedAt,
		&account.UpdatedAt,
		&account.DeletedAt,
	)

	return account, err
}
