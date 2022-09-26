package main

import "fmt"

// 문자열 값의 zero value는 빈 문자열이다. ("")
var y string

// 문자열 값의 zero value는 0.
var z int

func main() {
	fmt.Printf("%T", y)
	fmt.Printf("%T", z)
}
