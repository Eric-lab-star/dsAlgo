package main

import "fmt"

func changeMatrix(matrix [3][3]int) [3][3]int {
	row := [3]int{}
	col := [3]int{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if matrix[i][j] == 1 {
				row[i] = 1
				col[i] = 1
			}
		}
	}

	changed := [3][3]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if row[i] == 1 || col[j] == 1 {
				changed[i][j] = 1
			}
		}
	}
	return changed
}

func printMatrx(matrix [3][3]int) {
	for i := 0; i < 3; i++ {
		fmt.Println(matrix[i])
	}
}

func main() {
	matrix := [3][3]int{
		{1, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	printMatrx(matrix)
	changed := changeMatrix(matrix)
	printMatrx(changed)

}

