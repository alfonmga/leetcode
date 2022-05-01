package three_sum

import "sort"

// i != j, i != k, and j != k
// nums[i] + nums[j] + nums[k] == 0

func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	triplets := [][]int{}
	for i, iV := range nums {
		for j, jV := range nums {
			if j == i {
				continue
			}
			for k, kV := range nums {
				if k == i || k == j {
					continue
				}

				threeSum := iV + jV + kV
				if threeSum == 0 {
					_threeSumSlice := []int{iV, jV, kV}
					sort.Ints(_threeSumSlice)
					triplets = append(triplets, _threeSumSlice)
				}

			}
		}
	}

	finalTriplets := [][]int{}
	for _, v := range triplets {
		added := false
		for _, jV := range finalTriplets {
			if v[0] == jV[0] && v[1] == jV[1] && v[2] == jV[2] {
				added = true
				break
			}
		}
		if !added {
			finalTriplets = append(finalTriplets, v)
		}
	}

	return finalTriplets
}
