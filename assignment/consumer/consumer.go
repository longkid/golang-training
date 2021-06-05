package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"
)

func main() {

	brokersUrl := []string{"localhost:9092"}
	consumer, err := createConsumer(brokersUrl)
	if err != nil {
		panic(err)
	}

	defer consumer.Close()

	topic := "golang-assignment-events"
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
func createConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// NewConsumer creates a new consumer using the given broker addresses and configuration
	consumer, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}
