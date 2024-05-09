package dbpostgres

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbname := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbParams := os.Getenv("DB_PARAMS")

	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s", dbUsername, dbPassword, dbHost, dbPort, dbname, dbParams)
	fmt.Println(connection)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)

	err = db.PingContext(context.Background())
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil

}
