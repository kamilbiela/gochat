package redis

import (
	"github.com/garyburd/redigo/redis"
)

type PubSub struct {
	messageChan chan string
	psc         redis.PubSubConn
}

func InitPubSub() *PubSub {
	return &PubSub{
		messageChan: make(chan string),
		psc:         redis.PubSubConn{redisPool.Get()},
	}
}

func (pb PubSub) Subscribe(channelName string) error {
	return pb.psc.Subscribe(channelName)
}

func (pb PubSub) Unsubscribe(channelName string) error {
	return pb.psc.Unsubscribe(channelName)
}

func (pb PubSub) ReadMessage() chan string {

	go func() {
		for {
			switch v := pb.psc.Receive().(type) {
			case redis.Message:
				pb.messageChan <- string(v.Data)

			case redis.Subscription:
				if v.Count == 0 {
					pb.psc.Close()
					close(pb.messageChan)
					return
				}

			case error:
				// @todo handle redis errors
				panic(v)
				return
			}
		}
	}()

	return pb.messageChan
}

func (pb PubSub) Close() {
	pb.psc.PUnsubscribe("")
	return
}

func (pb PubSub) Publish(channelName string, message string) error {
	c := redisPool.Get()
	defer c.Close()

	c.Send("PUBLISH", channelName, message)
	if err := c.Flush(); err != nil {
		return err
	}

	return nil
}
