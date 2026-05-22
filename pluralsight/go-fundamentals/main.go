package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Dent, Arthur"
	score := 87

	fmt.Println("Student scores")
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println(name, score)
}
