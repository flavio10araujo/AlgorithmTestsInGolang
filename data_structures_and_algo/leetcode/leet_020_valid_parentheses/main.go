package main

func main() {
	s := "(([]))"
	print(isValid(s))
}

func isValid(s string) bool {
	var stack []rune

	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if !isPair(top, v) {
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}

func isPair(open, close rune) bool {
	switch open {
	case '(':
		return close == ')'
	case '{':
		return close == '}'
	case '[':
		return close == ']'
	default:
		return false
	}
}
