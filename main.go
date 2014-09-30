package main

import (
	"flag"
	"fmt"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"log"
	"net/http"
)

func main() {
	// parse flags
	flagFixture := flag.Bool("fixture", false, "specify to load fixtures")
	flag.Parse()

	// init di container
	container := NewContainer()

	if *flagFixture {
		fmt.Println("Loading fixtures...")
		FixturesLoad(container.getDB())
		return
	}

	http.Handle("/chat/", sockjs.NewHandler("/chat", sockjs.DefaultOptions, socketHandler))
	http.Handle("/", http.FileServer(http.Dir("old/public/")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
