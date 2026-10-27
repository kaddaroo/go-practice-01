package main

import "fmt"

func main() {
	a := 20
	b := 25
	c := 35
	var amount int = a + b + c
	var average float64 = float64(a+b+c) / 3
	fmt.Printf("Amount: %d\nAverage: %.2f\n", amount, average)
}
