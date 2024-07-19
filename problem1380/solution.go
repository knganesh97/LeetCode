package problem1380

func luckyNumbers(matrix [][]int) []int {

	m := len(matrix)
	n := len(matrix[0])

	var rowMin = make([]int, m)
	var colMax = make([]int, n)

	for i := 0; i < m; i++ {

		for j := 0; j < n; j++ {

			if j == 0 {
				rowMin[i] = matrix[i][j]
			} else if matrix[i][j] < rowMin[i] {
				rowMin[i] = matrix[i][j]
			}

			if i == 0 {
				colMax[j] = matrix[i][j]
			} else if matrix[i][j] > colMax[j] {
				colMax[j] = matrix[i][j]
			}
		}
	}

	var numsLucky []int

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == rowMin[i] && matrix[i][j] == colMax[j] {
				numsLucky = append(numsLucky, matrix[i][j])
			}
		}
	}

	return numsLucky
}
