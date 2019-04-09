package kafka

type Producer struct {
}

func (p Producer) Produce(msg string) error {
	return nil
}

type Consumer struct {
}

func (c Consumer) Consume() error {
	return nil
}
