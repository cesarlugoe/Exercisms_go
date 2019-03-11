package hamming

import (
	"errors"
)

// Distance comment
func Distance(a, b string) (int, error) {
	var d int
	var err error
	if len(a) != len(b) {
		return -1, errors.New("Strings have to have same length")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			d++
		}
	}
	return d, err
}
