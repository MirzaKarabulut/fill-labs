package main

import "fmt"

func Question2FirstComeToMind(n int) {
	if n < 2 {
		return
	}
	Question2FirstComeToMind(n / 2)
	fmt.Println(n)
}

func main() {
	Question2FirstComeToMind(9)
}
