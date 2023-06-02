package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Printf("1 + 2 = %v\n", sum(1, 2))
	fmt.Println(quote.Hello())
}

func Sum(a, b int) int {
	return a + b
}
