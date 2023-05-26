package main

import (
	"fmt"
	"time"
)

func DoWorkVer1(done <-chan interface{}, palseInterval time.Duration) (<-chan interface{}, <-chan time.Time) {
	heartbeat := make(chan interface{})
	results := make(chan time.Time)

	go func() {
		defer close(heartbeat)
		defer close(results)
		heartbeatPulse := time.Tick(palseInterval)
		workGen := time.Tick(2 * palseInterval)

		sendHeartBeatPulse := func() {
			select {
			case heartbeat <- struct{}{}:
			default:

			}
		}
		sendResult := func(result time.Time) {
			for {
				select {
				case <-done:
					return

				case <-heartbeatPulse:
					sendHeartBeatPulse()

				case results <- result.Local():
					return
				}
			}
		}

		for {
			select {
			case <-done:
				return

			case <-heartbeatPulse:
				sendHeartBeatPulse()

			case result := <-workGen:
				sendResult(result)
			}
		}
	}()
	return heartbeat, results
}

func main() {
	done := make(chan interface{})
	timeout := 2 * time.Second
	heartbeat, results := DoWorkVer1(done, timeout/2)

	for {
		select {
		case _, ok := <-heartbeat:
			if !ok {
				return
			}
			fmt.Println("receive heatbeat")

		case r, ok := <-results:
			if !ok {
				return
			}
			fmt.Printf("result: %v\n", r)
		case <-time.After(timeout):
			fmt.Println("worker goroutine is dead")
			return
		}

	}
}
