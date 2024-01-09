package main

import (
	"fmt"
)

func main() {
	var s string
	m := map[rune]rune{
		'A':'T',
		'T':'A',
		'C':'G',
		'G':'C',
	}
	fmt.Scanf("%s", &s)
	rs := []rune(s)

	for i := len(rs)-1; i >= 0; i-- {
		fmt.Printf("%c", m[rs[i]])
	}

	fmt.Println()
}