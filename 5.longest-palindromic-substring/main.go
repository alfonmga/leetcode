package longest_palindromic_substring

func reverseRune(r []rune) []rune {
	rReversed := make([]rune, len(r))
	copy(rReversed, r)
	for i, j := 0, len(rReversed)-1; i < len(rReversed)/2; i, j = i+1, j-1 {
		rReversed[i], rReversed[j] = rReversed[j], rReversed[i]
	}
	return rReversed
}
func runeSlicesEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func isRunePalindrome(s []rune) bool {
	return runeSlicesEqual(s, reverseRune(s))
}

func LongestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	sRuneChars := []rune(s)
	longestPalindrome := []rune{sRuneChars[0]}
	for cIdx, c := range sRuneChars {
		numCharsAfterCurrentChar := len(sRuneChars) - (cIdx + 1)
		longestPalindromeCandidate := []rune{c}
		for i := 1; i <= numCharsAfterCurrentChar; i++ {
			nextCharIdx := cIdx + i
			nextChar := sRuneChars[nextCharIdx]
			longestPalindromeCandidate = append(longestPalindromeCandidate, nextChar)
			if isRunePalindrome(longestPalindromeCandidate) && len(longestPalindromeCandidate) > len(longestPalindrome) {
				longestPalindrome = longestPalindromeCandidate
			}
		}
		if len(longestPalindrome) == len(sRuneChars) {
			break
		}
	}

	return string(longestPalindrome)
}
