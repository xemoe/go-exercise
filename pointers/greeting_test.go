package pointers

import "testing"

func TestAppendGreeting(t *testing.T) {
	type testCase struct {
		arg1 string
		want string
	}

	cases := []testCase{
		{"Bob", "Hi, Bob"},
	}

	for _, tc := range cases {
		got := tc.arg1
		AppendGreeting(&got)
		if tc.want != got {
			t.Errorf("Expected %s, but got %s", tc.want, got)
		}
	}
}
