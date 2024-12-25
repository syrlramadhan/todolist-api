package config

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

func OpenConnectionPostgres() (*sql.DB, error) {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	psqlMerge := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pwd, dbName)

	db, err := sql.Open("postgres", psqlMerge)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db, nil
}

func GetEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
