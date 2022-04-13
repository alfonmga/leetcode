package longest_palindromic_substring

import (
	"strings"
)

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func LongestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	sChars := strings.Split(s, "")

	longestPalindrome := string(s[0])
	for cIdx, c := range sChars {
		remainingCharsAfterCurrentChar := len(sChars) - (cIdx + 1)
		candidateLongestPalindrome := c
		for i := 1; i <= remainingCharsAfterCurrentChar; i++ {
			nextCharIdx := cIdx + i
			nextChar := sChars[nextCharIdx]
			candidateLongestPalindrome += nextChar
			isPalindrome := candidateLongestPalindrome == reverse(candidateLongestPalindrome)
			if isPalindrome && len(candidateLongestPalindrome) > len(longestPalindrome) {
				longestPalindrome = candidateLongestPalindrome
			}
		}
	}

	return longestPalindrome
}
