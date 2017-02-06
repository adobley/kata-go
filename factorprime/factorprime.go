package factorprime

func Factor(num int) []int {
	if num != 1 {
		return []int{num}
	}
	return []int{}
}
