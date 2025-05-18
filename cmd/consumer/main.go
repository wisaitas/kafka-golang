package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Create a reader connection to the Kafka broker
	// Use the exposed port from docker-compose
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "messages",
		GroupID:  "consumer-group-1",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	defer r.Close()

	fmt.Println("Consumer started. Press Ctrl+C to stop.")

	// Read messages from Kafka
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %s", err)
			continue
		}

		log.Printf("Received message from partition %d at offset %d: %s = %s",
			m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
