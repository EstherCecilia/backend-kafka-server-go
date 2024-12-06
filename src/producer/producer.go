package producer

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

func ProduceMessage(message string) error {
	brokerAddress := os.Getenv("KAFKA_BROKER")
	topic := os.Getenv("KAFKA_TOPIC")

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte(message),
	})
	if err != nil {
		return fmt.Errorf("Erro ao enviar mensagem: %v", err)
	}

	fmt.Printf("Producer enviou: %s\n", message)
	return nil
}

func StartProducer() {
	for i := 1; i <= 5; i++ {
		message := fmt.Sprintf("Mensagem %d", i)
		err := ProduceMessage(message)
		if err != nil {
			fmt.Errorf("Erro ao ler mensagem:", err)
		}
		time.Sleep(1 * time.Second)
	}
}
