package longest_substring_without_repeating_characters

import "strings"

func LengthOfLongestSubstring(s string) int {
	sChars := strings.Split(s, "")

	startSCharPosIdx := 0
	longestSubsStrCounter := 0
	encounterChars := make(map[string]bool)
	currentStreakCounter := 0
	for i := 0; i < len(sChars); {
		c := sChars[i]

		_, found := encounterChars[c]
		if found {
			encounterChars = make(map[string]bool)
			currentStreakCounter = 0
			startSCharPosIdx += 1
			i = startSCharPosIdx
			continue
		}

		encounterChars[c] = true
		currentStreakCounter += 1
		if currentStreakCounter > longestSubsStrCounter {
			longestSubsStrCounter += 1
		}
		i++
	}

	return longestSubsStrCounter
}
