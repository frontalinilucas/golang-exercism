package raindrops

import (
	"strconv"
)

type rain struct {
	number int
	msg    string
}

var (
	pling = rain{number: 3, msg: "Pling"}
	plang = rain{number: 5, msg: "Plang"}
	plong = rain{number: 7, msg: "Plong"}
)

func Convert(number int) string {
	result := concat(number, pling)
	result += concat(number, plang)
	result += concat(number, plong)
	if result == "" {
		return strconv.Itoa(number)
	}
	return result
}

func concat(n int, r rain) string {
	if isFactor(n, r.number) {
		return r.msg
	}
	return ""
}

func isFactor(number, module int) bool {
	return number%module == 0
}
