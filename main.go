package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	writer := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "info-topic",
	}

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte("mensagem"),
	})

	if err != nil {
		log.Fatal("cannot write a message:", err)
	}
}
