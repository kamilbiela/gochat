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

	// handlers
	sockjsHanlder := sockjs.NewHandler("/chat/[a-zA-Z0-9]+", sockjs.DefaultOptions, route.SocketHandler)

	// middlewares
	authMiddleware := middleware.Auth(container.GetAuth())

	// alice chains
	authChain := alice.New(authMiddleware)

	// router
	router := mux.NewRouter()

	// routes - auth
	router.Handle("/chat/{token}/{v:.*}", authChain.Then(sockjsHanlder))
	router.Methods("POST").Subrouter().Handle("/auth", route.AuthRoute(container.GetAuth()))

	// routes - other
	router.Handle("/test", authChain.Then(route.TestRoute("Works")))

	// serve files
	router.Handle("/", http.FileServer(http.Dir(container.GetConfig().Webdir)))

	log.Fatal(http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, router)))
}
