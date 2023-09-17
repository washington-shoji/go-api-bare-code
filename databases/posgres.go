package databases

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/washington-shoji/gobare/configs"
)

type PostgresStore struct {
	db *sql.DB
}

var DB PostgresStore

func PostgresDB() (*PostgresStore, error) {
	connStr := fmt.Sprintf("user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configs.EnvConfig("DB_USER"), configs.EnvConfig("DB_PASSWORD"), configs.EnvConfig("DB_NAME"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	DB = PostgresStore{
		db: db,
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func Init() error {
	return CreateTables()
}
