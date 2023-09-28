package romantointeger

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	s := "MCMXCIV"
	fmt.Printf("%v", romanToInt(s))
}
