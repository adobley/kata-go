package factorprime

import "testing"

func Test_FactorPrime_ReturnsEmptySlice_WhenFactoringOne(t *testing.T) {
	t.Log("Returns [] when factoring 1")
	result := Factor(1)
	assertSlicesEqual(t, []int{}, result)
}

func Test_FactorPrime_ReturnsSliceWithTwo_WhenFactoringTwo(t *testing.T) {
	t.Log("Returns [2] when factoring 2")
	result := Factor(2)
	assertSlicesEqual(t, []int{2}, result)
}

func Test_FactorPrime_ReturnsSliceWithThree_WhenFactoringThree(t *testing.T) {
	t.Log("Returns [3] when factoring 3")
	result := Factor(3)
	assertSlicesEqual(t, []int{3}, result)
}

func Test_FactorPrime_ReturnsSliceWithTwoTwos_WhenFactoringFour(t *testing.T) {
	t.Log("Returns [2, 2] when factoring 4")
	result := Factor(4)
	assertSlicesEqual(t, []int{2, 2}, result)
}

func Test_FactorPrime_ReturnsSliceWithFive_WhenFactoringFive(t *testing.T) {
	t.Log("Returns [5] when factoring 5")
	result := Factor(5)
	assertSlicesEqual(t, []int{5}, result)
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
		"Expected: %d\n"+
		"Actual: %d",
		expected,
		actual)
}
