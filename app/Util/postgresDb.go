package Util

import "github.com/go-pg/pg/v10"

type Database struct {
	*pg.DB
}

var dbConnection *pg.DB

func ConnectToDatabase() {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "example",
		Database: "postgres",
	})

	dbConnection = db
}

func GetDatabaseConnection() *pg.DB {
	return dbConnection
}
