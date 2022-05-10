package rest

import "testing"

func TestNextFibonacci(t *testing.T) {
	element1 := 1
	element2 := 2
	result := Next_fibonacci(&element1, &element2)

	if result != 3 {
		t.Errorf("next_fiboancci(1, 2) FAILED. Expected %d instead of %d/n", 3, result)
	} else {
		t.Logf("next_fibonacci(1, 2) PASSED. Expected %d, got %d", 3, result)
	}
}

type checkTest struct {
	arg1     []int
	arg2     int
	expected []int
}

var checkTests = []checkTest{
	checkTest{[]int{1, 2}, 3, 5},
	checkTest{4, 8, 12},
	checkTest{6, 9, 15},
	checkTest{3, 10, 13},
}

func TestCheckListAndAppend() (t *testing.T) {
	blacklist := []int{2, 5, 8}
	for _, value := range checkTests {
		if output := CheckListAndAppendClone(value.arg1, value.arg2, blacklist); output != value.expected {
			t.Errorf("Output %q not equal to expected %q", output, value.expected)
		}
	}
}
