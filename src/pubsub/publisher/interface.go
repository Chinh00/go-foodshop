package publisher

type Publisher interface {
	Configure() Publisher
	Publish(topic string, data []byte) error
}
