package db

import (
	"fmt"

	"github.com/cfung89/messaging/backend/pkg/utils"
	"github.com/lib/pq"
)

type AuthDB struct {
	BaseDB
}

type UsersAuth struct {
	UID        string `db:"uid"`
	DBPassword string
	Username   string `json:"username" db:"username"`
	Password   string `json:"password" db:"password"`
}

func (db *AuthDB) CreateTable() error {
	/* UserAuthTable
	   - UID
	   - Username
	   - Password
	   - Date Created
	*/
	query := `CREATE TABLE IF NOT EXISTS usersAuth (
				uid VARCHAR(100) PRIMARY KEY NOT NULL,
				username VARCHAR(100) NOT NULL,
				password VARCHAR(100) NOT NULL,
				created timestamp DEFAULT NOW() NOT NULL
				)`
	_, err := db.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// Returns primary key of inserted element
func (db *AuthDB) InsertUsers(user *UsersAuth) error {
	query := `INSERT INTO usersAuth (uid, username, password)
				VALUES ($1, $2, $3) RETURNING id`
	user.UID = utils.GenerateUUID().String()
	_, err := db.DB.Exec(query, user.UID, user.Username, user.Password)
	if err != nil {
		if e, ok := err.(*pq.Error); ok {
			return fmt.Errorf("%s: %s", err, e.Detail)
		}
		return err
	}
	return nil
}

// Saves user ID and password in user parameter, given the unique username
func (db *AuthDB) QueryUser(user *UsersAuth) error {
	query := "SELECT uid, password usersAuth where username = $1"
	err := db.DB.QueryRow(query, user.Username).Scan(&user.UID, &user.DBPassword)
	if err != nil {
		return err
	}
	return nil
}
