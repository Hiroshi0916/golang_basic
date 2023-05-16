package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string
}

func main() {

	smap := &sync.Map{}
	smap.Store("Hello", "World")
	smap.Store(1, 2)

	smap.Delete(1)

	smap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

	v, ok := smap.Load("Hello")
	if ok {
		fmt.Println(v)
	}
	smap.LoadOrStore("Hello", "Woooorld")

	smap.LoadOrStore(2, 3)

	smap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

}
