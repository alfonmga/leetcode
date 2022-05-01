package three_sum

import "sort"

// i != j, i != k, and j != k
// nums[i] + nums[j] + nums[k] == 0

func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	triplets := [][]int{}
	for i, num := range nums {
		if len(nums) < i+1 {
			break
		}

		jk := []int{}
		for jkIdx := 0; jkIdx < len(nums); jkIdx++ {
			if jkIdx == i {
				continue
			}

			if len(jk) == 0 {
				jk = append(jk, nums[jkIdx])
			} else {
				jk = []int{nums[jkIdx], jk[0]}
			}

			if len(jk) == 2 {
				threeSum := num + jk[0] + jk[1]
				if threeSum == 0 {
					threeSumSlice := []int{num, jk[0], jk[1]}
					sort.Ints(threeSumSlice)
					triplets = append(triplets, threeSumSlice)
				}
			}
		}
	}

	uniqueTriplets := [][]int{}
	for _, v := range triplets {
		added := false
		for _, jV := range uniqueTriplets {
			if v[0] == jV[0] && v[1] == jV[1] && v[2] == jV[2] {
				added = true
				break
			}
		}
		if !added {
			uniqueTriplets = append(uniqueTriplets, v)
		}
	}

	return uniqueTriplets
}
