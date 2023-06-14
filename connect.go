package sql

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func LoadUri(url string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(url)

}
