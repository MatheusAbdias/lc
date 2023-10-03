package goodpairs

func numIdenticalPairs(nums []int) int {
	appear := map[int]int{}
	var pair int
	for _, value := range nums {
		pair += appear[value]
		appear[value]++
	}

	return pair
}
