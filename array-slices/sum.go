package sum

func Sum(num []int) int {
	var sum int
	for _, number := range num {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfSlices := len(numbersToSum)
	sums := make([]int, lengthOfSlices)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
