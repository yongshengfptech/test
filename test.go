package main

import (
	"fmt"
	"project/test/mathQuestion"
)

func main() {
	q := mathQuestion.RandomQuestion()
	fmt.Printf("%v %c %v = ", q.Num1, q.Operator, q.Num2)

	var ans int
	_, err := fmt.Scanf("%d", &ans)
	if err != nil {
		fmt.Println("\nNot a number")
	} else {
		fmt.Printf("\n%s", mathQuestion.CheckAns(ans, q))
	}
}
