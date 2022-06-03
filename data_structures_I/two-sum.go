package main

import "fmt"

func twoSum(nums []int, target int) []int {
	elMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		elMap[nums[i]] = i
	}
	//fmt.Print(elMap, "\n")
	var ind1, ind2 int
	for i := 0; i < len(nums); i++ {
		search := target - nums[i]
		ind, ok := elMap[search]
		//fmt.Print("i:", i, " ind2:", ind2, " target:", target, " searchh:", search, ok, ind, "\n")
		if ind == i {
			continue
		} else if ok == true {
			ind1 = i
			ind2 = ind
			//fmt.Print("inelse - ", ind1, ind2, target, search, ok, ind, "\n")
			break
		}
	}
	//fmt.Print(ind1, ind2, target, "\n")
	return []int{ind1, ind2}
}

func twoSum1(nums []int, target int) []int {
	ind1 := 0
	ind2 := 0
	currsum := 0
	for i := 0; i < len(nums); {
		for j := i + 1; j < len(nums); {
			currsum = nums[i] + nums[j]
			//fmt.Print("(", i, ",", j, ") : ", nums[i], ",", nums[j], " | ", currsum, "\n")
			if currsum == target {
				ind1 = i
				ind2 = j
				break
			}
			j++
		}
		i++
	}
	return []int{ind1, ind2}
}

func main() {
	nums := []int{1, 3, 4, 2}
	fmt.Print("input: ", nums, "\n")
	fmt.Print(twoSum(nums, 6))

}
