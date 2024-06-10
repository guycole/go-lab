package main

import (
	"fmt"
	"strconv"
)

func countVowelSubstrings(word string) int {
	for ndx1 := 0; ndx1 < len(word); ndx1++ {
		for ndx2 := ndx1; ndx2 < len(word); ndx2++ {
			temp := word[ndx1 : ndx2+1]
			if len(temp) < 5 {
				continue
			}

			fmt.Println(temp)
			fmt.Println(word[ndx1 : ndx2+1])
		}
	}

	return 0
}

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
	results := countVowelSubstrings("aeiouu")
	fmt.Println(results)

	//	results = countVowelSubstrings("unicornarihan")
	//	fmt.Println(results)

	// results = countVowelSubstrings("cuaieuouac")
	// fmt.Println(results)
}
