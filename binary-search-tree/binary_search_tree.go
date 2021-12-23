package binarysearchtree

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// NewBst creates and returns a new SearchTreeData.
func NewBst(i int) SearchTreeData {
	return SearchTreeData{data: i}
}

// Insert inserts an int into the SearchTreeData.
// Inserts happen based on the rules of a BinarySearchTree
func (std *SearchTreeData) Insert(i int) {
	if i <= std.data {
		if std.left != nil {
			std.left.Insert(i)
		} else {
			t := NewBst(i)
			std.left = &t
		}
	} else {
		if std.right != nil {
			std.right.Insert(i)
		} else {
			t := NewBst(i)
			std.right = &t
		}
	}
}

// MapString returns the ordered contents of SearchTreeData as a []string.
// The values are in increasing order starting with the lowest int value.
// SearchTreeData that has the numbers [1,3,7,5] added will return the
// []string ["1", "3", "5", "7"].
func (std *SearchTreeData) MapString(f func(int) string) (result []string) {
	if std == nil {
		return result
	}
	result = append(result, std.left.MapString(f)...)
	result = append(result, f(std.data))
	result = append(result, std.right.MapString(f)...)
	return result
}

// MapInt returns the ordered contents of SearchTreeData as an []int.
// The values are in increasing order starting with the lowest int value.
// SearchTreeData that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (std *SearchTreeData) MapInt(f func(int) int) (result []int) {
	if std == nil {
		return result
	}
	result = append(result, std.left.MapInt(f)...)
	result = append(result, f(std.data))
	result = append(result, std.right.MapInt(f)...)
	return result
}
