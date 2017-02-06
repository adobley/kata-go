package factorprime

import "testing"

func TestFactorPrimeReturnsEmptySliceWhenFactoringOne(t *testing.T) {
	t.Log("Returns an empty slice when factoring 1")
	result := Factor(1)
	assertSlicesEqual(t, []int{}, result)
}

func TestFactorPrimeReturnsSliceWithTwoWhenFactoringTwo(t *testing.T) {
	t.Log("Returns a slice containg only 2 when factoring 2")
	result := Factor(2)
	assertSlicesEqual(t, []int{2}, result)
}

func TestFactorPrimeReturnsSliceWithThreeWhenFactoringThree(t *testing.T) {
	t.Log("Returns a slice containg only 3 when factoring 3")
	result := Factor(3)
	assertSlicesEqual(t, []int{3}, result)
}

func assertSlicesEqual(t *testing.T, expected, actual []int) {
	if len(expected) != len(actual) {
		errorSlicesNotEqual(t, expected, actual)
	}
	for i := range actual {
		if actual[i] != expected[i] {
			errorSlicesNotEqual(t, expected, actual)
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
