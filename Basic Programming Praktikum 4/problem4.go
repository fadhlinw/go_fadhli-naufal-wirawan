package main

import "fmt"

func primeNumber(number int) bool {
	if number < 2 {
		return false
	}

	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true

}

func main() {
	fmt.Println(primeNumber(3))
	fmt.Println(primeNumber(6))
}
