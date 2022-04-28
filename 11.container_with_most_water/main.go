package container_with_most_water

func MaxArea(height []int) int {
	maxArea := 0
	for i := 0; i < len(height); i++ {
		iHeight := height[i]

		for j := i + 1; j < len(height); j++ {
			xDistance := j - i
			jHeight := height[j]

			var area int
			if iHeight > jHeight {
				area = (iHeight - (iHeight - jHeight)) * xDistance
			} else if jHeight > iHeight {
				area = (jHeight - (jHeight - iHeight)) * xDistance
			} else {
				area = iHeight * xDistance
			}

			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}
