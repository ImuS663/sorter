package merge

import "github.com/ImuS663/sorter"

type Merge[T sorter.Number] struct{}

func (m Merge[T]) Sort(arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	leftArr := m.Sort(arr[:mid])
	rightArr := m.Sort(arr[mid:])

	return merge(leftArr, rightArr)
}

func merge[T sorter.Number](leftArr, rightArr []T) []T {
	var result []T

	i, j := 0, 0

	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] < rightArr[j] {
			result = append(result, leftArr[i])
			i++
		} else {
			result = append(result, rightArr[j])
			j++
		}
	}

	if i < len(leftArr) {
		result = append(result, leftArr[i:]...)
	}

	if j < len(rightArr) {
		result = append(result, rightArr[j:]...)
	}

	return result
}
