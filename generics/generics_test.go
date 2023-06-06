package generics

import (
	"math"
	"testing"
)

func TestSum(t *testing.T) {
	num1, num2 := 1, 2
	want := 3
	total := Sum(num1, num2)

	if want != total {
		t.Fatalf(`Sum(%v, %v) = %v, want %v`, num1, num2, total, want)
	}
}

func TestSumNums(t *testing.T) {
	// for int64
	numsInt64 := []int64{1, 2}
	wantInt64 := int64(3)
	totalInt64 := SumNums(numsInt64)

	if totalInt64 != wantInt64 {
		t.Fatalf(`SumNums(%v) = %v, want %v`, numsInt64, totalInt64, wantInt64)
	}

	// for float64
	numsfloat64 := []float64{1.1, 2.2}
	wantfloat64 := float64(3.3)
	totalfloat64 := SumNums(numsfloat64)
	// round off to 2 decimal place
	totalfloat64 = math.Round(totalfloat64*100) / 100

	if totalfloat64 != wantfloat64 {
		t.Fatalf(`SumNums(%v) = %v, want %v`, numsfloat64, totalfloat64, wantfloat64)
	}
}
