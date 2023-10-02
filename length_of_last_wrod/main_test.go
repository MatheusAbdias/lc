package lengthoflastwrod

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	s := "   fly me   to   the moon  "
	expected := 4

	output := lengthOfLastWord(s)

	if expected != output {
		t.Errorf("Expected %v, but receive %v", expected, output)
	}
}
