package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"
)

func main() {
	producer()
	consumer()
}

func producer() {
	writer := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "info-topic",
	}

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte("Catupiri com queijo"),
		Headers: []protocol.Header{
			{
				Key:   "session",
				Value: []byte("123"),
			},
		},
	})

	if err != nil {
		log.Fatal("cannot write a message: ", err)
	}
}

func consumer() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "consumer",
		Topic:    "info-topic",
		MinBytes: 0,
		MaxBytes: 10e6,
	})

	for i := 0; i < 1; i++ {
		message, err := reader.ReadMessage(context.Background())

		for _, val := range message.Headers {
			if val.Key == "session" && string(val.Value) == "123" {
				fmt.Print("correct session \n\n")
			}
		}

		if err != nil {
			log.Fatal("cannot receive a message: ", err)
			reader.Close()
		}

		fmt.Print("receive a message: ", string(message.Value))
	}

	reader.Close()
}
