package main

import "fmt"

func main() {
	// map
	//for

	m := map[string]int{
		"Apple":  100,
		"Banana": 200,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	for _, v := range m {
		fmt.Println(v)
	}

	for k := range m {
		fmt.Println(k)
	}
}
