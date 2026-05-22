package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	Name  string
	Score int
}

func main() {

	students := []Student{}

	shouldContinue := true

	for shouldContinue {

		fmt.Println("Select an option:")
		fmt.Println("1) Add student")
		fmt.Println("2) Print students")
		fmt.Println("q) Quit")

		var option string
		fmt.Scanln(&option)

		switch option {
		case "1":
			fmt.Println("Enter student name and score:")
			var name, rawScore string
			fmt.Scanln(&name, &rawScore)
			s, _ := strconv.Atoi(rawScore)
			students = append(students, Student{Name: name, Score: s})
		case "2":
			fmt.Println("Student scores")
			fmt.Println("----------------")
			for _, s := range students {
				fmt.Println(s.Name, s.Score)
			}
		case "q":
			shouldContinue = false
		}
	}
}
