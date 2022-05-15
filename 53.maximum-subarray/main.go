package maximum_subarray

func MaxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// fmt.Printf("-- %v --\n", nums)
	maxSubArrSum := 0
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= len(nums[i:]); j++ {
			_subArr := nums[i : i+j]
			_sum := 0
			for k := 0; k < len(_subArr); k++ {
				_sum += _subArr[k]
			}
			// fmt.Println(i, j, _subArr, _sum)
			if _sum > maxSubArrSum {
				maxSubArrSum = _sum
			}
		}
	}

	return maxSubArrSum
}
