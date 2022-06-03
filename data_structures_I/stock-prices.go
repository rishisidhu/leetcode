package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	diff := make([]int, 0, len(prices))
	diff = append(diff, 0)
	for i := 1; i < len(prices); i++ {
		price_delta := prices[i] - prices[i-1]
		diff = append(diff, price_delta)
	}
	return maxSubArray(diff)
}

func maxSubArray(nums []int) int {
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
	prices := []int{7, 6, 4, 3, 1}
	fmt.Print(maxProfit(prices))

}
