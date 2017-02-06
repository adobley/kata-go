package factorprime

import "testing"

func TestFactorPrimeReturnsEmptySliceWhenFactoringOne(t *testing.T) {
	t.Log("Returns an empty slice when factoring one")
	result := Factor(1)
	if len(result) != 0 {
		t.Errorf("Did not return an empty slice!\n"+
			"Actual: %d", result)
	}
}
