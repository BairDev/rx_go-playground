package iterable

// IntIterator could be an iterator object.
type IntIterable interface {
    Next() (value int, ok bool)
}

// this should be in line with IntIterator
type IterableInts struct {
    position int
    intCollection []int
}

// implementation of Next()
func (ints *IterableInts) Next() (value int, ok bool) {
	ints.position++ // increase current int
	if ints.position > len(ints.intCollection) {
        return 0, false
    }
    return ints.intCollection[ints.position], true
}

// constructor for iterableInts
func newIntInterator(ints []int) *IterableInts {
    return &iterableInts{-1, ints}
}