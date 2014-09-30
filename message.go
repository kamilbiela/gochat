package main

import (
	"encoding/json"
	// "errors"
	// "github.com/kamilbiela/gochat/pubsub"
	// "log"
)

const (
	MSG_ON_CHAT = iota
	MSG_JOIN_ROOM
	MSG_LEAVE_ROOM
	MSG_USER_AVAILABLE
	MSG_USER_AWAY
	MSG_USER_DONT_DISTURB
)

type JsonMessage struct {
	Command string
	Token   string
	Val     interface{}
}

type Message struct {
	Type   int
	Token  string
	UserId int
	Val    interface{}
}

func NewMessage(rawMessage string) (Message, error) {
	var jm JsonMessage
	var m Message

	err := json.Unmarshal([]byte(rawMessage), &jm)
	if err != nil {
		panic(err)
		return m, err
	}

	var t int

	switch jm.Command {
	case "message":
		t = MSG_ON_CHAT
	case "join":
		t = MSG_JOIN_ROOM
	default:
		// @todo this is error
		t = MSG_ON_CHAT
	}

	m = Message{
		Type:   t,
		Token:  jm.Token,
		UserId: GetUserIdForToken(jm.Token),
		Val:    jm.Val,
	}

	return m, nil
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
