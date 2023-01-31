package valid_parantheses

const (
	openRound   = '('
	closeRound  = ')'
	openCurly   = '{'
	closeCurly  = '}'
	openSquare  = '['
	closeSquare = ']'
)

func isValid(s string) bool {

	stack := []rune{}

	opposites := map[rune]rune{
		closeRound:  openRound,
		closeCurly:  openCurly,
		closeSquare: openSquare,
	}

	for _, c := range s {
		switch c {
		case openRound, openCurly, openSquare:
			stack = append(stack, c)
		case closeRound, closeCurly, closeSquare:
			if (len(stack) == 0) || stack[len(stack)-1] != opposites[c] {
				return false
			}
			stack = stack[0 : len(stack)-1]
		}
	}

	return len(stack) == 0
}
