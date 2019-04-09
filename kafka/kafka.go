package kafka

type Producer struct {
}

func NewProducer() Producer {
	return Producer{}
}

func (p Producer) Produce(msg string) error {
	return nil
}

type Consumer struct {
}

func NewConsumer() Consumer {
	return Consumer{}
}

func (c Consumer) Consume() error {
	return nil
}
