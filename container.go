package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Container struct {
	db *sql.DB
}

func NewContainer() Container {
	return Container{}
}

func (c Container) getDB() *sql.DB {
	if c.db == nil {
		db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/gochat")
		if err != nil {
			log.Fatal(err)
		}

		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}

		c.db = db
	}

	return c.db
}
