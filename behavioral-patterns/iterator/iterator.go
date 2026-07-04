package iterator

type Iterator interface {
	HasNext() bool
	Next() MenuItem
}

type Collection interface {
	CreateIterator() Iterator
}
