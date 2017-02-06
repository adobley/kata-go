package factorprime

func Factor(num int) []int {
	result := []int{}
	divisor := 2

	for num != 1 {
		for num%divisor == 0 {
			result = append(result, divisor)
			num /= divisor
		}
		divisor++
	}

	return result
}
