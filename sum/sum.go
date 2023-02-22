package sum

func Sum(numbs []int) int {

	sum := 0

	for _, number := range numbs {
		sum += number
	}

	return sum
}