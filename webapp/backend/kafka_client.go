package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type kafkaClient struct {
	producer *kafka.Producer
	msgChan  chan []byte
}

func (c *kafkaClient) produce(msg <-chan []byte) {
	go func() {
		for e := range c.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	topic := "tweets"
	for msg != nil {
		select {
		case m := <-msg:
			c.producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          []byte(m),
			}, nil)
			c.producer.Flush(15 * 1000)
		}
	}
	// Wait for message deliveries before shutting down

}

func newProducer() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})

	if err != nil {
		panic(err)
	}
	return p
}
