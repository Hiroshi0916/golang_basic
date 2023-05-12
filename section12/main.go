package main

import "fmt"

// //interface
// //違う型をまとめるもの

// type Stringfy interface {
// 	ToString() string
// }

// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p *Person) ToString() string {
// 	return fmt.Sprintf("Name=%V, Age=%v", p.Name, p.Age)
// }

// type Car struct {
// 	Number string
// 	Model  string
// }

// func (c *Car) ToString() string {
// 	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
// }

// func main() {

// 	//異なる型をまとめられる。
// 	vs := []Stringfy{
// 		&Person{Name: "Taro", Age: 21},
// 		&Car{Number: "123-456", Model: "AB-1234"},
// 	}

// 	for _, v := range vs {
// 		fmt.Println(v.ToString())
// 	}
// }

//interface
//カスタムエラー

// type error interface {
// 	Error() string
// }

// type MyError struct {
// 	Message string
// 	ErrCode int
// }

// func (e *MyError) Error() string {
// 	return e.Message
// }

// func RaiseError() error {
// 	return &MyError{Message: "カスタムエラーが発生しました", ErrCode: 1234}
// }

// func main() {
// 	err := RaiseError()

// 	fmt.Println(err.Error())

// 	e, ok := err.(*MyError)
// 	if ok {
// 		fmt.Println(e.ErrCode)
// 	}

// }

//Stringerのパッケージの定義
// type Stringer interface {
// 	String() string
// }

type Point struct {
	A int
	B string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

func main() {
	p := &Point{100, "ABC"}
	fmt.Println(p)

}
