package sum

func Sum(num []int) int {
	var sum int
	for _, number := range num {
		sum += number
	}

	return sum
}
