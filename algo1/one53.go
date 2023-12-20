package main

import "fmt"

func main() {
	for hundred := 1; hundred <= 9; hundred++ {
		for ten := 0; ten <= 9; ten++ {
			for one := 0; one <= 9; one++ {
				candidate1 := hundred * 100 + ten * 10 + one
				candidate2 := hundred * hundred * hundred + ten * ten * ten + one * one * one

				if candidate1 == candidate2 {
					fmt.Printf("%d\n", candidate1)
				}
			}
		}
	}
}