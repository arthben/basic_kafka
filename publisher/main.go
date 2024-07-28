package main

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
)

func main() {
	// kafka-1, kafka-2, kafka-3 is name of kafka broker
	// register kafka-1, kafka-2, kafka-3 on /etc/host
	// and point to IP address of kafka broker server
	servers := "kafka-1:9192,kafka-2:9292,kafka-3:9392"

	// if the kafka server doesn't enable auto topic creation,
	// create the topic first using kafka CLI.
	topic := "kafka_arthben"

	config := &kafka.ConfigMap{
		"bootstrap.servers": servers,
	}
	producer, err := kafka.NewProducer(config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	// publish message to kafka
	key := uuid.NewString()
	messageToPublish := composeMessage(topic, key, "Hai..")
	if err := producer.Produce(&messageToPublish, nil); err != nil {
		panic(err)
	}

	producer.Flush(1 * 1000)
}

func composeMessage(topic string, key string, message string) kafka.Message {
	return kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value:         []byte(message),
		Key:           []byte(key),
		Timestamp:     time.Time{},
		TimestampType: 0,
		Opaque:        nil,
		Headers:       []kafka.Header{},
	}
}
