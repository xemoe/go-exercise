package conditions

import "testing"

type testCase struct {
	arg1 int
	want bool
}

func TestIsOdd(t *testing.T) {
	cases := []testCase{
		{0, false},
		{1, true},
		{2, false},
		{3, true}}

	for _, tc := range cases {
		got := IsOdd(tc.arg1)
		if tc.want != got {
			t.Errorf("Expected %t, but got %t", tc.want, got)
		}
	}
}
