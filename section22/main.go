package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
	name  string
}

func main() {
	var wg sync.WaitGroup

	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		fmt.Printf("v がロックを取得しました\n", v1.name)
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		fmt.Printf("v がロックを取得しました\n", v2.name)
		defer v2.mu.Unlock()

		fmt.Println(v1.value + v2.value)
	}
	var a value = value{name: "a"}
	var b value = value{name: "b"}

	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)

	var memoryAccess sync.Mutex
	var data int
	go func() {
		defer wg.Done()
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
	}()

	wg.Wait()

	memoryAccess.Lock()
	if data == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(data)
	}
	memoryAccess.Unlock()

	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("1st Goroutin start")
		time.Sleep(1 * time.Second)
		fmt.Println("1st Goroutine Done")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("2nd Goroutin start")
		time.Sleep(1 * time.Second)
		fmt.Println("2nd Goroutine Done")
	}()

	wg.Wait()

	say := "Hello"

	wg.Add(1)

	go func() {
		defer wg.Done()
		say = "Good Bye"
	}()

	wg.Wait()

	fmt.Println(say)

	tasks := []string{"A", "B", "C"}

	for _, task := range tasks {
		wg.Add(1)

		go func(task string) {
			defer wg.Done()
			fmt.Println(task)
		}(task)
	}
	wg.Wait()

}
