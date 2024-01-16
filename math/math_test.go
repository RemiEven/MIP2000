package math

import "testing"

func TestDouble(t *testing.T) {
	expected := 4
	actual := Double(2)

	if expected != actual {
		t.Error("test failed, expected double of 2 to be 4")
	}
}
