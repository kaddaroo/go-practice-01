package main

import "fmt"

func main() {
	var memberName string = "Oleg"
	group := "TEP-03-26"
	const course = "Backend-Development on Golang language"
	doneWorkCount := 1
	var interestInBackend bool = true
	fmt.Printf("Участник: %s\n", memberName)
	fmt.Println("Группа:", group)
	fmt.Printf("Курс: %s\n", course)
	fmt.Println("Выполнено работ:", doneWorkCount)
	fmt.Printf("Интерес к бэкенду: %t\n", interestInBackend)
}
