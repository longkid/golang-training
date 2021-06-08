package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"
)

const (
	BROKER_URL = "localhost:9092"
	TOPIC      = "golang-assignment-events"
)

func main() {
	consumer, err := createConsumer()
	if err != nil {
		panic(err)
	}

	defer consumer.Close()

	topic := os.Getenv("TOPIC")
	if topic == "" {
		topic = TOPIC
	}

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	fmt.Println("Consumer started ")
	sigchan := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case err := <-partitionConsumer.Errors():
				fmt.Println(err)
			case msg := <-partitionConsumer.Messages():
				fmt.Printf("Received message: Topic (%s) | Message (%s) \n", string(msg.Topic), string(msg.Value))
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				done <- true
			}

		}
	}()

	fmt.Println("Awaiting terminating signal")
	<-done
	fmt.Println("Exiting")
}

// Connect to Apache Kafka as Consumer
func createConsumer() (sarama.Consumer, error) {
	brokerUrl := os.Getenv("BROKER_URL")
	if brokerUrl == "" {
		brokerUrl = BROKER_URL
	}

	brokersUrl := []string{brokerUrl}
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// NewConsumer creates a new consumer using the given broker addresses and configuration
	consumer, err := sarama.NewConsumer(brokersUrl, config)

	return consumer, err
}
