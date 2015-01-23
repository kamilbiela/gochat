package model

import (
	"encoding/json"
	// "errors"
	// "github.com/kamilbiela/gochat/pubsub"
	// "log"
)

const (
	MSG_TEXT = iota
	MSG_ROOM_JOIN
	MSG_ROOM_LEAVE
	MSG_USER_STATUS_AVAILABLE
	MSG_USER_STATUS_AWAY
	MSG_AUTH
)

type Message struct {
	Type int
	Val  interface{}
}

func NewMessage(rawMessage string) (*Message, error) {
	var m Message

	err := json.Unmarshal([]byte(rawMessage), &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// func HandleMessage(commChan chan string, pb pubsub.PubSuber, rawMessage string) error {
// 	message, err := NewMessage(rawMessage)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var closeChan chan bool

// 	switch message.Type {
// 	case MSG_JOIN_ROOM:
// 		log.Printf("User %d joins room %s \n", message.UserId, message.Val)
// 		chatName, ok := message.Val.(string)

// 		if !ok {
// 			return nil, errors.New("wrong msg")
// 		}
// 		AppCont.PubSub.Subscribe(chatName)

// 	case MSG_ON_CHAT:
// 		txt, _ := message.Val.(string)
// 		pb.Publish("chat", txt)

// 	default:
// 		log.Println("unsupported msg type")
// 	}

// 	return closeChan, nil
// }
