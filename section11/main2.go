package main

import "fmt"

type T struct {
	User User
}

type User struct {
	Name string
	Age  int
}

func (u *User) SetName() {
	u.Name = "A"
}

// Struct 埋め込み
func main() {

	t := T{User: User{Name: "user1", Age: 10}}

	fmt.Println(t)

	fmt.Println(t.User)
	fmt.Println(t.User.Name)

	//構造体が省略した場合のみアクセス可能
	// fmt.Println(t.Name)

	t.User.SetName()
	fmt.Println(t)
}
