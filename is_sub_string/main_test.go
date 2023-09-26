package issubstring

import "testing"

func TestIsSubString(t *testing.T) {
	s := "abc"
	raw := "ahbgdc"

	value := isSubsequence(s, raw)

	if !value {
		t.Error("Expected is sub")
	}
}
