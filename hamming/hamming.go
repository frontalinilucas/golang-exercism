package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	rA := []rune(a)
	rB := []rune(b)
	if len(rA) != len(rB) {
		return 0, errors.New("the strings have different length")
	}
	diff := 0
	for i := range rA {
		if rA[i] != rB[i] {
			diff++
		}
	}
	return diff, nil
}
