package database

import (
	"database/sql"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

func connection() *sql.DB {
	var db *sql.DB 
	var err error
	if db, err = sql.Open("/data/pravalika.db", ""); err != nil {
		slog.Error(err.Error())
	}
	return db
}

func Post()   {}
func Put()    {}
func Update() {}
func Fetch()  {}
