package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			ch2 <- i
		}
	}()

loop:
	for {
		select {
		case v, ok := <-ch1:
			if !ok {
				break loop
			}
			fmt.Printf("ch1: %v\n", v)

		case v, ok := <-ch2:
			if !ok {
				break loop
			}
			fmt.Printf("ch2: %v\n", v)

		}
	}
}
