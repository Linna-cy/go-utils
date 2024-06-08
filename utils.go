package utils

func Max(nums ...int) (max int) {
	for i, max := 1, nums[0]; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return
}
