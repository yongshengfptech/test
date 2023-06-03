package mathQuestion

import (
	"fmt"
	"math/rand"
)

type question struct {
	Num1     int
	Num2     int
	Operator rune
	Ans      int
}

func RandomQuestion() question {
	num1 := rand.Intn(50)
	num2 := rand.Intn(50)

	operators := []rune{'+', '-'}
	operator := operators[rand.Intn(len(operators))]

	var ans int
	if operator == '+' {
		ans = num1 + num2
	} else if operator == '-' {
		if num1 < num2 {
			num1, num2 = num2, num1
		}
		ans = num1 - num2
	}

	return question{Num1: num1, Num2: num2, Operator: operator, Ans: ans}
}

func CheckAns(ans int, q question) string {
	if ans == q.Ans {
		return fmt.Sprintln("Correct Answer")
	}
	return fmt.Sprintln("Wrong Answer")
}
