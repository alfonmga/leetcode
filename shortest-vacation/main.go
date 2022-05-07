package solution

import "math"

func Solution(A []int) int {
	locs := make(map[int]bool, len(A))
	for i := 0; i < len(A); i++ {
		locs[A[i]] = true
	}

	var shortestVacationLength int = math.MaxInt
	visitedLocs := make(map[int]bool, len(A))
	var visits int
	for i := 0; i < len(A); i++ {
		for j := i; j < len(A); j++ {
			visits++
			visitedLocs[A[j]] = true
			if len(visitedLocs) == len(locs) {
				if shortestVacationLength > visits {
					shortestVacationLength = visits
				}
				visitedLocs = make(map[int]bool, len(A))
				visits = 0
				break
			}
		}
	}

	return shortestVacationLength
}
