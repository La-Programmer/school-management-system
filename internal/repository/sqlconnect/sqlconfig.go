package sqlconnect

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb(dbName string) (*sql.DB, error) {
	fmt.Println("Trying to connect to MySql")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	connectionString := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", dbUser, dbPassword, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		// panic(err)
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	if err := db.Ping(); err != nil {
		log.Fatalf("Database unreachable: %v", err)
	}

	fmt.Println("Connected to MySQL.")
	return db, nil
}
