package producer

import (
	"github.com/segmentio/kafka-go"
)

const (
	topic               = "email-sending"
	brokerSenderAddress = "localhost:9092"
)

func EmailSending() *kafka.Writer {
	// intialize
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerSenderAddress},
		Topic:   topic,
	})
	return w

}
