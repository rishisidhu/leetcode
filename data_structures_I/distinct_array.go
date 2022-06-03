package main

import "fmt"

func containsDuplicate(nums []int) bool {
	truth_map := make(map[int]int)
	for _, v := range nums {
		truth_map[v] = truth_map[v] + 1
		if truth_map[v] > 1 {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Print(containsDuplicate(nums))
}
