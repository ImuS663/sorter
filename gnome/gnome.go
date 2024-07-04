package gnome

import "github.com/ImuS663/sorter"

type Gnome[T sorter.Number] struct{}

func (g Gnome[T]) Sort(arr []T) []T {
	i := 0

	for i < len(arr) {
		if i == 0 || arr[i] >= arr[i-1] {
			i++
		} else {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			i--
		}
	}

	return arr
}
