package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	data [][]int
}

var ErrMalformedMatrix = errors.New("malformed string")

func New(s string) (*Matrix, error) {
	m := &Matrix{}
	lines := strings.Split(s, "\n")
	m.data = make([][]int, len(lines))
	for i := range lines {
		n := strings.Split(strings.TrimSpace(lines[i]), " ")
		m.data[i] = make([]int, len(n))
		for j := range n {
			number, err := strconv.Atoi(n[j])
			if err != nil {
				return nil, ErrMalformedMatrix
			}
			m.data[i][j] = number
		}
	}
	if !m.valid() {
		return nil, ErrMalformedMatrix
	}
	return m, nil
}

func (m *Matrix) valid() bool {
	for i := 1; i < len(m.data); i++ {
		if len(m.data[i]) != len(m.data[i-1]) {
			return false
		}
	}
	return true
}

func (m *Matrix) Cols() [][]int {
	lenRows := len(m.data)
	lenCols := len(m.data[0])
	result := make([][]int, lenCols)
	for i := 0; i < lenCols; i++ {
		result[i] = make([]int, lenRows)
	}
	for i := range m.data {
		for j := range m.data[i] {
			result[j][i] = m.data[i][j]
		}
	}
	return result
}

func (m *Matrix) Rows() [][]int {
	result := make([][]int, len(m.data))
	for i := range m.data {
		result[i] = make([]int, len(m.data[i]))
		for j := range m.data[i] {
			result[i][j] = m.data[i][j]
		}
	}
	return result
}

func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m.data) {
		return false
	}
	if col < 0 || col >= len(m.data[0]) {
		return false
	}

	m.data[row][col] = val
	return true
}
