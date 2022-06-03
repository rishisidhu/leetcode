package main

import "fmt"

func expandSlice(row []int) []int {
	exp_row := make([]int, 0, len(row)+1)
	exp_row = append(exp_row, row[0])
	for i := 1; i < len(row); i++ {
		exp_row = append(exp_row, row[i]+row[i-1])
	}
	exp_row = append(exp_row, row[len(row)-1])
	return exp_row
}

func generate(numRows int) [][]int {
	pascalt := make([][]int, 0, 30)
	if numRows == 1 {
		pascalt = append(pascalt, []int{1})
	} else if numRows == 2 {
		pascalt = append(pascalt, []int{1})
		pascalt = append(pascalt, []int{1, 1})
	} else {
		pascalt = append(pascalt, []int{1})
		pascalt = append(pascalt, []int{1, 1})
		for i := 2; i < numRows; i++ {
			row := expandSlice(pascalt[i-1])
			pascalt = append(pascalt, row)
		}
	}
	return pascalt
}

func main() {
	fmt.Print(generate(30))
}
