package brackets

import "errors"

var match = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

type pile struct {
	pile []rune
}

func Bracket(input string) bool {
	r := []rune(input)
	p := pile{}
	for _, l := range r {
		if p.process(l) != nil {
			return false
		}
	}
	return p.isOK()
}

func (p *pile) isOK() bool {
	return len(p.pile) == 0
}

func (p *pile) process(s rune) error {
	switch s {
	case '{', '[', '(':
		p.pile = append(p.pile, s)
	case '}', ']', ')':
		if len(p.pile) == 0 {
			return errors.New("there are not any char")
		}
		pos := len(p.pile)-1
		last := p.pile[pos]
		if match[s] != last {
			return errors.New("there is not a polygon")
		}
		p.pile = p.pile[:pos]
	}
	return nil
}