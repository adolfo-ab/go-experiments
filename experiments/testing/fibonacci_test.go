package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFibonacci(t *testing.T) {
	expectedValues := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	f := fibonacci()

	for i := 0; i < 10; i++ {
		val := f()
		if val != expectedValues[i] {
			t.Fatalf("Wrong value. Expected %d, got %d.", expectedValues[i], val)
		}

	}

}

type testCase struct {
	n        int
	expected int
}

func TestFibonacciByFormula(t *testing.T) {
	testCases := []testCase{
		{0, 0},
		{15, 610},
		{25, 75025},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d", tc.n), func(t *testing.T) {
			val := fibonacciByFormula(tc.n)
			if val != tc.expected {
				t.Fatalf("Wrong value. Expected %d, got %d.", tc.expected, val)
			}
		})
	}
}

func TestFibonaccyByFormulaTestify(t *testing.T) {
	val := fibonacciByFormula(25)
	require.EqualValues(t, val, 75025, "Wrong value.")
}
