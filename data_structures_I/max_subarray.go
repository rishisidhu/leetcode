package main

import (
	"fmt"
	"math"
)

func sum(s int, e int, nums []int) int {
	sum_arr := 0
	for i := s; i <= e; i++ {
		sum_arr = sum_arr + nums[i]
	}
	return sum_arr
}

func maxSubArray1(nums []int) int {
	var start, end, max int
	//var max_start, max_end int
	start_ind := 0
	max = math.MinInt
	for start_ind < len(nums) {
		start = start_ind
		for end = start_ind; end < len(nums); end++ {
			sumarr := sum(start, end, nums)
			if sumarr > max {
				//max_start = start
				//max_end = end
				max = sumarr
			}
			//fmt.Print(start, end, sumarr, " | ", max_start, max_end, max)
			//fmt.Print("\n")
		}
		start_ind++
	}
	return max
}

func maxSubArray2(nums []int) int {
	sum_sequence := nums[0]

	cnums := make([]int, 0, 1000)
	if len(nums) == 1 {
		cnums = append(cnums, nums[0])
	}
	for i := 1; i < len(nums); i++ {
		if nums[i]*nums[i-1] > 0 {
			sum_sequence = sum_sequence + nums[i]
			if i == len(nums)-1 {
				cnums = append(cnums, sum_sequence)
			}
		} else {
			cnums = append(cnums, sum_sequence)
			sum_sequence = nums[i]
			if i == len(nums)-1 {
				cnums = append(cnums, sum_sequence)
			}
		}
	}
	fmt.Print("CNUMS= ", cnums, "\n")
	sum := 0
	prevsum := 0
	maxsum := math.MinInt
	negcount := 0
	minneg := math.MinInt
	for sind := 0; sind < len(nums); sind++ {
		if nums[sind] < 0 {
			if nums[sind] > minneg {
				minneg = nums[sind]
			}
			negcount++
		} else {
			break
		}
	}
	if negcount == len(nums) {
		return minneg
	}
	for ind := 0; ind < len(cnums); ind++ {
		prevsum = sum
		sum = sum + cnums[ind]
		fmt.Print(ind, ". ", cnums[ind], " : ", prevsum, " | ", sum, " | ", maxsum, "\n")
		if sum < 0 {
			sum = 0
		} else if prevsum > sum {

			if maxsum < prevsum {
				maxsum = prevsum
				sum = 0

			}
		} else {
			if maxsum < sum {
				maxsum = sum
			}
		}
	}

	return maxsum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Print("Method1: ", maxSubArray1(nums), "\n")
	fmt.Print("Method2: ", maxSubArray2(nums))
}
