package problem1605

func restoreMatrix(rowSum []int, colSum []int) [][]int {

	m := len(rowSum)
	n := len(colSum)

	var matrix = make([][]int, m)

	for i := range matrix {

		matrix[i] = make([]int, n)

		for j := range matrix[i] {

			matrix[i][j] = min(rowSum[i], colSum[j])

			rowSum[i] -= matrix[i][j]
			colSum[j] -= matrix[i][j]
		}
	}

	return matrix
}
