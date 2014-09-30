package main

import (
	redisPubSub "github.com/kamilbiela/gochat/pubsub/redis"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"log"
	"net/http"
)

func socketHandler(sockSession sockjs.Session) {
	pb := redisPubSub.InitPubSub()
	pb.Subscribe("all")
	log.Println("===== SOCKET OPEN ======")

	defer pb.Unsubscribe("all")
	defer pb.Close()

	go func() {
		for message := range pb.ReadMessage() {
			if err := sockSession.Send(message); err != nil {
				log.Println("99999")
				return
			}
		}
	}()

	// receive socket messages and handle them
	for {
		if rawMessage, err := sockSession.Recv(); err == nil {
			pb.Publish("all", rawMessage)
			continue
		}
		break
	}

	log.Println("===== SOCKET CLOSE ======")
}

func initSocketServer() {
	http.Handle("/chat/", sockjs.NewHandler("/chat", sockjs.DefaultOptions, socketHandler))
}
