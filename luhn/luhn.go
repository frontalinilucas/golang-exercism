package luhn

import "unicode"

func Valid(id string) bool {
	r := []rune(id)
	count := 0
	sum := 0
	for i := len(id) - 1; i >= 0; i-- {
		if unicode.IsSpace(r[i]) {
			continue
		}
		if !unicode.IsDigit(r[i]) {
			return false
		}
		count++
		sum += transform(r[i], isDivisibleBy(count, 2))
	}
	if count < 2 {
		return false
	}
	return isDivisibleBy(sum, 10)
}

func isDivisibleBy(num, divisible int) bool {
	return num%divisible == 0
}

func transform(r rune, transform bool) int {
	num := int(r - '0')
	if !transform {
		return num
	}
	num = num * 2
	if num > 9 {
		num -= 9
	}
	return num
}
