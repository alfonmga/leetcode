package find_min_counterxample

func FindMinCounterexample(N int) []int {
	ans := make([]int, N)
	ans[0] = N
	nextN := N + 1
	for i := 1; i < len(ans); i++ {
		ans[i] = nextN
		nextN++
	}
	return ans
}
