package pointers

import "testing"

func Double(n *int) {
	*n = *n * 2
}

func TestDouble(t *testing.T) {
	type testCase struct {
		arg1 int
		want int
	}

	cases := []testCase{
		{0, 0},
		{1, 2},
		{2, 4},
		{3, 6},
	}

	for _, tc := range cases {
		got := tc.arg1
		Double(&got)
		if tc.want != got {
			t.Errorf("Expected %d, but got %d", tc.want, got)
		}
	}
}
