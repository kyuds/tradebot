package executor

import (
	"fmt"
	"sync/atomic"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Run(stop *int32, conf *kafka.ConfigMap) {
	_, err := kafka.NewConsumer(conf)

	if err != nil {
		// will need better logging in the future
		fmt.Printf("Failed to create consumer: %s\n", err)
		// destroy the entire process, not just this thread.
		atomic.StoreInt32(stop, 1)
		return
	}
}
