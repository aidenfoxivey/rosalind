package main

import (
	"fmt"
)

func main() {
	var dnaString string
	m := make(map[rune]int)
	fmt.Scanf("%s", &dnaString)

	for _, r := range []rune(dnaString) {
		m[r]++
	}

	fmt.Printf("%d %d %d %d\n", m['A'], m['C'], m['G'], m['T'])
}
