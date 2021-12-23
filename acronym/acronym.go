// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) (result string) {
	words := strings.Fields(strings.ReplaceAll(s, "-", " "))
	for _, word := range words {
		result += getChar([]rune(word))
	}
	return result
}

func getChar(s []rune) string {
	for _, l := range s {
		if isValid(l) {
			return strings.ToUpper(string(l))
		}
	}
	return ""
}

func isValid(c rune) bool {
	return (c >= 65 && c <= 90) ||
		(c >= 97 && c <= 122)
}
