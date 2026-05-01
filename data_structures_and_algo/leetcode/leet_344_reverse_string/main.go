package main

func main() {
	print(string(reverseString([]rune{'h', 'e', 'l', 'l', 'o'})))
}

func reverseString(s []rune) []rune {
	l := 0
	r := len(s) - 1

	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}

	return s
}
