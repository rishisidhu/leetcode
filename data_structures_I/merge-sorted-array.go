package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	enums1 := make([]int, 0, m)
	for i := 0; i < m; i++ {
		enums1 = append(enums1, nums1[i])
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
	}
	if n != 0 {
		//fmt.Print(enums1, "\n")
		ptr1 := 0
		ptr2 := 0
		count := 0
		for {
			if ptr1 == m && ptr2 != n {
				nums1[count] = nums2[ptr2]
				ptr2++
				count++
				continue
			} else if ptr1 != m && ptr2 == n {
				nums1[count] = enums1[ptr1]
				ptr1++
				count++
				continue
			} else if ptr1 == m && ptr2 == n {
				break
			}
			if enums1[ptr1] <= nums2[ptr2] {
				nums1[count] = enums1[ptr1]
				ptr1++
				count++
			} else {
				nums1[count] = nums2[ptr2]
				ptr2++
				count++
			}
			fmt.Print("P1: ", ptr1, " | P2: ", ptr2, " | count: ", count, "\n")

		}
	}

}

func main() {
	nums1 := []int{2, 0}
	nums2 := []int{1}
	m := 1
	n := 1
	merge(nums1, m, nums2, n)
	fmt.Print(nums1)
}
