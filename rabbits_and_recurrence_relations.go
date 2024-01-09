package main

import (
	"fmt"
)
// for n=5, k=3
// f2 = f2 + f1 * 3
// 0 0
// 1 1
// 2 1
// 3 4
// 4 7
// 5 19

func main() {
	var offspring int
	var months int
	fmt.Scanf("%d %d", &months, &offspring)

	f1 := 0
	f2 := 1

	for i := 0; i < months; i++ {
		temp := f2
		f2 = f2 + offspring*f1
		f1 = temp
	}

	fmt.Println(f1)
}