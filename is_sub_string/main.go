package issubstring

func isSubsequence(s string, t string) bool {
	for subIndex, sub := range s {
		for index, b := range t {
			if b == sub {
				return isSubsequence(s[subIndex+1:], t[index+1:])
			}
		}
		return false
	}
	return true
}
