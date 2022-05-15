package spiral_matrix

import "fmt"

func SpiralOrder(matrix [][]int) []int {
	boardRows, boardCols := len(matrix), len(matrix[0])
	var m []int
	// first row
	m = append(m, matrix[0][0:]...)
	// last column
	for i := 1; i < boardRows; i++ {
		m = append(m, matrix[i][boardCols-1])
	}
	// last row
	for i := (boardCols - 1) - 1; 0 <= i; i-- {
		m = append(m, matrix[boardRows-1][i])
	}
	// first column
	for i := boardRows - 2; 1 <= i; i-- {
		m = append(m, matrix[i][0])
	}
	// inwards
	for i := 1; i <= boardRows-2; i++ {
		fmt.Println(i)
		for j := 1; j <= len(matrix[i])-2; j++ {
			fmt.Println("inwards", matrix[i][j])
			m = append(m, matrix[i][j])
		}
	}
	return m
}
