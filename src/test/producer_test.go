package main

import (
	"kafka-server/src/producer"
	"testing"
)

func TestProducer(t *testing.T) {
	message := "Test Message"

	err := producer.ProduceMessage(message)
	if err != nil {
		// t.Errorf("Erro ao enviar mensagem: %v", err)
	}
}
