package main

import (
	"fmt"
)

func main() {

	//int
	// var i int32 = 100

	// var i2 int64 = 200

	// fmt.Println(i + 50)
	// fmt.Println(i2)

	// fmt.Println(i + i2)

	// fmt.Printf("%T\n", i2)

	// fmt.Printf("%T\n", int32(i2))
	// fmt.Printf("%T\n", i)

	// fmt.Println(i + int32(i2))

	//float
	// var fl64 float64 = 2.4
	// fmt.Println(fl64)

	// fl := 3.2
	// fmt.Println(fl64 + fl)
	// fmt.Printf("%T, %T\n", fl64, fl)

	// var fl32 float32 = 1.2
	// fmt.Printf("%T\n", fl32)

	// fmt.Printf("%T\n", float64(fl32))

	// zero := 0.0
	// pinf := 1.0 / zero
	// fmt.Println(pinf)

	// ninf := -1.0 / zero
	// fmt.Println(ninf)

	// nan := zero / zero
	// fmt.Println(nan)

	//boolean
	// var t, f bool = true, false
	// fmt.Println(t, f)

	// //string
	// var s string = "Hello Golang"
	// fmt.Println(s)

	// fmt.Printf("%T\n", s)

	// var si string = "300"
	// fmt.Println(si)
	// fmt.Printf("%T\n", si)

	// fmt.Println(`test
	// test
	// 		test`)

	// fmt.Println("\"")

	// fmt.Println(`"`)

	// fmt.Println(s[0])
	// fmt.Println(string(s[0]))

	//Note:String is set of bytes

	//byte(uint8)
	// byteA := []byte{72, 73}
	// fmt.Println(byteA)

	// fmt.Println(string(byteA))

	// c := []byte("HI")
	// fmt.Println(c)

	// fmt.Println(string(c))

	//Array

	// var arr1 [3]int

	// fmt.Println(arr1)
	// fmt.Printf("%T\n", arr1)

	// var arr2 [3]string = [3]string{"A", "B"}
	// fmt.Println(arr2)

	// arr3 := [3]int{1, 2, 3}
	// fmt.Println(arr3)

	// arr4 := [...]string{"C"}
	// fmt.Println(arr4)
	// fmt.Printf("%T\n", arr4)

	// fmt.Println(arr2[0])
	// fmt.Println(arr2[1])
	// fmt.Println(arr2[2])
	// fmt.Println(arr2[2-1])

	// arr2[2] = "C"
	// fmt.Println(arr2)

	// // var arr5 [4]int
	// // arr5 = arr1
	// // fmt.Println(arr5)

	// fmt.Println(len(arr1))
	// fmt.Println(len(arr2))
	// fmt.Println(len(arr3))
	// fmt.Println(len(arr4))

	//Note; 配列型は要素数を変更できない。スライス型は要素数を変更可能

	//Interface & nil
	//Note:interfaceは互換性はあるが、数式計算ができない。
	// var x interface{}
	// fmt.Println(x)
	// x = 1

	// fmt.Println(x)

	// x = 3.14
	// fmt.Println(x)

	// x = "A"
	// fmt.Println(x)

	// x = [3]int{1, 2, 3}
	// fmt.Println(x)

	// // x = 2
	// // fmt.Println(x + 3)

	//型変換

	// var i int = 1
	// fl64 := float64(i)
	// fmt.Println(fl64)
	// fmt.Printf("i= %T\n", i)
	// fmt.Printf("fl64= %T\n", fl64)

	// i2 := int(fl64)
	// fmt.Printf("i2= %T\n", i2)

	// fl := 10.5
	// i3 := int(fl)
	// fmt.Printf("i3= %T\n", i3)
	// fmt.Println(i3)

	// var s string = "100"
	// fmt.Printf("s = %T\n", s)

	// i, err := strconv.Atoi(s)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(i)
	// fmt.Printf("i= %T\n", i)

	// var i2 int = 200
	// s2 := strconv.Itoa(i2)
	// fmt.Println(s2)
	// fmt.Printf("s2= %T\n", s2)

	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)

}
