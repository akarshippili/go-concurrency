package heap

type Comparable[T any] interface {
	// 	if == returns 0
	// 	else if > returns 1
	// 	else return -1
	Compare(val T) int
}
