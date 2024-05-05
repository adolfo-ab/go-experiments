package main

import "fmt"

func main() {
	lib.IterateBooks(myBookCallback)

	lib.IterateBooks(func(b Book) error {
		fmt.Println("Book author: ", b.Author)
		return nil
	})

	iter := lib.createIterator()
	for iter.hasNext() {
		book := iter.next()
		fmt.Printf("Book %+v\n", book)
	}
}

func myBookCallback(b Book) error {
	fmt.Println("Book title: ", b.Name)
	return nil
}
