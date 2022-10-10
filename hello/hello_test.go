package hello

import "testing"

func TestHello(t *testing.T) {

	expect := "Hello, Foo"
	actual := Hello("Foo")

	if actual != expect {
		t.Errorf("Expect %s, got %v", expect, actual)
	}
}
