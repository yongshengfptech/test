package main

import (
	"fmt"
	"project/test/mathQuestion"
	"strconv"
)

func main() {
	QandA()
}

func QandA() {
	q := mathQuestion.RandomQuestion()
	fmt.Println("Enter -1 to exit.")
	fmt.Printf("%v %c %v = ", q.Num1, q.Operator, q.Num2)

	var ansString string
	fmt.Scanln(&ansString)

	ans, err := strconv.Atoi(ansString)
	if err == nil {
		if ans == -1 {
			return
		}
		fmt.Printf("%s\n", mathQuestion.CheckAns(ans, q))
	} else {
		fmt.Printf("\033[31mNot a number\033[0m\n\n")
	}
	QandA()
}
