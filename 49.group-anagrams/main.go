package group_anagrams

func GroupAnagrams(strs []string) [][]string {
	var groupAnagrams [][]string
	charsMap := make(map[int][]string)

	for i := 0; i < len(strs); i++ {
		_word := strs[i]
		targetIdx := 0
		for _, runeV := range _word {
			targetIdx += int(runeV)
		}
		_, existing := charsMap[targetIdx]
		if existing {
			charsMap[targetIdx] = append(charsMap[targetIdx], _word)
			continue
		}
		charsMap[targetIdx] = []string{_word}
	}
	for _, words := range charsMap {
		groupAnagrams = append(groupAnagrams, words)
	}

	return groupAnagrams
}
