package main

import (
	"bytes"
	"fmt"
)

func main() {
	var s string
	var t string
	fmt.Scanf("%s", &s)
	fmt.Scanf("%s", &t)

	sRunes := []byte(s)
	tRunes := []byte(t)

	solutions := make([]int, 0)

	for i := 0; i < len(sRunes); i++ {
		if bytes.Equal(sRunes[i:i+len(tRunes)], tRunes) {
			solutions = append(solutions, i+1)
		}
	}

	fmt.Println(solutions)
}
