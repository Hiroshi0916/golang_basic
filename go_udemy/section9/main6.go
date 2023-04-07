package main

import "fmt"

func main() {
	// map
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "Japan"
	m4[2] = "USA"
	fmt.Println(m4)
	fmt.Println(m["A"])
	fmt.Println(m4[2])

	fmt.Println(m4[3])

	//mapにはエラーハンドリングの機能がある。
	s := m4[1]
	fmt.Println(s)

	// s, ok := m4[1]
	// fmt.Println(s, ok)

	// s, ok := m4[3]
	// fmt.Println(s, ok)

	s, ok := m4[3]
	if !ok {
		fmt.Println("error")
	}

	m4[2] = "US"
	fmt.Println(m4)

	m4[3] = "China"
	fmt.Println(m4)

	delete(m4, 3)
	fmt.Println(m4)
	fmt.Println(len(m4))
}
