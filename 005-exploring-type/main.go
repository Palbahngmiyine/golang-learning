package main

import (
	"fmt"
)

var y = 42

var z string = "흔들렸지만, 동요하진 않았습니다."

// 문자열 리터럴과 문자열의 차이점을 보여줌
var a string = `제임스가 말하길, 
"흔들렸지만, 

동요하진 않았습니다."`

// Golang은 정적 언어이기 때문에 타입을 명시적으로 선언할 수 있음
// 개인적인 생각: 정말 로컬 스코프 안에서 금방 휘발되지 않는 이상 타입을 명시적으로 써주자.

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
