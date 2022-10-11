package loops

import "testing"

func TestIsPrime(t *testing.T) {
	type testCase struct {
		arg1 int
		want bool
	}
	cases := []testCase{
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
	}

	for _, tc := range cases {
		got := IsPrime(tc.arg1)
		if tc.want != got {
			t.Errorf("Expected %t, but got %t", tc.want, got)
		}
	}
}
