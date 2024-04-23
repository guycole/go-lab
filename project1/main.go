package main

import "fmt"

func main() {
	fmt.Printf("project1\n")

	temp1 := []string{"a", "b", "c"}
	fmt.Printf("temp1: %v\n", temp1)

	var result string
	for _, element := range temp1 {
		if len(result) > 0 {
			result += ", "
		}

		result += element
	}

	fmt.Printf("result: %v\n", result)
}
