package insertion

import "github.com/ImuS663/sorter"

type Insertion[T sorter.Number] struct{}

func (i Insertion[T]) Sort(arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

	return arr
}
