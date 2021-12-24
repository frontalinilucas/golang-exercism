package palindrome

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type data struct {
	fmin, fmax int
	pmin, pmax Product
	errPrefix  string
}

var testData = []data{
	{1, 9,
		Product{}, // zero value means don't bother to test it
		Product{Value: 9, Factorizations: [][2]int{{1, 9}, {3, 3}}},
		""},
	{10, 99,
		Product{Value: 121, Factorizations: [][2]int{{11, 11}}},
		Product{Value: 9009, Factorizations: [][2]int{{91, 99}}},
		""},
	{100, 999,
		Product{Value: 10201, Factorizations: [][2]int{{101, 101}}},
		Product{Value: 906609, Factorizations: [][2]int{{913, 993}}},
		""},
	{4, 10, Product{}, Product{}, "no palindromes"},
	{10, 4, Product{}, Product{}, "fmin > fmax"},
}

// Bonus curiosities. Can a negative number be a palindrome? Most say no.

var bonusData = []data{
	// The following two test cases have the same input, but different expectations. Uncomment just one or the other.

	// Here you can test that you can reach the limit of the largest palindrome made of two 2-digit numbers.
	// {-99, -10, Product{}, Product{}, "Negative limits"},

	// You can still get non-negative products from negative factors.
	{-99, -10,
		Product{Value: 121, Factorizations: [][2]int{{-11, -11}}},
		Product{Value: 9009, Factorizations: [][2]int{{-99, -91}}},
		""},

	// The following two test cases have the same input, but different expectations. Uncomment just one or the other.

	//In case you reverse the *digits* you could have the following cases:
	//- the zero has to be considered
	//{-2, 2,
	//	Product{0, [][2]int{{-2, 0}, {-1, 0}, {0, 0}, {0, 1}, {0, 2}}},
	//	Product{4, [][2]int{{-2, -2}, {2, 2}}},
	//	""},

	// - you can keep the minus sign in place
	{-2, 2,
		Product{Value: -4, Factorizations: [][2]int{{-2, 2}}},
		Product{Value: 4, Factorizations: [][2]int{{-2, -2}, {2, 2}}},
		""},
}

func TestPalindromeProducts(t *testing.T) {
	t.Run("testData", func(t *testing.T) {
		testPalindromeProducts(t, testData)
	})
	t.Run("bonusData", func(t *testing.T) {
		testPalindromeProducts(t, bonusData)
	})
}

func testPalindromeProducts(t *testing.T, data []data) {
	// Uncomment the following line and the bonusData var above to add the bonus test to the default tests
	// testData = append(testData, bonusData...)
	for _, test := range data {
		// common preamble for test failures
		ret := fmt.Sprintf("Products(%d, %d) returned",
			test.fmin, test.fmax)
		// test
		pmin, pmax, err := Products(test.fmin, test.fmax)
		// we check if err is of error type
		var _ error = err
		switch {
		case err == nil:
			if test.errPrefix > "" {
				t.Fatalf(ret+" err = nil, want %q", test.errPrefix+"...")
			}
		case test.errPrefix == "":
			t.Fatalf(ret+" err = %q, want nil", err)
		case !strings.HasPrefix(err.Error(), test.errPrefix):
			t.Fatalf(ret+" err = %q, want %q", err, test.errPrefix+"...")
		default:
			continue // correct error, no further tests for this test case
		}
		matchProd := func(ww string, rp, wp Product) {
			if len(wp.Factorizations) > 0 && // option to skip test
				!reflect.DeepEqual(rp, wp) {
				t.Fatal(ret, ww, "=", rp, "want", wp)
			}
		}
		matchProd("pmin", pmin, test.pmin)
		matchProd("pmax", pmax, test.pmax)
	}
}

func BenchmarkPalindromeProducts(b *testing.B) {
	b.Run("testData", func(b *testing.B) {
		benchmarkPalindromeProducts(b, testData)
	})
	b.Run("bonusData", func(b *testing.B) {
		benchmarkPalindromeProducts(b, bonusData)
	})
}

func benchmarkPalindromeProducts(b *testing.B, data []data) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range data {
			Products(test.fmin, test.fmax)
		}
	}
}
