package palindromonumber

import "fmt"

func isPalindrome(x int) bool {
	s := fmt.Sprintf("%d", x)
	numberLen := len(s)
	for index := 0; index < numberLen; index++ {
		if numberLen%2 == 0 && index == numberLen/2 {
			break
		}
		if s[index] != s[(numberLen-1)-index] {
			return false
		}
	}

	return true
}
