package main

import (
	"fmt"

	"project/test/generics"

	"rsc.io/quote"
)

func main() {
	fmt.Printf("1 + 2 = %v\n", generics.Sum(1, 2))
	fmt.Println(quote.Hello())
}
