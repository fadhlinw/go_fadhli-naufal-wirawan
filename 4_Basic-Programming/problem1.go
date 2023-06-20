package main

import "fmt"

func problem1() {
	var T, r float64

	fmt.Println("Masukkan nilai T dan r:")
	fmt.Scanf("%f", &T)
	fmt.Scanf("%f", &r)

	LP := 2*3.14*r*r + 2*3.14*r*T
	fmt.Println("Luas Permukaan Tabung:", LP)
}
