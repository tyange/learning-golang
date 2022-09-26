package main

import "fmt"

type hotdog int

var a hotdog
var b int

func main() {
	a = 42
	b = 40
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	// 타입 변환으로 다른 타입을 가진 변수의 타입을 변경할 수 있다
	a = hotdog(b)
	b = int(a)
}
