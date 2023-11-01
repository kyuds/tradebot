package streamer

import (
	"fmt"
	"sync/atomic"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Run(stop *int32, conf *kafka.ConfigMap) {
	p, err := kafka.NewProducer(conf)

	if err != nil {
		// will need better logging in the future
		fmt.Printf("Failed to create producer: %s\n", err)
		// destroy the entire process, not just this thread.
		atomic.StoreInt32(stop, 1)
		return
	}

	defer p.Close()

	fmt.Println("CREATED PRODUCER")
}
