package conditions

import "testing"

func isOdd(n int) bool {
	return false
}

func TestIsOdd(t *testing.T) {
	expect := true
	actual := isOdd(1)

	if actual != expect {
		t.Errorf("Expect %t, got %t", expect, actual)
	}
}
