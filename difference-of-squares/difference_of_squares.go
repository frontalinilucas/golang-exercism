package diffsquares

func square(num int) int {
	return num * num
}

func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return square(sum)
}

func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += square(i)
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
