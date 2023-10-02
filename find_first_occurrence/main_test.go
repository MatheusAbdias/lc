package findfirstoccurrence

import "testing"

func TestSubString(t *testing.T) {
	s := "sadbutsad"
	sub := "sad"
	expected := 0
	output := strStr(s, sub)

	if expected != output {
		t.Errorf("Expected %v, but receive %v", expected, output)
	}
}
