package db

import (
	"database/sql"
	"fmt"

	"github.com/cfung89/messaging/backend/pkg/assert"
	_ "github.com/lib/pq"
)

type UsersAuth struct {
	UID      string
	Username string
	Password string
}

func OpenDB(secret string, port int) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://postgres:%s@localhost:%d/auth?sslmode=disable", secret, port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	assert.NotNil(CheckhealthDB(db))
	return db, nil
}

func CheckhealthDB(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		return err
	}
	return nil
}

func CreateUsersAuth(db *sql.DB) error {
	/* UserAuthTable
	- UID
	- Username
	- Password
	- Date Created
	*/
	query := `CREATE TABLE IF NOT EXISTS usersAuth (
	uid SERIAL PRIMARY KEY,
	username VARCHAR(100) NOT NULL PRIMARY KEY,
	password VARCHAR(100) NOT NULL,
	created timestamp DEFAULT NOW()
	)`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// Returns primary key of inserted element
func InsertUsersAuth(db *sql.DB, user UsersAuth) (int, error) {
	query := `INSERT INTO usersAuth (uid, username, password)
VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := db.QueryRow(query, user.UID, user.Username, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func QueryUsersAuth(db *sql.DB, username string) (string, error) {
	query := "SELECT username, password usersAuth where uid = $1"
	var uid int
	var password string
	err := db.QueryRow(query, username).Scan(&uid, &password)
	if err != nil {
		return "", err
	}
	return username, nil
}
