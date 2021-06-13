package interation

func Repeat(character string, repeatedCount int) string {
	var res string
	for i := 0; i < repeatedCount; i++ {
		res += character

	}

	return res
}
