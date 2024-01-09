package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scanf("%s", &str)
	fmt.Println(strings.Replace(str, "T", "U", -1))
}