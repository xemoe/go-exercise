package conditions

import "testing"

func TestIsOdd(t *testing.T) {
	type testCase struct {
		arg1 int
		want bool
	}

	cases := []testCase{
		{0, false},
		{1, true},
		{2, false},
		{3, true},
	}

	for _, tc := range cases {
		got := IsOdd(tc.arg1)
		if tc.want != got {
			t.Errorf("Expected %t, but got %t", tc.want, got)
		}
	}
}
