package main

import (
	"fmt"
)

func stringMask(arg string) string {
	//fmt.Println(arg)

	limit := len(arg)
	buffer := arg[limit-4:]

	template := "**** **** **** "
	results := template + buffer

	return results
}

func main() {
	results := stringMask("123456789")
	fmt.Println(results)
}
