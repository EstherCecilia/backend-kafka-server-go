package consumer

import (
	"context"
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
)

func ConsumeMessages() {
	brokerAddress := os.Getenv("KAFKA_BROKER")
	topic := os.Getenv("KAFKA_TOPIC")

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{brokerAddress},
		Topic:     topic,
		Partition: 0,
	})

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Errorf("Erro ao ler mensagem: %v\n", err)
		}
		fmt.Printf("Consumer consumiu: %s\n", string(msg.Value))
	}
}
