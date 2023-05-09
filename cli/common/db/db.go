package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", "user", "password", "dbname")
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}