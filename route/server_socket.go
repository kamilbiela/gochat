package route

import (
	"github.com/kamilbiela/gochat/pubsub"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"log"
)

func SocketHandler(sockSession sockjs.Session) {
	log.Println("===== SOCKET OPEN ======")

	pb := pubsub.InitPubSub()
	pb.Subscribe("all")

	go func() {
		for message := range pb.ReadMessage() {
			if err := sockSession.Send(message); err != nil {
				log.Fatalln(err)
				return
			}
		}
	}()

	// receive socket messages and handle them
	for {
		if rawMessage, err := sockSession.Recv(); err == nil {

			// decide what to do with received message

			pb.Publish("all", rawMessage)
			continue
		}
		break
	}

	pb.Unsubscribe("all")
	pb.Close()
	log.Println("===== SOCKET CLOSE ======")
}
