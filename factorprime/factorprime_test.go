package factorprime

import "testing"

func TestFactorPrime(t *testing.T) {
	t.Log("Returns an empty list when factoring one")
	result := Factor(1)
	if len(result) != 0 {
		t.Errorf("Did not return an empty list!\n"+
			"Actual: %d", result)
	}
}
