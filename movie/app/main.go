package main

import (
	"log"

	"github.com/IBM/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Version = sarama.V2_7_0_0 // Ensure this matches your Kafka version

	brokers := []string{"kafka:9092"} // or "kafka:9092" if running within Docker network
	client, err := sarama.NewClient(brokers, config)
	if err != nil {
		log.Fatalf("Error creating Kafka: %v", err)
	}

	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %v", err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Error closing Kafka producer: %v", err)
		}
	}()

	message := &sarama.ProducerMessage{
		Topic: "test",
		Value: sarama.StringEncoder("Hello Kafka"),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	log.Printf("Message sent to partition %d at offset %d", partition, offset)
}
