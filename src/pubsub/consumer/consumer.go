package consumer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var _ Consumer = (*consumer)(nil)

type consumer struct {
}

func NewConsumer() Consumer {
	return &consumer{}
}

func (c *consumer) StartConsumer() {
	c.Configure()
}

func (c *consumer) Configure() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	fmt.Println("Consumer started...")

	// Đăng ký các topic cần lắng nghe
	consumer.SubscribeTopics([]string{"*"}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received message: %s\n", string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v\n", err)
		}
	}
}
