package main

import (
	"fmt"
	"project/test/mathQuestion"
)

func main() {
	QandA()
}

func QandA() {
	q := mathQuestion.RandomQuestion()
	fmt.Println("\nEnter -1 to exit.")
	fmt.Printf("%v %c %v = ", q.Num1, q.Operator, q.Num2)

	var ans int
	_, err := fmt.Scanf("%d", &ans)
	if err == nil {
		if ans == -1 {
			return
		}
		fmt.Printf("%s\n", mathQuestion.CheckAns(ans, q))
	} else {
		fmt.Println("Not a number")
	}
	QandA()
}
