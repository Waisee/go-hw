package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Word struct {
	word  string
	count int
}

func (w *Word) Count(value string) {
	value = strings.TrimSpace(value)
	found := false
	for i, word := range wordsSlice {
		if word.word == value {
			wordsSlice[i].count++
			found = true
			break
		}
	}
	if !found {
		wordsSlice = append(wordsSlice, Word{
			value,
			1,
		})
	}
}

var wordsSlice []Word

func Top10(text string) []string {
	var w Word
	for _, value := range strings.Fields(text) {
		w.Count(value)
	}
	sort.Slice(wordsSlice, func(i, j int) bool {
		if wordsSlice[i].count != wordsSlice[j].count {
			return wordsSlice[i].count > wordsSlice[j].count
		}
		return wordsSlice[i].word < wordsSlice[j].word
	})

	result := make([]string, 0, 10)
	for j := range wordsSlice {
		result = append(result, wordsSlice[j].word)
		if len(result) == 10 {
			break
		}
	}

	return result
}
