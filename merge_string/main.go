package mergestring

func mergeAlternately(word1 string, word2 string) string {
	result := ""
	appended := 0
	for index, b := range word1 {
		result += string(b)
		if index < len(word2)-1 {
			appended++
			result += string(word2[index])
		}
	}

	for index := appended; index < len(word2)-1; index++ {
		result = string(word2[index])
	}

	return result
}
