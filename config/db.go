package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host,
		port,
		dbName,
		user,
		password,
		sslMode)
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("connot connect to DB: ", err)
	}

	DB.SetConnMaxLifetime(time.Minute * 30)
	DB.SetMaxOpenConns(5)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxIdleTime(5 * time.Minute)

	if err := DB.Ping(); err != nil {
		log.Fatal("connot ping DB: ", err)
	}
	log.Println("Connected to DB")
}
