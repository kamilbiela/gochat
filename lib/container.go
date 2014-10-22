package lib

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kamilbiela/gochat/mapper"
	"log"
)

type Container struct {
	db         *sql.DB
	config     *Config
	auth       *AuthService
	userMapper *mapper.UserMapper
}

func NewContainer() *Container {
	c := new(Container)
	c.config = NewConfig()
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
	if c.auth == nil {
		c.auth = NewAuthService(c.GetUserMapper())
	}
	return c.auth
}

func (c *Container) GetUserMapper() *mapper.UserMapper {
	if c.userMapper == nil {
		c.userMapper = mapper.NewUserMapper(c.GetDB())
	}

	return c.userMapper
}
