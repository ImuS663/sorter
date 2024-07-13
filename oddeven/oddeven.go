package oddeven

import "github.com/ImuS663/sorter"

type OddEven[T sorter.Number] struct{}

func (o OddEven[T]) Sort(arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	for sorted := false; !sorted; {
		sorted = true

		for i := 1; i < len(arr)-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
		for i := 0; i < len(arr)-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
	}

	return arr
}
