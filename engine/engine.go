package main

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/kyuds/tradebot/engine/executor"
	"github.com/kyuds/tradebot/engine/streamer"
)

func main() {
	// Setup
	pubConf := producerConfigs()
	subConf := consumerConfigs()
	err := createKafkaTopics()

	if err != nil {
		fmt.Println(err)
		return
	}

	// Run stock processor
	var stop int32 = 0

	go func() {
		streamer.Run(&stop, pubConf)
	}()
	go func() {
		executor.Run(&stop, subConf)
	}()

	for {
		time.Sleep(60 * time.Second)
		if atomic.LoadInt32(&stop) == 1 {
			return
		}
	}
}
