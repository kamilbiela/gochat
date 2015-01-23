package route

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 10024
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func SocketRoute() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", 405)
			return
		}

		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		defer ws.Close()
		ws.SetReadLimit(maxMessageSize)

		ws.SetReadDeadline(time.Now().Add(pongWait))
		ws.SetWriteDeadline(time.Now().Add(writeWait))

		ws.SetPongHandler(func(string) error {
			ws.SetReadDeadline(time.Now().Add(pongWait))
			return nil
		})

		writeChan := make(chan []byte)

		// write
		go func() {
			ticker := time.NewTicker(pingPeriod)
			defer ticker.Stop()

			for {
				select {
				case message, ok := <-writeChan:
					if !ok {
						ws.WriteMessage(websocket.CloseMessage, []byte{})
						return
					}
					if err := ws.WriteMessage(websocket.TextMessage, message); err != nil {
						return
					}
				case <-ticker.C:
					if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
						return
					}
				}
				ws.SetWriteDeadline(time.Now().Add(writeWait))
			}
		}()

		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				break
			}

			log.Println("Read message: ", string(message))

			writeChan <- message
		}

	})
}

/*

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

				_, err := model.NewMessage(rawMessage)
				if err != nil {
					log.Println("====1")
					log.Println(err)
					log.Println("====1")
					continue
				}
				log.Println("======")
				log.Println(sockSession.ID())
				log.Println("======")

				pb.Publish("all", rawMessage)
				continue
			}
			break
		}

		pb.Unsubscribe("*")
		pb.Close()
		log.Println("===== SOCKET CLOSE ======")
	}

*/
