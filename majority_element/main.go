package majorityelement

func majorityElement(nums []int) []int {
	ocurrencies := map[int][]int{}
	result := []int{}

	for index, number := range nums {
		ocurrencies[number] = append(ocurrencies[number], index)
	}

	for k, v := range ocurrencies {
		if len(v) > len(nums)/3 {
			result = append(result, k)
		}
	}

	return result
}
