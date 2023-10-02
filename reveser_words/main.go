package reveserwords

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")

	for index, word := range words {
		words[index] = reverseWord(word)
	}
	return strings.Join(words, " ")
}

func reverseWord(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
