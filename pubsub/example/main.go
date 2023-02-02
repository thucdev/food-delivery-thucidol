package main

import (
	"context"
	"log"
	"thucidol/pubsub"
	"thucidol/pubsub/localpubsub"
	"time"
)

func main() {
	var localPS pubsub.PubSub = localpubsub.NewPubSub()

	var topic pubsub.Topic = "Order created"

	sub1, close1 := localPS.Subscribe(context.Background(), topic)
	sub2, _ := localPS.Subscribe(context.Background(), topic)

	localPS.Publish(context.Background(), topic, pubsub.NewMessage(1))
	localPS.Publish(context.Background(), topic, pubsub.NewMessage(2))

	go func() {
		for {
			log.Println("sub1:", (<-sub1).Data())
			time.Sleep(time.Millisecond * 10)
		}
	}()

	go func() {
		for {
			log.Println("sub2:", (<-sub2).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	time.Sleep(time.Second * 3)

	close1()

	localPS.Publish(context.Background(), topic, pubsub.NewMessage(3))
	time.Sleep(time.Second * 3)

}
