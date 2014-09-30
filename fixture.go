package main

import (
	"database/sql"
	"log"
)

func FixturesLoad(db *sql.DB) {
	queries := []string{
		// organization
		`DROP TABLE IF EXISTS organization`,
		`CREATE TABLE organization (
			id INT NOT NULL AUTO_INCREMENT,
			name TEXT NOT NULL,
			PRIMARY KEY (id)
	    );`,
		`INSERT INTO organization VALUES (null, "gochat")`,

		// room
		`DROP TABLE IF EXISTS room`,
		`CREATE TABLE room (
			id INT NOT NULL AUTO_INCREMENT,
			organization_id INT NOT NULL,
			name TEXT NOT NULL,
			PRIMARY KEY (id)
	    );`,
		`INSERT INTO room VALUES (null, 1, "room")`,

		// user
		`DROP TABLE IF EXISTS user`,
		`CREATE TABLE user (
			id INT NOT NULL AUTO_INCREMENT,
			organization_id INT NOT NULL,
			name TEXT NOT NULL,
			PRIMARY KEY (id)
	    );`,
		`INSERT INTO user VALUES 
			(null, 1, "user1"),
			(null, 1, "user2"),
			(null, 1, "user3"),
			(null, 1, "user4")
		`,
	}

	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			log.Fatal(err)
		}
	}
}
