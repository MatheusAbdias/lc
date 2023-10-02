package colorsneithbors

import "testing"

func TestWinnerGame(t *testing.T) {
	testCases := []struct {
		Colors   string
		Expected bool
	}{
		{Colors: "AAABABB", Expected: true},
		{Colors: "AA", Expected: false},
		{Colors: "ABBBBBBBAAA", Expected: false},
	}

	for _, tc := range testCases {
		output := winnerOfGame(tc.Colors)
		if output != tc.Expected {
			t.Errorf("Exepected %v, output %v", tc.Expected, output)
		}
	}

}
