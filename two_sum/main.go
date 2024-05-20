package main

import "fmt"

func twoSum(nums []int, target int) []int {
	//fmt.Println(len(nums))

	for ndx1 := 0; ndx1 < len(nums); ndx1++ {
		for ndx2 := ndx1 + 1; ndx2 < len(nums); ndx2++ {
			//fmt.Printf("%d %d\n", ndx1, ndx2)
			if nums[ndx1]+nums[ndx2] == target {
				result := []int{ndx1, ndx2}
				return result
			}
		}
	}

	return make([]int, 2)
}

func main() {
	candidates := []int{2, 7, 11, 15}
	results := twoSum(candidates, 9)

	for _, value := range results {
		fmt.Println(value)
	}
}
