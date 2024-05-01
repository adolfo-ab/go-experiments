package main

import "fmt"

func newPublication(pubType string, name string, pages int, publisher string) (iPublication, error) {
	switch pubType {
	case "newspaper":
		return createNewspaper(name, pages, publisher), nil
	case "magazine":
		return createMagazine(name, pages, publisher), nil
	default:
		return nil, fmt.Errorf("no such publication type: %s", pubType)
	}

}
