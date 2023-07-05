package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func NewDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("%v://%v:%v@%v/%v?sslmode=disable", os.Getenv("DB_TYPE"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	dbCon, err := sql.Open(os.Getenv("DB_TYPE"), connStr)

	if err != nil {
		return nil, err
	}

	db = dbCon

	return dbCon, nil
}

func GetDB() *sql.DB {
	return db
}
