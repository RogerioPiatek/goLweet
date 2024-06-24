package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func Init() {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Erro loading .env file.")
		}

		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")

		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbPassword, dbName)

		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Fatalf("Error opening database: %q", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Error opening database: %q", err)
		}

		fmt.Println("Succesfully connected to the database!")
	})
}

func GetDB() *sql.DB {
	return db
}
