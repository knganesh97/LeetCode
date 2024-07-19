package problem1380

import (
	"LeetCode/util"
	"fmt"
)

type Input struct {
	matrix [][]int
}

type Inputs []Input

type Output struct {
	luckyNumbers []int
}

type Outputs []Output

/**
Input Format :
t (number of test cases)
m n (m = number of rows, n = number of columns) (for each test case)
n space separated numbers representing each row of the matrix (m lines) (for each test case)
*/

func (inputs *Inputs) ScanInputs(filePath string) {

	var r = util.GetFileAsString(filePath)

	var testCases int
	fmt.Sscanf(r, "%d\n", &testCases)
	r = util.RemoveFirstLine(r)

	// *input = make(Inputs, testCases)

	for range testCases {

		var m, n int
		fmt.Sscanf(r, "%d %d\n", &m, &n)
		r = util.RemoveFirstLine(r)

		var input Input

		input.matrix = make([][]int, m)
		for i := range input.matrix {
			input.matrix[i] = make([]int, n)
			for j := range input.matrix[i] {
				var element int
				if j == n-1 {
					fmt.Sscanf(r, "%d\n", &element)
					r = util.RemoveFirstLine(r)
				} else {
					fmt.Sscanf(r, "%d ", &element)
					r = util.RemoveFirstWord(r)
				}
				input.matrix[i][j] = element
			}
		}

		*inputs = append(*inputs, input)
	}

}

func (outputs *Outputs) CalculateOutputs(inputs any) {
	testCases := len(inputs.(Inputs))
	for i := 0; i < testCases; i++ {
		var output Output
		output.luckyNumbers = luckyNumbers(inputs.(Inputs)[i].matrix)
		*outputs = append(*outputs, output)
	}

}

func (outputs *Outputs) PrintOutputs() {
	for _, output := range *outputs {
		n := len(output.luckyNumbers)
		for i := 0; i < n; i++ {
			if i < n-1 {
				fmt.Print(output.luckyNumbers[i], " ")
			} else {
				fmt.Print(output.luckyNumbers[i], "\n")
			}
		}
	}

}
