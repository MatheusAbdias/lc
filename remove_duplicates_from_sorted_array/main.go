package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	index := 1
	for index < len(nums) {
		if nums[index-1] == nums[index] {
			nums = append(nums[:index], nums[index+1:]...)
		} else {
			index++
		}
	}
	return index
}
