package palindrome

import (
	"errors"
	"math"
	"sync"
)

type Product struct {
	Value          int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return Product{}, Product{}, errors.New("fmin > fmax")
	}
	min, max, err := getMinMax(fmin, fmax)
	if err != nil {
		return Product{}, Product{}, err
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		min.addFactorizations(fmin, fmax)
	}()
	go func() {
		defer wg.Done()
		max.addFactorizations(fmin, fmax)
	}()
	wg.Wait()

	return min, max, nil
}

func getMinMax(fmin, fmax int) (Product, Product, error) {
	min := Product{Value: math.MaxInt}
	max := Product{Value: math.MinInt}
	err := errors.New("no palindromes")
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			p := i * j
			if isPalindrome(p) {
				err = nil
				min.add(p, func(p1, p2 int) bool { return p1 < p2 })
				max.add(p, func(p1, p2 int) bool { return p1 > p2 })
			}
		}
	}

	return min, max, err
}

func (p *Product) add(v int, f func(p1, p2 int) bool) {
	if f(v, p.Value) {
		p.Value = v
	}
}

func (p *Product) addFactorizations(fmin, fmax int) {
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			if i*j == p.Value {
				p.Factorizations = append(p.Factorizations, [2]int{i, j})
			}
		}
	}
}

func isPalindrome(num int) bool {
	if num < 0 {
		num = num * -1
	}
	inputNum := num
	var (
		remainder int
		res       int
	)
	for num > 0 {
		remainder = num % 10
		res = (res * 10) + remainder
		num = num / 10
	}
	return inputNum == res
}
