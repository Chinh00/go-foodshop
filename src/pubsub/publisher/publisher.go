package publisher

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type publisher struct{}

func (*publisher) Configure() Publisher {
	//TODO implement me
	panic("implement me")
}

func (*publisher) Publish(topic string, data []byte) error {
	pub, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		panic(err)
	}

	chanel := make(chan kafka.Event)
	pub.Produce(&kafka.Message{
		Value: data,
		TopicPartition: kafka.TopicPartition{
			Topic: &topic,
		},
	}, chanel)
	e := <-chanel
	print(e)
	return nil
}
