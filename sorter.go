package sorter

type Sorter[T Number] interface {
	Sort(arr []T) []T
}
