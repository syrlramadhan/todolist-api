package config

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go-restful/util"
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
	util.SendPanicIfError(err)

	err = db.Ping()
	util.SendPanicIfError(err)

	return db, nil
}

func GetEnv() {
	err := godotenv.Load()
	util.SendPanicIfError(err)
}
