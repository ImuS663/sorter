package quick

import "github.com/ImuS663/sorter"

type Quick[T sorter.Number] struct{}

func (q Quick[T]) Sort(arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	sort(arr, 0, len(arr)-1)

	return arr
}

func sort[T sorter.Number](arr []T, p, q int) {
	if p >= q {
		return
	}

	i := partition(arr, p, q)

	sort(arr, p, i-1)
	sort(arr, i+1, q)
}

func partition[T sorter.Number](arr []T, p, q int) int {
	x := arr[q]
	i := p - 1

	for j := p; j < q; j++ {
		if arr[j] <= x {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[q] = arr[q], arr[i+1]
	return i + 1
}
