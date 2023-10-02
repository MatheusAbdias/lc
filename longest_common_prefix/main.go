package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	var smal string
	var smalIndex int
	if len(strs) == 1 {
		return strs[0]
	}
	for index, s := range strs {
		if index == 0 {
			smal = s
			smalIndex = index
		} else {
			if len(s) < len(smal) {
				smal = s
				smalIndex = index
			}
		}
	}
	strs = append(strs[:smalIndex], strs[smalIndex+1:]...)
	for smal != "" {
	out:
		for index, s := range strs {
			if s[:len(smal)] != smal {
				smal = smal[:len(smal)-1]
				break out
			} else {
				if index == len(strs)-1 {
					return smal
				}

			}

		}
	}
	return smal
}
