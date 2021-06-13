package interation

const repeatedCount = 5

func Repeat(character string) string {
	res := ""
	for i := 0; i < repeatedCount; i++ {
		res += character

	}

	return res
}
