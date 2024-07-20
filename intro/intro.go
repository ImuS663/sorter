package intro

import (
	"github.com/ImuS663/sorter"
	"github.com/ImuS663/sorter/heap"
	"github.com/ImuS663/sorter/insertion"
	"math"
)

type Intro[T sorter.Number] struct{}

func (i Intro[T]) Sort(arr []T) []T {
	maxDepth := math.Floor(math.Log2(float64(len(arr)))) * 2

	i.sort(arr, maxDepth)

	return arr
}

func (i Intro[T]) sort(arr []T, maxDepth float64) {
	if len(arr) < 16 {
		s := insertion.Insertion[T]{}
		s.Sort(arr)
	} else if maxDepth == 0 {
		s := heap.Heap[T]{}
		s.Sort(arr)
	} else {
		p := partition(arr)
		i.sort(arr[:p-1], maxDepth-1)
		i.sort(arr[p+1:], maxDepth-1)
	}
}

func partition[T sorter.Number](arr []T) int {
	x := arr[len(arr)-1]
	i := 0

	for j := 1; j < len(arr); j++ {
		if arr[j] <= x {
			i++

			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[len(arr)-1] = arr[len(arr)-1], arr[i+1]

	return i + 1
}
