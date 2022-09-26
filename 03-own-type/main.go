package main

import "fmt"

// type 지정자로 자신만의 타입을 만들 수 있다.
type hotdog int

var a hotdog

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
