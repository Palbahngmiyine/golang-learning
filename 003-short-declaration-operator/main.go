package main

import (
	"fmt"
)

func main() {
	x := 42
	fmt.Println(x)
	x = 99
	fmt.Println(x)
	// 선언을 하고 사용하지 않으면 에러남
	y := 100 + 70
	fmt.Println(y)
	z := "Bond, James"
	fmt.Println(z)
}
