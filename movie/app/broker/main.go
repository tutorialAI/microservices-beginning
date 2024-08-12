package broker

import (
	"log"

	// db "app/db/mysql"

	"github.com/IBM/sarama"
)

func main() {
	// db.Connect()
	// alex, err := db.SelectByName("Alex")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// users, err := db.SelectUsers()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(users)
	// fmt.Println(alex)

	// id, _ := db.AddUser(db.User{
	// 	Name:  "Emma",
	// 	Email: "email@mail.com",
	// })

	// fmt.Printf("new user was added with id %d \n", id)

	kafkaConnect()
}

func kafkaConnect() {
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
