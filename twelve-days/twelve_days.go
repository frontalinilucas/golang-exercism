package twelve

import (
	"fmt"
	"strings"
)

const lineSong = "On the %s day of Christmas my true love gave to me: %s."

type lyric struct {
	number, position, letter string
}

var lyrics = []lyric{
	{
		number:   "a",
		position: "first",
		letter:   "Partridge in a Pear Tree",
	},
	{
		number:   "two",
		position: "second",
		letter:   "Turtle Doves",
	},
	{
		number:   "three",
		position: "third",
		letter:   "French Hens",
	},
	{
		number:   "four",
		position: "fourth",
		letter:   "Calling Birds",
	},
	{
		number:   "five",
		position: "fifth",
		letter:   "Gold Rings",
	},
	{
		number:   "six",
		position: "sixth",
		letter:   "Geese-a-Laying",
	},
	{
		number:   "seven",
		position: "seventh",
		letter:   "Swans-a-Swimming",
	},
	{
		number:   "eight",
		position: "eighth",
		letter:   "Maids-a-Milking",
	},
	{
		number:   "nine",
		position: "ninth",
		letter:   "Ladies Dancing",
	},
	{
		number:   "ten",
		position: "tenth",
		letter:   "Lords-a-Leaping",
	},
	{
		number:   "eleven",
		position: "eleventh",
		letter:   "Pipers Piping",
	},
	{
		number:   "twelve",
		position: "twelfth",
		letter:   "Drummers Drumming",
	},
}

func Verse(i int) string {
	l := lyrics[i-1]
	line := make([]string, i)
	for j := i - 1; j >= 0; j-- {
		line[i-j-1] = fmt.Sprintf("%s %s", lyrics[j].number, lyrics[j].letter)
	}
	length := len(line)
	if length > 1 {
		line[length-1] = fmt.Sprintf("and %s", line[length-1])
	}
	return fmt.Sprintf(lineSong, l.position, strings.Join(line, ", "))
}

func Song() string {
	lines := make([]string, 12)
	for i := 1; i <= 12; i++ {
		lines[i-1] = Verse(i)
	}
	return strings.Join(lines, "\n")
}
