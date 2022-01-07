package erratum

import (
	"errors"
)

func Use(opener ResourceOpener, input string) error {
	r, err := open(opener)
	if err != nil {
		return err
	}
	defer r.Close()

	return frob(r, input)
}

func open(opener ResourceOpener) (Resource, error) {
	r, err := opener()
	if err != nil {
		if errors.As(err, &TransientError{}) {
			return open(opener)
		}
		return nil, err
	}
	return r, nil
}

func frob(res Resource, input string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch t := r.(type) {
			case FrobError:
				res.Defrob(t.defrobTag)
				err = t.inner
			default:
				err = r.(error)
			}
		}
	}()
	res.Frob(input)
	return
}
