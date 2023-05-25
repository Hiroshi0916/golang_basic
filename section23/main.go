package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func shortProcess(ctx context.Context) (bool, error) {
	shortWork := time.NewTicker(1 * time.Second)
	ctx, cancel := context.WithTimeout(ctx, 1*time.Millisecond)
	defer cancel()

	select {
	case <-ctx.Done():
		return false, fmt.Errorf("cancel")
	case <-shortWork.C:
	}
	return true, nil
}

func longProcess(ctx context.Context) (bool, error) {
	longWork := time.NewTicker(5 * time.Second)

	select {
	case <-ctx.Done():
		return false, fmt.Errorf("cancel")
	case <-longWork.C:
	}
	return true, nil
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)

	go func() {
		defer wg.Done()

		if isDone, err := shortProcess(ctx); err != nil {
			fmt.Printf("shortProcess: %v\n", err)
			fmt.Println(isDone)

			cancel()

			return
		}
		fmt.Println("shortProcess is Done")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()

		if isDone, err := longProcess(ctx); err != nil {
			fmt.Printf("longProcess: %v\n", err)
			fmt.Println(isDone)

			return
		}
		fmt.Println("longProcess is Done")
	}()

	wg.Wait()
}
