package pblocal

import (
	"context"
	"food_delivery/common"
	"food_delivery/pubsub"
	"log"
	"sync"
)

type localPubSub struct {
	messageQueue    chan *pubsub.Message
	mapTopicChannel map[pubsub.Topic][]chan *pubsub.Message
	lock            *sync.RWMutex
}

func NewLocalPubSub() *localPubSub {
	ps := &localPubSub{
		messageQueue:    make(chan *pubsub.Message, 1000),
		mapTopicChannel: make(map[pubsub.Topic][]chan *pubsub.Message),
		lock:            new(sync.RWMutex),
	}

	ps.Run()

	return ps
}

func (ps *localPubSub) Run() {
	log.Println("Start pubsub")

	go func() {
		for {
			mess := <-ps.messageQueue

			if chans, ok := ps.mapTopicChannel[mess.GetTopic()]; ok {
				for i := range chans {
					go func(c chan *pubsub.Message) {
						c <- mess
					}(chans[i])
				}
			}
		}
	}()
}

func (ps *localPubSub) Publish(ctx context.Context, title pubsub.Topic, message *pubsub.Message) error {
	message.SetTopic(title)

	go func() {
		defer common.AppRecover()
		ps.messageQueue <- message
	}()

	return nil
}

func (ps *localPubSub) Subscribe(ctx context.Context, title pubsub.Topic) (ch <-chan *pubsub.Message, close func()) {
	c := make(chan *pubsub.Message)

	ps.lock.Lock()

	if val, ok := ps.mapTopicChannel[title]; ok {
		val = append(ps.mapTopicChannel[title], c)
		ps.mapTopicChannel[title] = val
	} else {
		ps.mapTopicChannel[title] = []chan *pubsub.Message{c}
	}

	ps.lock.Unlock()
	return c, func() {
		log.Println("Unsubscribe")

		if chans, ok := ps.mapTopicChannel[title]; ok {
			for i := range chans {
				if c == chans[i] {
					chans = append(chans[:i], chans[i+1:]...)

					ps.lock.Lock()
					ps.mapTopicChannel[title] = chans
					ps.lock.Unlock()
					break
				}
			}
		}
	}
}
