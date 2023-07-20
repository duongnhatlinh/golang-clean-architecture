package pubsub

import "time"

type Topic string

type Message struct {
	id        int
	title     Topic
	data      interface{}
	createdAt time.Time
}

func NewMessage(data interface{}) *Message {
	now := time.Now().UTC()
	return &Message{
		id:        time.Now().Nanosecond(),
		data:      data,
		createdAt: now,
	}
}

func (m *Message) SetTopic(title Topic) {
	m.title = title
}

func (m *Message) Data() interface{} {
	return m.data
}

func (m *Message) GetTopic() Topic {
	return m.title
}
