package pubsub

type PubSuber interface {
	Subscribe(channelName string) (chan string, chan bool)
	Unsubscribe() error
	ReadMessage() string
}
