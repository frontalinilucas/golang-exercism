package grains

import (
	"errors"
	"fmt"
)

const (
	min = 1
	max = 64
)

func Square(number int) (uint64, error) {
	if number < min || number > max {
		return 0, errors.New(fmt.Sprintf("a chessboard has between %d and %d squares", min, max))
	}
	square, _ := calculate(number)
	return square, nil
}

func Total() uint64 {
	_, total := calculate(max)
	return total
}

func calculate(n int) (uint64, uint64) {
	var square uint64 = 1
	var total uint64 = 1
	for i := 2; i <= n; i++ {
		square *= 2
		total += square
	}
	return square, total
}
