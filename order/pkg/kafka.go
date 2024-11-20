package pkg

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func InitProducer(broker, topic string) (*kafka.Writer, error) {
	return &kafka.Writer{
		Addr:         kafka.TCP(broker),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchSize:    100,
		RequiredAcks: kafka.RequireOne,
	}, nil
}

func ProduceMessage(writer *kafka.Writer, message string) error {
	start := time.Now()
	defer func() {
		fmt.Printf("time taken to produce message: %v\n", time.Since(start))
	}()

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte("key"),
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
