package duplicateletters

import "testing"

func TestRemoveDubplicateString(t *testing.T) {
	s := "cbabc"
	expected := "abc"
	result := removeDuplicateLetters(s)
	if result != expected {
		t.Errorf("Expected %s, receive %s", expected, result)
	}
}
