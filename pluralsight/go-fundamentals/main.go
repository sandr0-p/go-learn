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

	fmt.Println("Student scores")
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println(students[0].Name, students[0].Score)
	fmt.Println(students[1].Name, students[1].Score)
	fmt.Println(students[2].Name, students[2].Score)
}
