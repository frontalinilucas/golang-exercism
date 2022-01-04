package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

type Robot struct {
	name string
}

var nameUsed = map[string]bool{}

const max = 26 * 26 * 10 * 10 * 10

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(nameUsed) >= max {
		return "", errors.New("no names left")
	}
	for {
		name := getName()
		if !nameUsed[name] {
			nameUsed[name] = true
			r.name = name
			return name, nil
		}
	}
}

func (r *Robot) Reset() {
	r.name = ""
}

func getName() string {
	return fmt.Sprintf("%v%v%v%v%v",
		string(rand.Int31n(26)+66),
		string(rand.Int31n(26)+66),
		rand.Int31n(10),
		rand.Int31n(10),
		rand.Int31n(10))
}
