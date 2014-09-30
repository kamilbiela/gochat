package main

import (
	"flag"
	"fmt"
	"log"
	// "github.com/jinzhu/gorm"
	// elastigo "github.com/mattbaird/elastigo/lib"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"

	// "github.com/kamilbiela/gochat/pubsub"
	// redisPubSub "github.com/kamilbiela/gochat/pubsub/redis"
)

type App struct {
	DB *sql.DB
}

var AppCont App = App{}

func main() {
	flagFixture := flag.Bool("fixture", false, "specify to load fixtures")
	flag.Parse()

	AppCont.DB = initDB()

	if *flagFixture {
		fmt.Println("Loading fixtures...")
		FixturesLoad(AppCont.DB)
		return
	}

	initSocketServer()
	initHttp()
	startHttp()
}

func initHttp() {
	http.Handle("/", http.FileServer(http.Dir("old/public/")))
}

func startHttp() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/gochat")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func messageHandler() {

}
