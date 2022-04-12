package two_sum

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TwoSum(nums []int, target int) []int {
	for i, v := range nums {
		subSliceStartIdx := i + 1
		if len(nums) <= subSliceStartIdx {
			break
		}
		for idx, val := range nums[subSliceStartIdx:] {
			sum := v + val
			if sum == target {
				return []int{i, subSliceStartIdx + idx}
			}
		}
	}
	return []int{}
}
