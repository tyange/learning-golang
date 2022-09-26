package main

import (
	"fmt"
)

// var 키워드로 함수 패키지 내부에서 선언된 변수는 전역 변수가 된다.

var y = 42

// 변수명 옆에 타입명을 지정하여 변수를 선언할 수 있다.

var z string = "Shaken, not stirred"

// 백틱을 사용하여 띄어쓰기와 따옴표 등 줄바꿈과 기호를 포함한 문자열을 선언할 수 있다.

var a string = `James said, 
"Shaken, 

not stirred"`

// go는 정적 타입 언어이다.

func main() {
	fmt.Println(y)
	// '%T\n' 키워드를 사용하면 해당 변수의 타입을 출력한다.
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	// z = 43
	// fmt.Println(z)
	// fmt.Printf("%T\n", z)
}
