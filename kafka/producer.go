package kafka

type Producer struct {
}

func NewProducer() Producer {
	return Producer{}
}

func (p Producer) Produce(msg string) error {
	return nil
}
