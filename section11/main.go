package main

import (
	"fmt"
)

// //構造体はJavaのClassのようなもの。

// type User struct {
// 	Name string
// 	Age  int
// 	// X, Y int
// }

// func UpdateUser(user User) {
// 	user.Name = "A"
// 	user.Age = 1000
// }

// func Updateuser2(user *User) {
// 	user.Name = "A"
// 	user.Age = 1000
// }

// func main() {
// 	var user1 User
// 	fmt.Println(user1)
// 	user1.Name = "user1"
// 	user1.Age = 10
// 	fmt.Println(user1)

// 	user2 := User{}
// 	fmt.Println(user2)
// 	user2.Name = "user2"
// 	fmt.Println(user2)

// 	user3 := User{Name: "user3", Age: 30}
// 	fmt.Println(user3)

// 	user4 := User{"user4", 40}
// 	fmt.Println(user4)

// 	user6 := User{Name: "user6"}
// 	fmt.Println(user6)

// 	//newは構造体のポインタ型を返す
// 	user7 := new(User)
// 	fmt.Println(user7)

// 	user8 := &User{}
// 	fmt.Println(user8)

// 	UpdateUser(user1)
// 	// UpdateUser2(user8)

// 	fmt.Println(user1)
// 	fmt.Println(user8)

// }

// type User struct {
// 	Name string
// 	Age  int
// }

// func (u User) SayName() {
// 	fmt.Println(u.Name)
// }

// func (u User) SetName(name string) {
// 	u.Name = name
// }

// // 原則:メソッドのレシーバーはポインタ型にする。
// // ->呼び出し元はポインタ型。呼び出し先はポインタ型宣言無し。
// func (u *User) SetName2(name string) {
// 	u.Name = name
// }

// func main() {
// 	user1 := User{Name: "user1"}
// 	user1.SayName()

// 	user1.SetName("A")
// 	user1.SayName()

// 	user1.SetName2("A")
// 	user1.SayName()

// 	user2 := &User{Name: "user2"}
// 	user2.SetName2("B")
// 	user2.SayName()

// }

// type T struct {
// 	User User
// }

// type User struct {
// 	Name string
// 	Age  int
// }

// func (u *User) SetName() {
// 	u.Name = "A"
// }

// // Struct 埋め込み
// func main() {

// 	t := T{User: User{Name: "user1", Age: 10}}

// 	fmt.Println(t)

// 	fmt.Println(t.User)
// 	fmt.Println(t.User.Name)

// 	//構造体が省略した場合のみアクセス可能
// 	// fmt.Println(t.Name)

// 	t.User.SetName()
// 	fmt.Println(t)
// }

// type User struct {
// 	Name string
// 	Age  int
// }

// // Struct型コンストラクタ
// func NewUser(name string, age int) *User {
// 	return &User{Name: name, Age: age}
// }

// func main() {
// 	user1 := NewUser("user1", 10)

// 	fmt.Println(user1)

// 	fmt.Println(*user1)
// }

// type User struct {
// 	Name string
// 	Age  int
// }

// // ポインタのスライスで定義する方が、望ましい。
// // ->データ管理+ メモリ効率
// type Users []*User

// func main() {

// 	user1 := User{Name: "user1", Age: 10}
// 	user2 := User{Name: "user2", Age: 20}

// 	user3 := User{Name: "user3", Age: 30}
// 	user4 := User{Name: "user4", Age: 40}

// 	//構造体スライス
// 	users := Users{}

// 	users = append(users, &user1)
// 	users = append(users, &user2)

// 	users = append(users, &user3, &user4)

// 	for _, u := range users {
// 		fmt.Println(*u)
// 	}

// 	//make関数を使って、構造体のスライスを作ることができる。

// 	users2 := make([]*User, 0)
// 	users2 = append(users2, &user1, &user2)

// 	for _, u := range users2 {
// 		fmt.Println(*u)
// 	}

// }

// type User struct {
// 	Name string
// 	Age  int
// }

// func main() {

// 	//struct map
// 	m := map[int]User{
// 		1: {Name: "user1", Age: 10},
// 		2: {Name: "user2", Age: 20},
// 	}

// 	fmt.Println(m)

// 	m2 := map[User]string{
// 		{Name: "user1", Age: 10}: "Tokyo",
// 		{Name: "user2", Age: 20}: "LA",
// 	}

// 	fmt.Println(m2)

// 	m3 := make(map[int]User)
// 	fmt.Println(m3)

// 	m3[1] = User{Name: "user3"}
// 	fmt.Println(m3)

// 	for _, v := range m {
// 		fmt.Println(v)
// 	}
// }

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
