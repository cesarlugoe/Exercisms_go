package grains

import "errors"

// Square func comment
func Square(n int) (uint64, error) {
	var r uint64
	var res uint64
	r = 1

	if n < 1 || n > 64 {
		return 0, errors.New("invalid number")
	}
	for i := 1; i <= 64; i++ {

		if n == i {
			res = r
		}
		r = r * 2
	}
	return res, nil
}

// Total comment
func Total() uint64 {
	return 18446744073709551615
}
