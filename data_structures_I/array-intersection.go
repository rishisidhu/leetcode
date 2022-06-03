package main

import (
	"fmt"
	"math"
)

func intersect(nums1 []int, nums2 []int) []int {
	int_hash := make(map[int][]int)
	for i := 0; i < len(nums1); i++ {
		search := nums1[i]
		c_slice, ok := int_hash[search]
		if !ok {
			c_slice = []int{1, 0}
			int_hash[search] = c_slice
		} else {
			c_slice = int_hash[search]
			c_slice[0] = c_slice[0] + 1
		}
	}
	//fmt.Print(int_hash)
	//fmt.Print("\n")
	for i := 0; i < len(nums2); i++ {
		search := nums2[i]
		c_slice, ok := int_hash[search]
		if !ok {
			c_slice = []int{0, 1}
			int_hash[search] = c_slice
		} else {
			c_slice = int_hash[search]
			c_slice[1] = c_slice[1] + 1
		}
	}
	//fmt.Print(int_hash)

	int_slice := make([]int, 0, 1000)
	for key, val := range int_hash {
		if val[0] == 0 || val[1] == 0 {
			continue
		} else {
			times := int(math.Min(float64(val[0]), float64(val[1])))
			for i := 0; i < times; i++ {
				int_slice = append(int_slice, key)
			}
		}
	}
	return int_slice
}

func main() {
	nums1 := []int{1, 2, 2, 3, 4, 3, -1}
	nums2 := []int{2, 2, 6, 8, 9, -1}
	fmt.Print("\n", intersect(nums1, nums2))
}
