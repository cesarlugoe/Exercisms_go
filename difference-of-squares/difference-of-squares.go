package diffsquares

// SquareOfSum comment
func SquareOfSum(n int) int {
	r := n * (n + 1) / 2
	return r * r
}

// SumOfSquares Comment
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference Comment
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
