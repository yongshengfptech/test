package generics

func Sum(num1, num2 int) int {
	return num1 + num2
}

// add all the numbers of a slice and return total
func SumNums[numType int64 | float64](nums []numType) numType {
	var total numType

	for _, num := range nums {
		total += num
	}

	return total
}
