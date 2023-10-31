package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// configure these to setup kafka pubsub
func producerConfigs() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": "kafka-1:29092",
	}
}

func consumerConfigs() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": "kafka-1:29092",
		"auto.offset.reset": "latest",
		"group.id":          "group",
	}
}

func createKafkaTopics() error {
	// Open topic lists to generate Kafka topics.
	topicFile, err := os.Open("kafka-topic-list.conf")
	topics := make([]kafka.TopicSpecification, 0)

	if err != nil {
		fmt.Printf("Failed to open topic list: %s\n", err)
		return err
	}

	topicScanner := bufio.NewScanner(topicFile)
	topicScanner.Split(bufio.ScanLines)

	// Generate Topic Slice
	for topicScanner.Scan() {
		topicName := topicScanner.Text()
		topics = append(topics, kafka.TopicSpecification{
			Topic:             topicName,
			NumPartitions:     1,
			ReplicationFactor: 1,
		})
	}

	// Kafka Admin Client
	admin, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": "kafka-1:29092",
	})

	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
		return err
	}
	defer admin.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	maxDuration, err := time.ParseDuration("60s")
	if err != nil {
		panic("ParseDuration(60s)")
	}

	createdTopics, err := admin.CreateTopics(
		ctx,
		topics,
		kafka.SetAdminOperationTimeout(maxDuration))
	if err != nil {
		fmt.Printf("Failed to create topic: %v\n", err)
		return err
	}

	for _, tp := range createdTopics {
		if tp.Error.Code() != kafka.ErrNoError {
			fmt.Printf("Failed to create %s\n", tp.Topic)
			return errors.New("failed to create a topic")
		}
	}

	return nil
}
