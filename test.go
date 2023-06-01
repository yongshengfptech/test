package main

import "fmt"

func main() {
	fmt.printf("1 + 2 = %v\n", sum(1, 2))
}

func sum(a, b int) int {
	return a + b
}
