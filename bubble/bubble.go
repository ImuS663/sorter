package bubble

import "github.com/ImuS663/sorter"

type Bubble[T sorter.Number] struct{}

func (b Bubble[T]) Sort(arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			arr[i], arr[i-1] = arr[i-1], arr[i]
		}
	}

	b.Sort(arr[:len(arr)-1])

	return arr
}
