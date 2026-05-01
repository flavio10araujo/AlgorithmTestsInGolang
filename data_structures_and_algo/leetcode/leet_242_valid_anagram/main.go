package main

func main() {
	println(isAnagram("anagram", "nagaram"))
	print(isAnagram("rat", "car"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	rs := []rune(s)
	rt := []rune(t)
	charFreq := make(map[rune]int)

	for i := range rs {
		charFreq[rs[i]]++
	}

	for i := range rt {
		if charFreq[rt[i]] == 0 {
			return false
		}

		charFreq[rt[i]]--
	}

	return true
}
