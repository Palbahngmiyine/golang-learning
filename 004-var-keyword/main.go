package main

import (
	"fmt"
)

var y = 43

// 아무 값도 없이 변수를 지정할 경우 꼭 타입을 명시해줘야 함!
var z int

func main() {
	x := 42
	fmt.Println(x)
	foo()
	y = 3
	foo()
	fmt.Println(z)
}

func foo() {
	fmt.Println("Declared:", y)
}
