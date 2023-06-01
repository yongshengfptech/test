package main

import "fmt"

func main() {
	fmt.Printf("1 + 2 = %v\n", sum(1, 2))
}

func sum(a, b int) int {
	return a + b
}
