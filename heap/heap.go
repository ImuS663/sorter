package heap

import "github.com/ImuS663/sorter"

type Heap[T sorter.Number] struct{}

func (h Heap[T]) Sort(arr []T) []T {
	heapIfy(arr)

	end := len(arr) - 1

	for end > 0 {
		arr[end], arr[0] = arr[0], arr[end]
		end -= 1

		siftDown(arr, 0, end)
	}

	return arr
}

func heapIfy[T sorter.Number](arr []T) {
	start := len(arr)/2 - 1

	for start >= 0 {
		siftDown(arr, start, len(arr)-1)

		start--
	}
}

func siftDown[T sorter.Number](arr []T, start, end int) {
	root := start
	for root*2+1 <= end {
		child := root*2 + 1

		swap := root
		if arr[swap] < arr[child] {
			swap = child
		}
		if child+1 <= end && arr[swap] < arr[child+1] {
			swap = child + 1
		}
		if swap != root {
			arr[root], arr[swap] = arr[swap], arr[root]

			root = swap
		} else {
			return
		}
	}
}
