package main

func main() {
    pattern := "abba"
    s := "dog cat cat dog"
    fmt.Println(wordPattern(pattern, s))
}

func wordPattern(pattern string, s string) bool {
    words := strings.Split(s, " ")
    if len(pattern) != len(words) {
        return false
    }

    charToWord := make(map[byte]string)
    wordToChar := make(map[string]byte)

    for i := 0; i < len(pattern); i++ {
        char := pattern[i]
        word := words[i]

        if mappedWord, exists := charToWord[char]; exists {
            if mappedWord != word {
                return false
            }
        } else {
            charToWord[char] = word
        }

        if mappedChar, exists := wordToChar[word]; exists {
            if mappedChar != char {
                return false
            }
        } else {
            wordToChar[word] = char
        }
    }

    return true
}