package sumall

import (
	"learngo/sum"
)

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, sum.Sum(numbers))
	}

	return sums
}