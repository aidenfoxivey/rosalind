package main

import (
	"fmt"
	"math/big"
)

func main() {
	var AA, Aa, aa int
	fmt.Scanf("%d %d %d", &AA, &Aa, &aa)
	possibleCombinations := choose(AA+Aa+aa, 2)

	validCombinations := choose(
		AA,
		2,
	) + float64(
		AA*Aa,
	) + float64(
		AA*aa,
	) + 0.5*float64(
		Aa*aa,
	) + 0.75*choose(
		Aa,
		2,
	)

	fmt.Printf("%f\n", float64(validCombinations/possibleCombinations))
}

func factorial(n int) *big.Int {
	var f big.Int
	f.MulRange(1, int64(n))

	return f.MulRange(1, int64(n))
}

func choose(n int, k int) float64 {
	var bottom big.Int
	var quotient big.Int

	top := factorial(n)
	quotient.Div(top, bottom.Mul(factorial(k), factorial(n-k)))

	return float64(quotient.Uint64())
}
