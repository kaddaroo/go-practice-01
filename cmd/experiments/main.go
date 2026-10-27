package main

import "fmt"

func main() {
	const language = "Go"
	var name string = "Ethan"
	age := 17
	var interestInCareer bool = true
	percentageOfCourseCompleted := 0.22
	fmt.Printf("Hello! My name is %s, I'm %d years old.\n", name, age)
	fmt.Println("I am really interested in a backend development in language that named:", language)
	fmt.Print("I have completed ", percentageOfCourseCompleted*100, "% of the course.\n")
	fmt.Printf("My interest in a career in backend development is: %t\n", interestInCareer)
}
