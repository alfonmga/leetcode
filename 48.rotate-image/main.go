package rotate_image

func Rotate(matrix *[][]int) {
	boardRows := len((*matrix)[0])
	// boardSize := int(math.Pow(float64(boardRows), 2))
	// fmt.Printf("Board size: %d\n", boardSize)

	// allocate temporal empty matrix
	_m := make([][]int, boardRows)
	for i := 0; i < boardRows; i++ {
		_m[i] = make([]int, boardRows)
	}

	for i := 0; i < boardRows; i++ {
		for j := 0; j < boardRows; j++ {
			v := (*matrix)[i][j]
			xPos := (boardRows - 1) - i
			yPos := j
			// fmt.Printf("[%d][%d] to pos: (%v,%v)\n", i, j, xPos, yPos)
			_m[yPos][xPos] = v
		}
	}
	// fmt.Printf("_m: %v", _m)

	for i := 0; i < boardRows; i++ {
		for j := 0; j < boardRows; j++ {
			(*matrix)[i][j] = _m[i][j]
		}
	}
}
