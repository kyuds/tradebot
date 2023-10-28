package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// configure these to setup kafka pubsub
func producerConfigs() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	}
}

func consumerConfigs() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	}
}

func createKafkaTopics() error {
	// Kafka Admin Client

	// Open topic lists to generate Kafka topics.
	topicFile, err := os.Open("kafka-topic-list.conf")

	if err != nil {
		return err
	}

	topicScanner := bufio.NewScanner(topicFile)
	topicScanner.Split(bufio.ScanLines)

	// Generate Topics
	for topicScanner.Scan() {
		topicName := topicScanner.Text()
		// Create topic here.
		fmt.Println(topicName)
	}

	return nil
}
