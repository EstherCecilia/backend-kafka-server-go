package main

import (
	"fmt"
	"kafka-server/src/config"
	"kafka-server/src/consumer"
	"kafka-server/src/producer"
	"time"
)

func main() {
	config.InitEnvs()

	go producer.StartProducer()

	time.Sleep(2 * time.Second)

	fmt.Println("Iniciando Consumer...")
	consumer.ConsumeMessages()
}
