package main

import "fmt"

// // 型スイッチ
// func anything(a interface{}) {
// 	// fmt.Println(a)
// 	switch v := a.(type) {
// 	case string:
// 		fmt.Println(v + "!?")
// 	case int:
// 		fmt.Println(v + 10000)
// 	}
// }

// //defer
// func TestDefer() {
// 	defer fmt.Println("END")
// 	fmt.Println("START")
// }

// // Note: deferがついたものは、後ろから実行される。この場合、６、５、４の順番に出力される。
// func RunDefer() {
// 	defer fmt.Println("4")
// 	defer fmt.Println("5")
// 	defer fmt.Println("6")
// }

// // fo foroutin(Goの並列処理)

// func sub() {
// 	for {
// 		fmt.Println("Sub Loop")
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }

// // init
func init() {
	fmt.Println("init")
	fmt.Println("init2")
}
func init() {
	fmt.Println("init2")
}

func main() {

	fmt.Println("Main")

	// go sub()
	// go sub()
	// for {
	// 	fmt.Println("Main Loop")
	// 	time.Sleep(200 * time.Millisecond)
	// }

	// //panic recover
	// //panicが起これば、recoverが値を返して、xに値が入る。そのxをPrintlnでプリントしている。
	// //推奨はエラーハンドリングを使う方が推奨
	// defer func() {
	// 	if x := recover(); x != nil {
	// 		fmt.Println(x)
	// 	}
	// }()
	// panic("Runtime error")
	// fmt.Println("Start")

	// TestDefer()

	// // func() {
	// // 	fmt.Println("1")
	// // 	fmt.Println("2")
	// // 	fmt.Println("3")
	// // }()
	// //Note: defer をつけると、main()が終わって、最後にdefer がついている関数を実行する。
	// defer func() {
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// 	fmt.Println("3")
	// }()
	// fmt.Println("TEST")

	// RunDefer()

	// //ファイルのリソースの解放処理を実行している。
	// file, err := os.Create("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()

	// file.Write([]byte("Hello"))

	//if文
	// a := 0
	// if a == 2 {
	// 	fmt.Println("two")
	// } else if a == 1 {
	// 	fmt.Println("one")
	// } else {
	// 	fmt.Println("I don't know")
	// }
	// if b := 100; b == 100 {
	// 	fmt.Println("one hundred")
	// }

	// x := 0
	// if x := 2; true {
	// 	fmt.Println(x)
	// }
	// fmt.Println(x)

	//エラーハンドリング
	// var s string = "100"
	// i, err := strconv.Atoi(s)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("i= %T\n", i)

	//for文
	// g := 0
	// for {
	// 	g++
	// 	if g == 3 {
	// 		break
	// 	}
	// 	fmt.Println("loop")
	// }

	//条件for
	// point := 0
	// for point < 10 {
	// 	fmt.Println(point)
	// 	point++
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	if i == 6 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	//配列を使ったfor
	// arr := [3]int{1, 2, 3}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// arr := [3]int{1, 2, 3}

	// for i := range arr {
	// 	fmt.Println(i)
	// }

	// sl := []string{"Python", "PHP", "GO"}
	// for i, v := range sl {
	// 	fmt.Println(i, v)
	// }

	// m := map[string]int{"apple": 100, "banana": 200}

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	//Switch
	// n := 5
	// switch n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("I dont know")
	// }

	// switch n := 2; n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")

	// default:
	// 	fmt.Println("I dont know")
	// }

	// k := 6
	// switch {
	// case k > 0 && k < 4:
	// 	fmt.Println("0<n<4")
	// case k > 3 && k < 7:
	// 	fmt.Println("3<n<7")
	// }

	// anything("aaa")
	// anything(1)

	// var x interface{} = 3

	// i, isInt := x.(int)
	// fmt.Println(i, isInt)
	// // fmt.Println(x + 2)

	// f, isFloat64 := x.(float64)
	// fmt.Println(f, isFloat64)

	// if x == nil {
	// 	fmt.Println("None")
	// } else if i, isInt := x.(int); isInt {
	// 	fmt.Println(i, "x is Int")
	// } else if s, isString := x.(string); isString {
	// 	fmt.Println(s, isString)
	// } else {
	// 	fmt.Println("I don't know")
	// }

	// switch x.(type) {
	// case int:
	// 	fmt.Println("int")
	// case string:
	// 	fmt.Println("string")
	// default:
	// 	fmt.Println("I dont know")
	// }

	// switch v := x.(type) {
	// case bool:
	// 	fmt.Println(v, "bool")
	// case int:
	// 	fmt.Println(v, "int")
	// case string:
	// 	fmt.Println(v, "string")
	// }

	//ラベル付きfor
	// Loop:
	// 	for {
	// 		for {
	// 			for {
	// 				fmt.Println("START")
	// 				break Loop
	// 			}
	// 			fmt.Println("No operation")
	// 		}
	// 		fmt.Println("No operation")
	// 	}
	// 	fmt.Println("END")

	//ラベル付きfor
	// Loop:
	// 	for i := 0; i < 3; i++ {
	// 		for j := 1; j < 3; j++ {
	// 			if j > 1 {
	// 				continue Loop
	// 			}
	// 			fmt.Println(i, j, i*j)
	// 		}
	// 		fmt.Println("No operation")
	// 	}

}
