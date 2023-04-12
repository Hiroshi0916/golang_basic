package main

import "fmt"

func main() {

	//slice map channelは参照渡し
	// sl := []int{100, 200}
	// sl2 := sl
	// sl2[0] = 1000
	// fmt.Println(sl)

	// var i int = 10
	// i2 := i
	// i2 = 100
	// fmt.Println(i, i2)

	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5, 10)
	fmt.Println(sl2)

	//nはcopyに成功した数
	//先頭から塗りつぶすようにcopyしていく
	n := copy(sl2, sl)
	fmt.Println(n, sl2)

}
