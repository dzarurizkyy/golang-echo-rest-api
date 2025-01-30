package db

import (
	"database/sql"
	"fmt"
	"golang-echo-rest-api/config"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.DB_HOST, conf.DB_PORT, conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_NAME)

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		panic("Conection error...")
	}

	err = db.Ping()
	if err != nil {
		panic("Data source name error...")
	}

	fmt.Println("Successfully connected..")
}

func CreateCon() *sql.DB {
	return db
}
