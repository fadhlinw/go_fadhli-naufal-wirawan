package main

import "fmt"

func akarbilangan() {

	var bilangan int

	fmt.Println("Masukan bilangan:")
	fmt.Scanf("%d", &bilangan)

	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Println(i)
		}
	}

}
