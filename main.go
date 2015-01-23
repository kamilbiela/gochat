package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kamilbiela/gochat/lib"
	"github.com/kamilbiela/gochat/route"
)

func main() {
	// parse flags
	flagFixture := flag.Bool("fixture", false, "specify to load fixtures")
	flag.Parse()

	// init di container
	container := lib.NewContainer()

	if *flagFixture {
		fmt.Println("Loading fixtures...")
		lib.FixturesLoad(container.GetDB(), container.GetAuth())
		return
	}

	// middlewares
	//authMiddleware := middleware.Auth(container.GetAuth())

	// alice chains
	//authChain := alice.New(authMiddleware)

	// router
	router := mux.NewRouter()

	// routes - auth
	router.Methods("POST").Subrouter().Handle("/auth", route.AuthRoute(container.GetAuth()))

	// socket
	socketRoute := route.SocketRoute()
	//router.Handle("/chat/{token}", authChain.Then(socketRoute))
	router.Handle("/chat", socketRoute)

	// serve files
	router.Handle("/", http.FileServer(http.Dir(container.GetConfig().Webdir)))

	log.Fatal(http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, router)))
}
