package main

import "fmt"

type observer interface {
	onUpdate(data string)
}

type DataListener struct {
	Name string
}

func (dl *DataListener) onUpdate(data string) {
	fmt.Println("Listener: ", dl.Name, "got data change: ", data)
}
