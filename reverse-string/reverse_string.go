package reverse

// String comment
func String(s string) string {

	r := []rune(s)
	var rs []rune

	for i := len(r) - 1; i >= 0; i-- {
		rs = append(rs, r[i])
	}

	return string(rs)
}
