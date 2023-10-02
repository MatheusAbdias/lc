package findfirstoccurrence

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			sub := string(haystack[i])
			for j := 1; j < len(needle); j++ {
				if i+j > len(haystack)-1 {
					return -1
				}
				sub = sub + string(haystack[i+j])
			}
			if sub == needle {
				return i
			}
		}
	}
	return -1
}
