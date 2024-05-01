package main

import "fmt"

func main() {
	mag1, _ := newPublication("magazine", "times", 120, "times")
	mag2, _ := newPublication("magazine", "economist", 121, "econ")
	np1, _ := newPublication("newspaper", "okdiario", 120, "ok")
	np2, _ := newPublication("newspaper", "peperiodico", 121, "pepe")

	for _, pub := range []iPublication{mag1, mag2, np1, np2} {
		func(pub iPublication) {
			fmt.Println("----------------------")
			fmt.Println("Name: ", pub.getName())
			fmt.Println("Name: ", pub.getPages())
			fmt.Println("Name: ", pub.getPublisher())
		}(pub)
	}
}
