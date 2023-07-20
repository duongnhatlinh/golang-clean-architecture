package main

import (
	"context"
	"fmt"
	"food_delivery/pubsub"
	"food_delivery/pubsub/pblocal"
	"time"
)

func main() {
	var title pubsub.Topic = "title1"

	ps1 := pblocal.NewLocalPubSub()

	ch1, close1 := ps1.Subscribe(context.Background(), title)
	ch2, close2 := ps1.Subscribe(context.Background(), title)

	ps1.Publish(context.Background(), title, pubsub.NewMessage(1))
	ps1.Publish(context.Background(), title, pubsub.NewMessage(2))

	close1()
	close2()

	go func() {
		fmt.Println("pubsub 1: ", (<-ch1).Data())
	}()
	go func() {
		fmt.Println("pubsub 2: ", (<-ch2).Data())
	}()

	time.Sleep(time.Second * 2)

}
