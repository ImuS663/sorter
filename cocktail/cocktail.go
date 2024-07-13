package cocktail

import (
	"github.com/ImuS663/sorter"
)

type Cocktail[T sorter.Number] struct{}

func (c Cocktail[T]) Sort(arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	start, end := 0, len(arr)-1

	for swapped := true; swapped; {
		swapped = false

		for i := 0; i < end; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}

		start++

		swapped = false

		for i := end - 1; i >= 0; i-- {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}

		end--
	}

	return arr
}
