package lengthoflastwrod

import "strings"

func lengthOfLastWord(s string) int {
	words := strings.Split(s, " ")
	lastWord := words[0]
	for index := 1; index < len(words); index++ {
		if len(SpaceFieldsJoin(words[index])) != 0 {
			lastWord = words[index]
		}
	}
	return len(lastWord)
}

func SpaceFieldsJoin(str string) string {
	return strings.Join(strings.Fields(str), "")
}
