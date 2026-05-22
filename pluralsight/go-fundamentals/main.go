package main

import (
	"fmt"
	"strings"
)

type Student struct {
	Name  string
	Score int
}

func main() {

	students := []Student{
		{Name: "Dent, Arthur", Score: 87},
		{Name: "MacMillan, Tricia", Score: 96},
		{Name: "Prefect, Ford", Score: 64},
	}

	fmt.Println("Select score to print:")
	var option string
	fmt.Scanln(&option)
	var index int
	switch option {
	case "1":
		index = 0
	case "2":
		index = 1
	case "3":
		index = 2
	default:
		fmt.Println("Unknown option, defaulting to 1")
		index = 0
	}
	fmt.Println()

	fmt.Println("Student scores")
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println(students[index].Name, students[index].Score)
}
