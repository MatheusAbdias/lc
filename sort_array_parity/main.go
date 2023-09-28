package sortarrayparity

func sortArrayByParity(nums []int) []int {
	even := []int{}
	odd := []int{}

	for index := 0; index < len(nums); index++ {
		if nums[index]%2 == 0 {
			even = append(even, nums[index])
		} else {
			odd = append(odd, nums[index])
		}
	}
	return append(even, odd...)
}
