package executor

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Run() {
	_, _ = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	fmt.Println("executor")
}
