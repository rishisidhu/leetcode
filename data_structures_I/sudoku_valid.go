package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func findChar(str string) bool {
	//char := str[i]
	mila := strings.ContainsAny(str, "2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 ")
	//fmt.Printf("%c", char)
	//fmt.Print(" - ", mila, "\n")
	if mila {
		return mila
	}
	return false
}

func isEmpty(board [][]byte) bool {
	valid_array := make([]bool, 18, 18)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if !valid_array[i] {
					valid_array[i] = true
				}
				if !valid_array[j+9] {
					valid_array[j+9] = true
				}
			}
		}
	}
	for i := 0; i < 18; i++ {
		if !valid_array[i] {
			return true
		}
	}
	return false
}

func isValidSudoku(board [][]byte) bool {
	num_map := make(map[int][]int)
	for i := 1; i < 10; i++ {
		num_map[i] = []int{0, 0, 0}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			curr_num, _ := strconv.Atoi(string(board[i][j]))
			data := num_map[curr_num]
			//fmt.Print(i, j, curr_num, data, board[i][j], "\n")
			data[0] = data[0] + int(math.Pow(10, float64(i)))
			data[1] = data[1] + int(math.Pow(10, float64(j)))
			sq_num_i := i / 3
			sq_num_j := j / 3
			sq_num := 3*sq_num_i + sq_num_j
			data[2] = data[2] + int(math.Pow(10, float64(sq_num)))
			//fmt.Print(i, j, sq_num_i, sq_num_j, sq_num, "\n")
			num_map[curr_num] = data
			//fmt.Print("post: ", num_map, data, "\n")
		}
	}
	//fmt.Print(num_map)
	found := false
	for key, value := range num_map {
		fmt.Print(key, value, "\n")
		row := strconv.Itoa(value[0])
		col := strconv.Itoa(value[1])
		square := strconv.Itoa(value[2])
		found = findChar(row)
		if found {
			break
		}
		found = findChar(col)
		if found {
			break
		}
		found = findChar(square)
		if found {
			break
		}
	}
	return !found
}

func main() {
	/*board := [][]byte{{'.', '3', '.', '.', '.', '.', '7', '.', '.'},
	{'6', '.', '.', '1', '9', '5', '.', '8', '.'},
	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	{'8', '.', '.', '.', '6', '.', '.', '9', '3'},
	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	{'.', '6', '.', '.', '.', '.', '2', '.', '.'},
	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}*/
	/*board := [][]byte{
	{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
	{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '3', '.', '.', '.'},
	{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
	{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
	{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
	{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
	{'.', '.', '4', '.', '.', '.', '.', '.', '.'}}*/
	board := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'}}
	fmt.Print("Final Result: ", isValidSudoku(board))
}
