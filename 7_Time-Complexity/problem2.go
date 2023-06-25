package main

import "fmt"

func pow(x, n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	result := pow(x*x, n/2)
	if n%2 == 1 {
		result *= x
	}

	return result

}

func main() {

	fmt.Println(pow(2, 3)) // 8

	fmt.Println(pow(5, 3)) // 125

	fmt.Println(pow(10, 2)) // 100

	fmt.Println(pow(2, 5)) // 32

	fmt.Println(pow(7, 3)) // 343

}
