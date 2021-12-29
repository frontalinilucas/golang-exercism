package isogram

import (
	"strings"
)

const (
	space  = 32
	hyphen = 45
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	letters := make(map[rune]bool)
	for _, l := range word {
		if isSpaceOrHyphen(l) {
			continue
		}
		if ok := letters[l]; ok {
			return false
		}
		letters[l] = true
	}
	return true
}

func isSpaceOrHyphen(l int32) bool {
	return l == space || l == hyphen
}
