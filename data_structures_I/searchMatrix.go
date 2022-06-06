package main

import (
	"fmt"
)

func searchMatrix1(matrix [][]int, target int) bool {
	var i, j int = 0, 0
	m := len(matrix)
	n := len(matrix[0])
	var row = m - 1
	for i < m {
		fmt.Print(i, j, matrix[i][j], target, "\n")
		if matrix[i][j] < target {
			i++
		} else if matrix[i][j] >= target {
			if matrix[i][j] == target {
				return true
			} else {
				if i == 0 {
					return false
				}
				row = i - 1
				break
			}
		}
	}
	j++
	for j < n {
		fmt.Print(row, j, matrix[row][j], target, "\n")
		if matrix[row][j] < target {
			j++
		} else if matrix[row][j] >= target {
			if matrix[row][j] == target {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {

	m := len(matrix)
	n := len(matrix[0])
	var low_i, high_i int = 0, m - 1
	var row int
	/*if matrix[0][0] > target || matrix[m-1][n-1] < target {
		return false
	}*/
	if m == 1 {
		row = 0
	} else {
		for {
			mid_i := (low_i + high_i) / 2
			fmt.Print(low_i, high_i, mid_i, matrix[mid_i][0], target, "\n")
			if matrix[mid_i][0] < target {
				fmt.Print("mid is less than target", matrix[mid_i][0], target, "\n")
				low_i = mid_i
			} else if matrix[mid_i][0] > target {
				fmt.Print("i mid is more than target", matrix[mid_i][0], target, "\n")
				high_i = mid_i
			} else if matrix[mid_i][0] == target {
				return true
			}
			if high_i == low_i {
				return false
			}
			if high_i-low_i == 1 {
				fmt.Print("high and low aree close", high_i, low_i, "\n")
				if matrix[high_i][0] > target {
					row = low_i
				} else if matrix[high_i][0] < target {
					row = high_i
				} else {
					return true
				}
				fmt.Print("entry row:", row, "\n")
				break
			}
		}

	}
	low_j := 0
	high_j := n - 1
	for {
		mid_j := (low_j + high_j) / 2
		fmt.Print(low_j, high_j, mid_j, matrix[row][mid_j], target, "\n")
		if matrix[row][mid_j] < target {
			fmt.Print("mid is less than target", matrix[row][mid_j], target, "\n")
			low_j = mid_j + 1
		} else if matrix[row][mid_j] > target {
			fmt.Print("mid is more than target", matrix[row][mid_j], target, "\n")
			high_j = mid_j - 1
		} else if matrix[row][mid_j] == target {
			return true
		}
		if high_j < low_j {
			return false
		}
		if high_j == low_j {
			fmt.Print("entry j", "\n")
			if matrix[row][high_j] == target {
				return true
			} else {
				return false
			}
		}
	}
}

func main() {
	/*matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}*/
	matrix := [][]int{
		{1}, {3},
	}
	target := 0
	fmt.Print(searchMatrix1(matrix, target), "\n********************\n\n")
	fmt.Print(searchMatrix(matrix, target))
}
