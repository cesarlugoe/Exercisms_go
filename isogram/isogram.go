package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram comment
func IsIsogram(s string) bool {
	if len(s) < 1 {
		return true
	}

	s = strings.ToLower(s)
	w := ""
	for _, r := range s {

		if strings.ContainsRune(w, r) && unicode.IsLetter(r) {
			return false
		}
		w += string(r)
	}
	return true
}

//Better answer found

// func IsIsogram(s string) bool {
// 	s = strings.ToLower(s)

// 	for i, c := range s {
// 		if unicode.IsLetter(c) && strings.ContainsRune(s[i+1:], c) {
// 			return false
// 		}
// 	}

// 	return true
// }
