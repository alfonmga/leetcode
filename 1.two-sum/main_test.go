package two_sum

import "testing"

func TestTwoSum(t *testing.T) {
	ans := []int{0, 1}
	actual := TwoSum([]int{2, 7, 11, 15}, 9)
	if !Equal(ans, actual) {
		msg := `
Result: %v
Expected result: %v`
		t.Fatalf(msg, actual, ans)
	}
}
func TestTwoSumNonConsecutive(t *testing.T) {
	ans := []int{0, 2}
	actual := TwoSum([]int{3, 2, 3}, 6)
	if !Equal(ans, actual) {
		msg := `
Result: %v
Expected result: %v`
		t.Fatalf(msg, actual, ans)
	}
}
