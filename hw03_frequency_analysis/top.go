package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	wordsMap := make(map[string]int)
	for _, value := range strings.Fields(text) {
		value = strings.TrimSpace(value)
		wordsMap[value]++
	}

	var keys []string
	for key := range wordsMap {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		if wordsMap[keys[i]] == wordsMap[keys[j]] {
			return strings.Compare(keys[i], keys[j]) < 0
		}
		return wordsMap[keys[i]] > wordsMap[keys[j]]
	})

	if len(keys) > 10 {
		result := keys[:10]
		return result
	}
	return keys
}
