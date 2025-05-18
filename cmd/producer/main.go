package main

import producer_initials "github.com/wisaitas/kafka-golang/internal/producer/initials"

// createTopic creates a topic if it doesn't exist
// func createTopic(brokerAddr, topic string, partitions int) {
// 	conn, err := kafka.Dial("tcp", brokerAddr)
// 	if err != nil {
// 		log.Fatalf("Failed to dial leader: %s", err)
// 	}
// 	defer conn.Close()

// 	// Get controller
// 	controller, err := conn.Controller()
// 	if err != nil {
// 		log.Fatalf("Failed to get controller: %s", err)
// 	}

// 	// Connect to controller
// 	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, fmt.Sprintf("%d", controller.Port)))
// 	if err != nil {
// 		log.Fatalf("Failed to connect to controller: %s", err)
// 	}
// 	defer controllerConn.Close()

// 	// Create topic
// 	err = controllerConn.CreateTopics(kafka.TopicConfig{
// 		Topic:             topic,
// 		NumPartitions:     partitions,
// 		ReplicationFactor: 1,
// 	})
// 	if err != nil {
// 		log.Printf("Error creating topic (may already exist): %s", err)
// 	} else {
// 		log.Printf("Created topic: %s", topic)
// 	}
// }

// func main() {
// 	// Topic name
// 	topic := "messages"
// 	brokerAddr := "localhost:9092"

// 	// Create topic first
// 	createTopic(brokerAddr, topic, 1)

// 	// Wait a moment for topic creation to propagate
// 	time.Sleep(2 * time.Second)

// 	// Create a writer connection to the Kafka broker
// 	w := &kafka.Writer{
// 		Addr:     kafka.TCP(brokerAddr),
// 		Topic:    topic,
// 		Balancer: &kafka.LeastBytes{},
// 	}
// 	defer w.Close()

// 	fmt.Println("Producer started. Press Ctrl+C to stop.")

// 	// Send messages every 1 second
// 	for i := 0; ; i++ {
// 		msg := fmt.Sprintf("Message #%d - sent at %s", i, time.Now().Format(time.RFC3339))

// 		err := w.WriteMessages(context.Background(), kafka.Message{
// 			Key:   []byte(fmt.Sprintf("key-%d", i)),
// 			Value: []byte(msg),
// 		})

// 		if err != nil {
// 			log.Printf("Failed to send message: %s", err)
// 		} else {
// 			log.Printf("Sent: %s", msg)
// 		}

// 		time.Sleep(1 * time.Second)
// 	}
// }

func main() {
	producer_initials.Initial()
}
