package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/cfung89/messaging/backend/pkg/assert"
	_ "github.com/lib/pq"
)

type DB interface {
	OpenDB(secret string, port int) error
	CheckHealthDB() error
	CreateTable() error
	InsertUsers(user *UsersAuth) error
	QueryUsers(user *UsersAuth) error
}

type BaseDB struct {
	DB *sql.DB
}

func (db *BaseDB) OpenDB(secret string, port int) error {
	connStr := fmt.Sprintf("postgres://postgres:%s@localhost:%d/auth?sslmode=disable", secret, port)
	b, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	db.DB = b
	defer db.DB.Close()
	assert.NotNil(db.CheckHealthDB())
	return nil
}

func (db *BaseDB) CheckHealthDB() error {
	if db.DB == nil {
		return errors.New("Database connection is nil.")
	}
	if err := db.DB.Ping(); err != nil {
		db.DB.Close()
		return err
	}
	return nil
}
