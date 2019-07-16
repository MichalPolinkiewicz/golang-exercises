package persistence

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DbInstance *sql.DB

func init() {
	dbUri := createDbUri()
	db, err := sql.Open("mysql", dbUri)
	if err != nil {
		log.Println("error during opening db connection:", err)
	}
	DbInstance = db
}

func createDbUri() string {
	e := godotenv.Load("phone-number-normalizer/persistence/.env")
	if e != nil {
		log.Println("error during loading env file:", e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	charset := os.Getenv("charset")
	pt := os.Getenv("parseTime")
	dbUri := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=%s", username, password, dbName, charset, pt)

	return dbUri
}
