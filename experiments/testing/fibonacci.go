package main

import (
	"fmt"
	"math"
)

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		f := x
		x, y = y, f+y

		return f
	}
}

func fibonacciByFormula(n int) int {
	phi := -1 / math.Phi
	return int((math.Pow(math.Phi, float64(n)) - math.Pow(phi, float64(n))) / math.Sqrt(5.0))
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	fmt.Println(fibonacciByFormula(0))
	fmt.Println(fibonacciByFormula(1))
	fmt.Println(fibonacciByFormula(2))
	fmt.Println(fibonacciByFormula(5))
}
