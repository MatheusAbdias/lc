package removeelement

func removeElement(nums []int, val int) int {
	index := 0

	for index < len(nums) {
		if nums[index] == val {
			nums = append(nums[:index], nums[index+1:]...)
		} else {
			index++
		}
	}

	return index
}
