package main

import "fmt"

func main() {
	fmt.Print("Enter a number of minutes: ")
	var minutes int
	const minutesInHour = 60
	_, err := fmt.Scan(&minutes)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	if minutes < 0 {
		fmt.Println("Minutes cannot be negative.")
		return
	}

	fmt.Printf("Hours: %d, Minutes: %d\n", minutes/minutesInHour, minutes%minutesInHour)
}
