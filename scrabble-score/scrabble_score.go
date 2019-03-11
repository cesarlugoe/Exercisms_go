package scrabble

import (
	"strings"
)

// Score comment
func Score(word string) int {
	s := 0
	upW := strings.ToUpper(word)
	for i := 0; i < len(upW); i++ {
		s += findValue(upW[i])
	}
	return s
}

func findValue(l byte) int {
	switch l {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 0
	}
}
