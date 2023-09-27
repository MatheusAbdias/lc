package palindromonumber

import (
	"testing"
)

func TestIsPalindromo(t *testing.T) {
	x := -101
	expected := true

	value := isPalindrome(x)

	if expected != value {
		t.Errorf("Expected %v receive %v", expected, value)
	}
}
