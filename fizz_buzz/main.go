package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func fizzBuzz(n int) []string {
	results := make([]string, n)

	for ndx := 1; ndx <= n; ndx++ {
		if ndx%3 == 0 && ndx%5 == 0 {
			results[ndx-1] = "FizzBuzz"
		} else if ndx%5 == 0 {
			results[ndx-1] = "Buzz"
		} else if ndx%3 == 0 {
			results[ndx-1] = "Fizz"
		} else {
			results[ndx-1] = strconv.Itoa(ndx)
		}
	}

	return results
}

func main() {
	results := fizzBuzz(15)
	fmt.Println(reflect.TypeOf(results))

	for _, value := range results {
		fmt.Println(value)
	}
}
