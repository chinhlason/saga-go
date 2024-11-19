package pkg

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func ProduceMessage(broker, topic, message string) error {
	writer := kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	defer writer.Close()

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte("inventory"),
		Value: []byte(message),
	})
	if err != nil {
		return fmt.Errorf("failed to write message: %s", err)
	}
	return nil
}

func ConsumeMessages(broker, topic, groupId string) error {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: groupId,
	})
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			return fmt.Errorf("failed to read message: %s", err)
		}
		fmt.Printf("message received: %s\n", string(msg.Value))
	}
}
