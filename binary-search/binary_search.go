package binarysearch

func SearchInts(list []int, key int) int {
	length := len(list)
	if length == 0 {
		return -1
	}
	var middle = length / 2
	if key == list[middle] {
		return middle
	} else if key > list[middle] {
		return toRight(middle, list, key)
	} else {
		return toLeft(middle, list, key)
	}
}

func toRight(middle int, list []int, key int) int {
	pos := middle + 1
	val := SearchInts(list[pos:], key)
	if val == -1 {
		pos = 0
	}
	return pos + val
}

func toLeft(middle int, list []int, key int) int {
	return SearchInts(list[:middle], key)
}
