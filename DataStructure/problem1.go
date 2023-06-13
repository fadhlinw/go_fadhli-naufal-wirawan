package main

import "fmt"

func mergeArrays(arrayA, arrayB []string) []string {
	merged := append(arrayA, arrayB...)

	seen := make(map[string]bool)
	result := []string{}

	for _, name := range merged {

		if !seen[name] {
			seen[name] = true
			result = append(result, name)
		}
	}

	return result
}

func main() {
	arrayA := []string{"kazuya", "jin", "lee"}
	arrayB := []string{"kazuya", "feng"}

	merged := mergeArrays(arrayA, arrayB)
	fmt.Println(merged)

	arrayC := []string{"lee", "jin", "panda"}
	arrayD := []string{"kazuya", "panda"}

	merged = mergeArrays(arrayC, arrayD)
	fmt.Println(merged)
}
