package kafka

type Consumer struct {
}

func NewConsumer() Consumer {
	return Consumer{}
}

func (c Consumer) Consume() error {
	return nil
}
