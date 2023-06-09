package main

import "fmt"

func main() {

	// input

	var studentScore int = 60

	// Process: Your Solution Code Here
	if studentScore >= 80 && studentScore <= 100 {
		fmt.Println("A", studentScore)
	} else if studentScore >= 65 && studentScore <= 79 {
		fmt.Println("B", studentScore)
	} else if studentScore >= 50 && studentScore <= 64 {
		fmt.Println("C", studentScore)
	} else if studentScore >= 35 && studentScore <= 49 {
		fmt.Println("D", studentScore)
	} else {
		fmt.Println("E", studentScore)
	}
	// Output

	// Nilai A
}
