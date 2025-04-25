package sqlconnect

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() (*sql.DB, error) {

	// connectionString := "root:root@tcp(127.0.0.1:3306)/" + dbname
	connectionString := os.Getenv("CONNECTION_STRING")
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
