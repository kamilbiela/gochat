package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/kamilbiela/gochat/lib"
	"github.com/kamilbiela/gochat/middleware"
	"github.com/kamilbiela/gochat/route"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"log"
	"net/http"
	"os"
)

func main() {
	// parse flags
	flagFixture := flag.Bool("fixture", false, "specify to load fixtures")
	flag.Parse()

	// init di container
	container := lib.NewContainer()

	if *flagFixture {
		fmt.Println("Loading fixtures...")
		lib.FixturesLoad(container.GetDB())
		return
	}

	authChain := alice.New(middleware.Auth(container.GetAuth()))

	router := mux.NewRouter()
	router.Handle("/chat/{v:.*}", sockjs.NewHandler("/chat", sockjs.DefaultOptions, route.SocketHandler))
	router.Methods("POST").Subrouter().Handle("/auth", route.AuthRoute(container.GetAuth()))
	router.Handle("/test", authChain.Then(route.TestRoute("test")))
	router.Handle("/", http.FileServer(http.Dir("public")))

	log.Fatal(http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, router)))
}
