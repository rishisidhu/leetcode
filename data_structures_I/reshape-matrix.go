package main

import "fmt"

func size(mat [][]int) []int {
	array_size := []int{len(mat), len(mat[0])}
	return array_size
}

func printSerpentine(mat [][]int) {
	output := make([]int, 0, 10000)
	size := size(mat)
	for i := 0; i < size[0]; i++ {
		for j := 0; j < size[1]; j++ {
			output = append(output, j+size[1]*i)
			//output = append(output, mat[i][j])
			//new_y := (j + size[1]*i) % size[1]
			//new_x := (j + size[1]*i) / size[1]
			//output
		}
	}
	fmt.Print("\n", output)
}

func reshape(mat [][]int, r int, c int) [][]int {
	size := size(mat)
	if size[0]*size[1] != r*c {
		return mat
	}
	output := make([][]int, 0, r)
	for i := 0; i < r; i++ {
		row := make([]int, 0, c)
		for j := 0; j < c; j++ {
			old_y := (j + c*i) % size[1]
			old_x := (j + c*i) / size[1]
			row = append(row, mat[old_x][old_y])
		}
		output = append(output, row)
	}
	return output
}

func main() {
	mat := [][]int{{11, 12, 13, 14}, {23, 34, 45, -1}, {-3, -4, -5, -1}, {-13, -24, -15, -1}}
	fmt.Print(size(mat), "\n")
	//printSerpentine(mat)
	fmt.Print(reshape(mat, 16, 1))

}
