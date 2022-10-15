package loops

import "testing"

func TestSumOfFirst(t *testing.T) {
	type testCase struct {
		arg1 int
		want int
	}

	cases := []testCase{
		{0, 0},
		{1, 1},
		{2, 3},
		{3, 6},
	}

	for _, tc := range cases {
		got := SumOfFirst(tc.arg1, 0)
		if tc.want != got {
			t.Errorf("Expected %d, but got %d", tc.want, got)
		}
	}
}
