package main

func main() {
	s := "pwwkew"
	print(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	left := 0
	set := map[rune]struct{}{}
	runes := []rune(s)

	for _, v := range runes {
		for {
			if _, exists := set[v]; !exists {
				break
			}

			delete(set, runes[left])
			left++
		}

		set[v] = struct{}{}

		maxLength = max(maxLength, len(set))
	}

	return maxLength
}
