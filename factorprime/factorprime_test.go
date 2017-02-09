package factorprime

import "testing"

func TestFactorPrime(t *testing.T) {
	t.Log("Returns [] when factoring 1")
	assertSlicesEqual(t, []int{}, Factor(1))

	t.Log("Returns [2] when factoring 2")
	assertSlicesEqual(t, []int{2}, Factor(2))

	t.Log("Returns [3] when factoring 3")
	assertSlicesEqual(t, []int{3}, Factor(3))

	t.Log("Returns [2, 2] when factoring 4")
	assertSlicesEqual(t, []int{2, 2}, Factor(4))

	t.Log("Returns [5] when factoring 5")
	assertSlicesEqual(t, []int{5}, Factor(5))

	t.Log("Returns [2, 3] when factoring 6")
	assertSlicesEqual(t, []int{2, 3}, Factor(6))

	t.Log("Returns [7] when factoring 7")
	assertSlicesEqual(t, []int{7}, Factor(7))

	t.Log("Returns [2, 2, 2] when factoring 8")
	assertSlicesEqual(t, []int{2, 2, 2}, Factor(8))

	t.Log("Returns [2, 2, 2, 3, 5, 5, 7, 11, 13, 17, 17] when factoring 2 * 2 * 2 * 3 * 5 * 5 * 7 * 11 * 13 * 17 * 17")
	assertSlicesEqual(t, []int{2, 2, 2, 3, 5, 5, 7, 11, 13, 17, 17}, Factor(2*2*2*3*5*5*7*11*13*17*17))
}

func assertSlicesEqual(t *testing.T, expected, actual []int) {
	if len(expected) != len(actual) {
		errorSlicesNotEqual(t, expected, actual)
		return
	}
	for i := range actual {
		if actual[i] != expected[i] {
			errorSlicesNotEqual(t, expected, actual)
			return
		}
	}

}

func errorSlicesNotEqual(t *testing.T, expected, actual []int) {
	t.Errorf("Did not return expected slice:\n"+
		"\t\tExpected: %d\n"+
		"\t\tActual: %d",
		expected,
		actual)
}
