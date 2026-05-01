package main

import (
	"regexp"
	"strings"
)

func main() {
	print(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
}

func mostCommonWord(paragraph string, banned []string) string {
	bannedSet := map[string]struct{}{}
	for _, word := range banned {
		bannedSet[word] = struct{}{}
	}

	paragraph = strings.ToLower(paragraph)
	re := regexp.MustCompile(`[^a-z]`)
	paragraph = re.ReplaceAllString(paragraph, " ")

	words := strings.Fields(paragraph)
	var mcw string
	max := -1
	freq := make(map[string]int)

	for _, word := range words {
		if _, exists := bannedSet[word]; exists {
			continue
		}

		freq[word]++

		if freq[word] > max {
			max = freq[word]
			mcw = word
		}

	}

	return mcw
}
