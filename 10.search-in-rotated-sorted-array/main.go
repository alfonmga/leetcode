package search_in_rotated_sorted_array

func Search(nums []int, target int) int {
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] == target {
	// 		return i
	// 	}
	// }
	// return -1
	return idxBinarySearch(nums, target)
}

func idxBinarySearch(nums []int, target int) int {
	lo := 0
	hi := len(nums) - 1

	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[lo] <= nums[mid] {
			if target >= nums[lo] && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	if nums[lo] == target {
		return lo
	} else {
		return -1
	}
}
