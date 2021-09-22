func islandPerimeter(grid [][]int) int {

	iResult := 0
	row := len(grid)
	col := len(grid[0])
	for i := 0; i < row; i++ {
		for k := 0; k < col; k++ {
			if grid[i][k] == 1 {
				if i-1 < 0 || grid[i-1][k] == 0 {
					iResult++
				}

				if k-1 < 0 || grid[i][k-1] == 0 {
					iResult++
				}

				if i+1 >= row || grid[i+1][k] == 0 {
					iResult++
				}

				if k+1 >= col || grid[i][k+1] == 0 {
					iResult++
				}
			}
		}
	}

	return iResult
}
