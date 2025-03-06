package repeatedElement

func repeatedNTimes(nums []int) int {
	repeatedElements := map[int]int{}
	for i := range nums {
		repeatedElements[nums[i]]++
		if repeatedElements[nums[i]] == len(nums)/2 {
			return nums[i]
		}
	}
	return 0
}