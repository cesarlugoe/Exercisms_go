package luhn

import (
	"strconv"
	"strings"
)

// Valid comment
func Valid(number string) bool {
	var sum int64
	var alternate bool

	if len(number) < 2 {
		return false
	}
	number = strings.Join(strings.Fields(number), "")

	for i := range number {
		n, err := strconv.ParseInt(number[i:i+1], 0, 64)
		if err != nil {
			return false
		}
		if alternate {
			n = n * 2
			if n > 9 {
				n = (n % 10) + 1
			}
		}
		alternate = !alternate
		sum = sum + n
	}
	return sum%10 == 0
}
