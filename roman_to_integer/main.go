package romantointeger

func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	number := 0

	for index := 0; index < len(s); index++ {
		currentValue := romanMap[s[index]]
		if index == len(s)-1 {
			number += currentValue
		} else {
			nextValue := romanMap[s[index+1]]
			if nextValue > currentValue {
				number -= currentValue
			} else {
				number += currentValue
			}
		}

	}
	return number
}
