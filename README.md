# sorter

`sorter` is a library designed to provide efficient and flexible sorting algorithms. It supports a variety of sorting techniques, making it a versatile tool for different use cases, whether educational or practical.

## Installation

To use the `sorter` library in your Go project, you need to install it using `go get`:

```bash
go get github.com/ImuS663/sorter/[sort-algorithm-package]
```

## Usage

Here's a basic example of how to use the `sorter` library:

```go
package main

import (
    "fmt"
    "github.com/ImuS663/sorter/merge"
)

func main() {
    data := []int{5, 3, 8, 1, 2}
    sorter := Merge[int]{}
    sortedData := sorter.Sort(data)
    fmt.Println("Sorted Data:", sortedData)
}
```

Replace Merge[int]{} with any other sorting struct provided by the library based on your needs.

## Supported Algorithms

- Bubble Sort [`go get github.com/ImuS663/sorter/bubble`]
- Cocktail Sort [`go get github.com/ImuS663/sorter/cocktail`]
- Gnome Sort [`go get github.com/ImuS663/sorter/gnome`]
- Heap Sort [`go get github.com/ImuS663/sorter/heap`]
- Insertion Sort [`go get github.com/ImuS663/sorter/insertion`]
- Intro Sort [`go get github.com/ImuS663/sorter/intro`]
- Merge Sort [`go get github.com/ImuS663/sorter/merge`]
- Quick Sort [`go get github.com/ImuS663/sorter/quick`]
- Tree Sort [`go get github.com/ImuS663/sorter/tree`]

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/ImuS663/sorter/blob/main/LICENSE) file for details.