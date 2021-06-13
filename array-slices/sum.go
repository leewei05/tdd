package sum

func Sum(num [5]int) int {
	var sum int
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}

	return sum
}
