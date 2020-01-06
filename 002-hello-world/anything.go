package main

import "fmt"

func main() {
	fmt.Println("Sometimes need to rest within a life")
	foo()
	fmt.Println("Something more")

	for i := 0; i < 100; i++ {
		if i != 0 && i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm foo")
}

func bar() {
	fmt.Println("and I'm bar")
}
