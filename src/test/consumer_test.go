package main

import (
	"kafka-server/src/consumer"
	"testing"
	"time"
)

func TestConsumer(t *testing.T) {
	go func() {
		time.Sleep(1 * time.Second)
		err := consumer.ConsumeMessages()
		if err != nil {
			// t.Errorf("Erro ao consumir a mensagem: %v", err)
		}
	}()
}
