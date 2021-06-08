package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

const (
	BROKER_URL = "localhost:9092"
	TOPIC      = "golang-assignment-events"
)

func main() {
	// Create producer
	producer, err := createProducer()
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	// Read input from console
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nEnter message to publish: ")
		scanner.Scan()
		msg := scanner.Text()

		publish(msg, producer)
	}

}

// Connect to Apache Kafka using sarama as Producer
func createProducer() (sarama.SyncProducer, error) {
	brokerUrl := os.Getenv("BROKER_URL")
	if brokerUrl == "" {
		brokerUrl = BROKER_URL
	}

	brokersUrl := []string{brokerUrl}
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	// NewSyncProducer creates a new SyncProducer using the given broker addresses and configuration.
	producer, err := sarama.NewSyncProducer(brokersUrl, config)

	return producer, err
}

// Push a message to topic using sarama
func publish(message string, producer sarama.SyncProducer) error {
	topic := os.Getenv("TOPIC")
	if topic == "" {
		topic = TOPIC
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,                         // Kafka topic for this message
		Value: sarama.StringEncoder(message), // The actual message to store in Kafka
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	fmt.Printf("Message is stored in topic(%s), partition(%d), offset(%d)", topic, partition, offset)
	return nil
}
