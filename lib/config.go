package lib

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Mysql struct {
		Db   string
		User string
		Pass string
	}
	Webdir string
}

func NewConfig() *Config {
	c := new(Config)

	data, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		log.Fatalln(err)
	}

	if err = yaml.Unmarshal([]byte(data), c); err != nil {
		log.Fatalf("Can't parse config file error: %v", err)
	}

	return c
}
