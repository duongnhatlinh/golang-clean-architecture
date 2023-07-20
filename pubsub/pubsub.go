package pubsub

import "context"

type Pubsub interface {
	Publish(ctx context.Context, title Topic, message *Message) error
	Subscribe(ctx context.Context, title Topic) (ch <-chan *Message, close func())
}
