package main

import (
	"fmt"
	"time"

	"github.com/kyuds/tradebot/engine/executor"
	"github.com/kyuds/tradebot/engine/streamer"
)

func main() {
	go func() {
		streamer.Run()
	}()
	go func() {
		executor.Run()
	}()
	for {
		fmt.Println("reached")
		time.Sleep(5 * time.Second)
	}
}
