package main

import (
	"fmt"
	"strconv"
)

type Vector[T any] []T
type IntVector = Vector[int]

type T[A any, B []C, C *A] struct {
	a A
	b B
	c C
}

type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Max[T Number](x, y T) T {
	if x >= y {
		return x
	}
	return y
}

type Stringer interface {
	String() string
}

func f[T fmt.Stringer](xs []T) []string {
	result := []string{}
	for _, x := range xs {
		result = append(result, x.String())
	}
	return result
}

type MyInt int

func (i MyInt) String() string {
	return strconv.Itoa(int(i))
}

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](xs ...T) Set[T] {
	s := make(Set[T])
	for _, v := range xs {
		s[v] = struct{}{}
	}
	return s
}

func (s Set[T]) Add(x T) {
	s[x] = struct{}{}
}

func (s Set[T]) Remove(x T) {
	delete(s, x)
}

func main() {
	fmt.Println(f([]MyInt{1, 2, 3, 4}))
	fmt.Println(Max(1, 2))
	fmt.Println(Max(1.1, 1.3))

	var x, y MyInt = 1, 2
	fmt.Println(Max(x, y))

	var v Vector[int] = Vector[int]{10, 20, 30}
	fmt.Println(v)
	var v2 Vector[float64] = Vector[float64]{1.3, 3.4, 5.6}
	fmt.Println(v2)

	v3 := IntVector{1, 2, 3}
	fmt.Println(v3)

	var t T[int, []*int, *int]
	fmt.Print("A: %T, B: %T, C: %\n", t.a, t.b, t.c)

	var a any = 1
	a = true
	fmt.Println(a)

	s := NewSet(1, 2, 3)
	fmt.Println(s)

	s.Remove(1)
	fmt.Println(s)

}
