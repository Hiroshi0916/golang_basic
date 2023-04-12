package main

import "fmt"

// 独自の型
type MyInt int

func (mi MyInt) Print() {
	fmt.Println(mi)
}

func main() {
	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)

	// i := 100
	// //同じ型でなければ計算できない
	// fmt.Println(mi + i)

	mi.Print()
}
