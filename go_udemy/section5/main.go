package main

import "fmt"

//const value

const Pi = 3.14
const (
	URL      = "http://yahoo.co.jp"
	SiteName = "YAHOO"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// var Big int = 999999999999999999999 + 1

const (
	c0 = iota
	c1
	c2
	c3
)

func main() {

	fmt.Println(Pi)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)
	// fmt.Println(Big - 1)
	fmt.Println(c0, c1, c2, c3)
}
