package main

import (
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
}
