package jump_game

func CanJump(nums []int) bool {
	lastGoodIdxPos := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		idxPlusNumV := i + nums[i]
		if idxPlusNumV >= lastGoodIdxPos {
			lastGoodIdxPos = i
		}
	}
	return lastGoodIdxPos == 0
}
