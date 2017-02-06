package factorprime

import "testing"

func TestFactorPrimeReturnsEmptySliceWhenFactoringOne(t *testing.T) {
	t.Log("Returns an empty slice when factoring 1")
	result := Factor(1)
	if len(result) != 0 {
		t.Errorf("Did not return an empty slice!\n"+
			"Actual: %d", result)
	}
}

func TestFactorPrimeReturnsSliceWithTwoWhenFactoringTwo(t *testing.T) {
	t.Log("Returns a slice containg only 2 when factoring 2")
	result := Factor(2)
	if len(result) != 1 || result[0] != 2 {
		t.Errorf("Did not return an slice containing only 2!\n"+
			"Expected: %d\n"+
			"Actual: %d",
			[]int{2},
			result)
	}
}
