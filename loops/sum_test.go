package loops

import "testing"

type testCase struct {
	arg1 int
	want int
}

func TestSumOfFirst(t *testing.T) {
	cases := []testCase{
		{0, 0},
		{1, 1},
		{2, 3},
		{3, 6},
	}

	for _, tc := range cases {
		got := SumOfFirst(tc.arg1)
		if tc.want != got {
			t.Errorf("Expected %d, but got %d", tc.want, got)
		}
	}
}
