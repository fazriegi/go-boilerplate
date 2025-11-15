package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Get() *sqlx.DB {
	if DB == nil {
		log.Fatal("database connection is not initialized")
	}
	return DB
}
