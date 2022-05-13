package combination_sum

func CombinationSum(candidates []int, target int) [][]int {
	var combinations [][]int

	for i := 0; i < len(candidates); i++ {
		n := candidates[i]
		if n == target {
			combinations = append(combinations, []int{n})
			continue
		} else if n > target {
			continue
		} else if target%n == 0 {
			c := make([]int, target/n)
			for j := 0; j < target/n; j++ {
				c[j] = n
			}
			combinations = append(combinations, c)
			continue
		} else {
			var c []int
			sum := 0
			for sum+n < target {
				sum += n
				if sum < target {
					c = append(c, n)
				}
			}
			for j := 0; j < len(candidates); j++ {
				if candidates[j] == candidates[i] {
					continue
				}
				if sum+candidates[j] == target {
					c = append(c, candidates[j])
					break
				} else if (target-sum)%candidates[j] == 0 {
					_c := make([]int, target/n)
					for j := 0; j < target/n; j++ {
						_c[j] = n
					}
					c = append(c, _c...)
					break
				} else if len(c) > 1 && sum-n+candidates[j] == target {
					c = c[:len(c)-1]
					c = append(c, candidates[j])
					break
				}
			}
			if SumSliceInts(c) == target {
				combinations = append(combinations, c)
			}
		}
	}

	return combinations
}

func SumSliceInts(s []int) int {
	sum := 0
	for j := 0; j < len(s); j++ {
		sum += s[j]
	}
	return sum
}
