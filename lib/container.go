package lib

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Container struct {
	db     *sql.DB
	config *Config
	auth   *AuthService
}

func NewContainer() *Container {
	c := new(Container)
	c.config = NewConfig()
	c.auth = NewAuthService()
	return c
}

func (c *Container) GetDB() *sql.DB {
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

func (c *Container) GetConfig() *Config {
	return c.config
}

func (c *Container) GetAuth() *AuthService {
	return c.auth
}
