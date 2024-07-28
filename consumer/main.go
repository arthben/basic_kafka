package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// kafka-1, kafka-2, kafka-3 is name of kafka broker
	// register kafka-1, kafka-2, kafka-3 on the /etc/host
	// and point to IP address of kafka broker server
	servers := "kafka-1:9192,kafka-2:9292,kafka-3:9392"

	// if the kafka server doesn't enable auto topic creation,
	// create the topic first using kafka CLI.
	topic := "kafka_arthben"
	group := "firstClient1"
	readTimeout := 5 * time.Second

	config := &kafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          group,
		"auto.offset.reset": "earliest",
	}
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	if err := consumer.Subscribe(topic, nil); err != nil {
		panic(err)
	}

	for {
		message, err := consumer.ReadMessage(readTimeout)
		if err == nil {
			fmt.Printf("message.Key: %v\n", string(message.Key))
			fmt.Printf("message.Value: %v\n", string(message.Value))
		}
	}
}
