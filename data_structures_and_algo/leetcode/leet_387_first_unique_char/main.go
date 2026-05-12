package main

import "fmt"

func main() {
	//s := "leetcode" // 0
	s := "leetcodel" // 3
	fmt.Println(firstUniqChar(s))
}

func firstUniqChar(s string) int {
	freq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if freq[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}
