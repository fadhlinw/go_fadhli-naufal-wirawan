package main

import (
	"fmt"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	total := 0
	for _, score := range s.score {
		total += score
	}
	average := float64(total) / float64(len(s.score))
	return average
}

func (s Student) Min() (min int, name string) {
	min = s.score[0]
	name = s.name[0]

	for i := 1; i < len(s.score); i++ {
		if s.score[i] < min {
			min = s.score[i]
			name = s.name[i]
		}
	}

	return min, name
}

func (s Student) Max() (max int, name string) {
	max = s.score[0]
	name = s.name[0]

	for i := 1; i < len(s.score); i++ {
		if s.score[i] > max {
			max = s.score[i]
			name = s.name[i]
		}
	}

	return max, name
}

func main() {
	var a = Student{}

	for i := 0; i < 6; i++ {
		var name string
		fmt.Print("Input Student ", i+1, " Name: ")
		fmt.Scan(&name)
		a.name = append(a.name, name)

		var score int
		fmt.Print("Input ", name, " Score: ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\nAvarage Score of Students is", a.Average())

	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Student is: ", nameMax, "(", scoreMax, ")")

	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Student is: ", nameMin, "(", scoreMin, ")")
}
