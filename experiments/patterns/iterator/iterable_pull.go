package main

type IterableCollection interface {
	createIterator() iterator
}

type iterator interface {
	hasNext() bool
	next() *Book
}

type BookIterator struct {
	current int
	books   []Book
}

func (b *BookIterator) hasNext() bool {
	if b.current < len(b.books) {
		return true
	}
	return false
}

func (b *BookIterator) next() *Book {
	if b.hasNext() {
		bk := b.books[b.current]
		b.current++
		return &bk
	}
	return nil
}
