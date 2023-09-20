package databases

func CreateTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS account (
		id SERIAL PRIMARY KEY NOT NULL,
		user_name VARCHAR(50) UNIQUE NOT NULL,
		email VARCHAR(50) UNIQUE NOT NULL,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		encrypted_password VARCHAR(256) NOT NULL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP,
		deleted_at TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS hobby (
		id SERIAL PRIMARY KEY NOT NULL,
		title VARCHAR(50) UNIQUE NOT NULL,
		description VARCHAR(512),
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP,
		deleted_at TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS collectable (
		id SERIAL PRIMARY KEY NOT NULL,
		title VARCHAR(50) UNIQUE NOT NULL,
		description VARCHAR(512),
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP,
		deleted_at TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS author (
		id uuid PRIMARY KEY NOT NULL,
		full_name VARCHAR(50) UNIQUE NOT NULL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP,
		deleted_at TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS book (
		id uuid PRIMARY KEY NOT NULL,
		title VARCHAR(50) UNIQUE NOT NULL,
		description VARCHAR(500),
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP,
		deleted_at TIMESTAMP
	);
	`

	_, err := DB.Database.Exec(query)
	return err
}
