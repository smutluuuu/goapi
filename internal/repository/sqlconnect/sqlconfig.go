package sqlconnect

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() (*sql.DB, error) {
	fmt.Println("Trying to connect MariaDB")

	// connectionString := "root:root@tcp(127.0.0.1:3306)/" + dbname
	connectionString := os.Getenv("CONNECTION_STRING")
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to Maria DB")
	return db, nil
}
