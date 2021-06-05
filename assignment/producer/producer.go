package producer

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// Connect to Apache Kafka using sarama as Producer
func createProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	// NewSyncProducer creates a new SyncProducer using the given broker addresses and configuration.
	producer, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}

// Push a message to topic using sarama
func Publish(topic string, message string) error {
	// TODO How to parameterize brokerUrl "localhost:9092"?
	brokersUrl := []string{"localhost:9092"}
	producer, err := createProducer(brokersUrl)
	if err != nil {
		return err
	}

	defer producer.Close()

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
