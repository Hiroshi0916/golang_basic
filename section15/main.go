package main

import (
	"fmt"
	"os"
)

// //OS
// //Java のSystem.exit(0)と同じ
// func main() {
// 	// os.Exit(1)
// 	// fmt.Println("Start")

// 	defer func() {
// 		fmt.Println("defer")
// 	}()
// 	os.Exit(0)
// }

// func main() {
// 	//log.Fatal

// 	_, err := os.Open("A.txt")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

// func main() {
// 	//Args
// 	fmt.Println(os.Args[0])
// 	fmt.Println(os.Args[1])
// 	fmt.Println(os.Args[2])
// 	fmt.Println(os.Args[3])

// 	//os.Argsの要素数
// 	fmt.Printf("length = %d\n", len(os.Args))

// 	//os.Argsの中身の全てを表示
// 	for i, v := range os.Args {
// 		fmt.Println(i, v)
// 	}
// }

// // ファイル操作
// func main() {
// 	f, err := os.Open("test.txt")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer f.Close()
// }

// //ファイル作成
// func main() {
// 	//Create
// 	f, _ := os.Create("foo.txt")
// 	f.Write([]byte("Hello\n"))

// 	f.WriteAt([]byte("Golang"), 6)
// 	f.Seek(0, os.SEEK_END)
// 	f.WriteString("Year")
// }

// ファイル作成
// func main() {
// 	//Read
// 	f, err := os.Open("foo.txt")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	defer f.Close()

// 	bs := make([]byte, 128)

// 	n, err := f.Read(bs)
// 	fmt.Println(n)
// 	fmt.Println(string(bs))

// 	bs2 := make([]byte, 128)

// 	nn, err := f.ReadAt(bs2, 10)
// 	fmt.Println(nn)
// 	fmt.Println(string(bs2))
// }

// func main() {
// 	//ファイル操作
// 	//Open File
// 	//O_RDONLY
// 	//O_WRONLY
// 	//O_RDWR
// 	//O_APPEND
// 	//O_CREATE
// 	//O_TRUNC

// 	f, err := os.OpenFile("foo.txt", os.O_RDONLY|os.O_CREATE, 0666)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	defer f.Close()

// 	bs := make([]byte, 128)

// 	n, err := f.Read(bs)
// 	fmt.Println(n)
// 	fmt.Println(string(bs))

// }

// Time
// func main() {
// 	t := time.Now()
// 	fmt.Println(t)

// 	t2 := time.Date(2020, 6, 10, 10, 10, 10, 0, time.Local)
// 	fmt.Println(t2)
// 	fmt.Println(t2.Year())
// 	fmt.Println(t2.YearDay())
// 	fmt.Println(t2.Month())
// 	fmt.Println(t2.Weekday())
// 	fmt.Println(t2.Day())
// 	fmt.Println(t2.Hour())
// 	fmt.Println(t2.Minute())
// 	fmt.Println(t2.Second())
// 	fmt.Println(t2.Nanosecond())
// 	fmt.Println(t2.Zone())
// }

// Time duration
// func main() {
// 	fmt.Println(time.Hour)
// 	fmt.Printf("%T\n", time.Hour)
// 	fmt.Println(time.Minute)
// 	fmt.Println(time.Second)
// 	fmt.Println(time.Millisecond)
// 	fmt.Println(time.Microsecond)
// 	fmt.Println(time.Nanosecond)

// 	//time.Duration型を文字列から生成する
// 	d, _ := time.ParseDuration("2h30m")
// 	fmt.Println(d)

// 	t3 := time.Now()
// 	t3 = t3.Add(2*time.Minute + 15*time.Second)
// 	fmt.Println(t3)
// 	//現在自国の2分30秒後を表す、time.Time型の取得

// }

// // 時間の差分を取得
// func main() {
// 	t5 := time.Date(2021, 7, 24, 0, 0, 0, 0, time.Local)
// 	t6 := time.Now()
// 	fmt.Println(t6)

// 	//t5-t6
// 	d2 := t5.Sub(t6)
// 	fmt.Println(d2)

// 	//時刻を比較
// 	fmt.Println(t6.Before(t5))
// 	fmt.Println(t6.After(t5))
// 	fmt.Println(t5.Before(t6))
// 	fmt.Println(t5.After(t6))

// }

// // 指定時間のスリープ
// func main() {
// 	// 5秒間の停止
// 	time.Sleep(5 * time.Second)
// 	fmt.Println("5秒停止")

// }

// // Math
// func main() {

// 	fmt.Println(math.Pi)

// 	fmt.Println(math.Sqrt2)

// 	fmt.Println(math.MaxFloat32)
// 	fmt.Println(math.SmallestNonzeroFloat32)
// 	fmt.Println(math.MaxFloat64)
// 	fmt.Println(math.SmallestNonzeroFloat64)
// 	fmt.Println(math.MaxInt64)
// 	fmt.Println(math.MinInt64)
// }

// Math
// //Javaと同じ。
// func main() {
// 	fmt.Println(math.Abs(-100))

// 	fmt.Println(math.Pow(0, 2))

// 	fmt.Println(math.Sqrt(2))

// 	fmt.Println(math.Cbrt(8))

// 	fmt.Println(math.Max(1, 2))
// 	fmt.Println(math.Min(1, 2))
// }

// Math
// Javaと同じ。
// func main() {
// 	fmt.Println(math.Abs(-100))

// 	fmt.Println(math.Pow(0, 2))

// 	fmt.Println(math.Sqrt(2))

// 	fmt.Println(math.Cbrt(8))

// 	fmt.Println(math.Max(1, 2))
// 	fmt.Println(math.Min(1, 2))
// }

// func main() {
// 	//小数点以下の切り捨て、切り上げ
// 	fmt.Println(math.Trunc(3.14))
// 	fmt.Println(math.Trunc(-3.14))

// 	fmt.Println(math.Floor((3.5)))
// 	fmt.Println(math.Floor((-3.5)))

// 	fmt.Println(math.Ceil((3.5)))
// 	fmt.Println(math.Ceil((-3.5)))
// }

// func main() {
// 	//Rand
// 	rand.Seed(256)

// 	fmt.Println(rand.Float64())
// 	fmt.Println(rand.Float64())
// 	fmt.Println(rand.Float64())

// 	fmt.Println(time.Now().UnixNano())
// 	rand.Seed(time.Now().UnixNano())

// 	fmt.Println(rand.Intn(100))
// 	fmt.Println(rand.Int())
// 	fmt.Println(rand.Int31())

// 	fmt.Println(rand.Int63())

// 	fmt.Println(rand.Uint32())

// }

// func main() {
// 	//flag
// 	// go run main.go -n 20 -m message -x
// 	//オプションの値を格納する変数の定義

// 	var (
// 		max int
// 		msg string
// 		x   bool
// 	)
// 	flag.IntVar(&max, "n", 32, "処理数の最大値")

// 	flag.StringVar(&msg, "m", "", "処理メッセージ")

// 	flag.BoolVar(&x, "x", false, "拡張オプション")

// 	//コマンドラインをパース
// 	flag.Parse()

// 	fmt.Println("処理数の最大値 = ", max)
// 	fmt.Println("処理メッセージ = ", msg)
// 	fmt.Println("拡張オプション = ", x)
// }

func main() {
	// // fmt標準
	// fmt.Println("表示")

	// // fmt標準
	// fmt.Print("Hello\n")

	// // 改行
	// fmt.Println("Hello!")

	// fmt.Println("Hello", "world!")
	// fmt.Println("Hello", "world!")

	// fmt.Printf("%s\n", "Hello")
	// fmt.Printf("%#v\n", "Hello")

	fmt.Fprint(os.Stdout, "Hello")
	fmt.Fprintf(os.Stdout, "Hello")
	fmt.Fprintln(os.Stdout, "Hello")

	f, _ := os.Create("test.txt")
	defer f.Close()
	fmt.Fprintln(f, "Fprint")

}
