package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	// for文
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while文
	n := 0
	for n < 10 {
		fmt.Println(n)
	}
}
