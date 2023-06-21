package sql_connection

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectDB(envName string) (*sql.DB, error) {

	url := LoadUri(envName)

	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err.Error())
	}
	return db, err
}

func LoadUri(url string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(url)

}

func CloseDB(db *sql.DB) {
	db.Close()
}

func PingDB(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err.Error())
	}
}

func GetTable(db *sql.DB, table string) (*sql.Rows, error) {
	rows, err := db.Query("SELECT * FROM " + table)
	if err != nil {
		panic(err.Error())
	}
	return rows, err
}
