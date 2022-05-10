package rest

import "testing"

func TestAddBlacklistElement(t *testing.T) {

	result := numberDosentExistDeleteNumber(5, 5)

	if result == nil {
		t.Logf("AddBlacklistElement(3,4) PASSED")

	} else {
		t.Errorf("AddBlacklistElement(3,4) FAILED.")
	}

}
