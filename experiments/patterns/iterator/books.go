package main

import "fmt"

type BookType int

const (
	HardCover BookType = iota
	SoftCover
	Paperback
	Ebook
)

type Book struct {
	Name      string
	Author    string
	PageCount int
	Type      BookType
}

type Library struct {
	Collection []Book
}

func (l *Library) IterateBooks(f func(book Book) error) {
	var err error
	for _, b := range l.Collection {
		err = f(b)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func (l *Library) createIterator() iterator {
	return &BookIterator{
		books: l.Collection,
	}
}

var lib *Library = &Library{
	Collection: []Book{
		{
			Name:      "N1",
			Author:    "A1",
			PageCount: 877,
			Type:      HardCover,
		},
		{
			Name:      "N2",
			Author:    "A2",
			PageCount: 878,
			Type:      Ebook,
		},
		{
			Name:      "N3",
			Author:    "A3",
			PageCount: 177,
			Type:      SoftCover,
		},
		{
			Name:      "N4",
			Author:    "A4",
			PageCount: 277,
			Type:      Paperback,
		},
	},
}
