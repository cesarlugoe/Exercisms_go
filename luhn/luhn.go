package luhn

import (
	"strconv"
	"strings"
)

// Valid comment
func Valid(number string) bool {
	var sum int64
	var alternate bool
	count := 0

	number = strings.Join(strings.Fields(number), "")

	for i := len(number) - 1; i >= 0; i-- {
		n, err := strconv.ParseInt(number[i:i+1], 0, 64)
		if err != nil {
			return false
		}
		if alternate {
			n = n * 2
			if n > 9 {
				n = n - 9
			}
		}
		alternate = !alternate
		sum = sum + n
		count++
	}

	if count < 2 {
		return false
	}
	return sum%10 == 0
}
