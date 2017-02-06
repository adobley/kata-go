package factorprime

func Factor(num int) []int {
	result := []int{}

	if num%2 == 0 {
		result = append(result, 2)
		num /= 2
	}

	if num != 1 {
		result = append(result, num)
	}

	return result
}
