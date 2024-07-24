package problem1605

import (
	"LeetCode/constants"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const INPUT_FILE_PATH = "problem1605\\sampleInputs.txt"

type Input struct {
	rowSum []int
	colSum []int
}

type Inputs []Input

type Output struct {
	matrix [][]int
	m      int
	n      int
}

type Outputs []Output

/**
Input Format :
t (number of test cases)
m n (m = number of rows, n = number of columns) (for each test case)
m space separated numbers each representing the sum of a row of the matrix (for each test case)
n space separated numbers each representing the sum of a column of the matrix (for each test case)
*/

func (inputs *Inputs) ScanInputs() {

	//open file
	file, err := os.Open(INPUT_FILE_PATH)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	//create scanner
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	//scan for testCases
	if !scanner.Scan() {
		log.Println(constants.NO_TOKEN_TO_SCAN_ERROR, "test cases")
	}
	scan := scanner.Text()

	testCases, err := strconv.Atoi(scan)
	if err != nil {
		log.Println(err)
	}

	for range testCases {

		var input Input

		//scan for m and n
		if !scanner.Scan() {
			log.Println(constants.NO_TOKEN_TO_SCAN_ERROR, "dimensions")
		}
		scan = scanner.Text()

		dimensions := strings.Split(scan, " ")

		m, err := strconv.Atoi(dimensions[0])
		if err != nil {
			log.Println(err)
		}

		n, err := strconv.Atoi(dimensions[1])
		if err != nil {
			log.Println(err)
		}

		//scan for rowSum
		if !scanner.Scan() {
			log.Println(constants.NO_TOKEN_TO_SCAN_ERROR, "row sums")
		}
		scan = scanner.Text()

		input.rowSum = make([]int, m)
		rowSum := strings.Split(scan, " ")

		for j := 0; j < m; j++ {
			input.rowSum[j], err = strconv.Atoi(rowSum[j])
			if err != nil {
				log.Println(err)
			}
		}

		//scan for colSum
		if !scanner.Scan() {
			log.Println(constants.NO_TOKEN_TO_SCAN_ERROR, "column sums")
		}
		scan = scanner.Text()

		input.colSum = make([]int, n)
		colSum := strings.Split(scan, " ")

		for j := 0; j < m; j++ {
			input.colSum[j], err = strconv.Atoi(colSum[j])
			if err != nil {
				log.Println(err)
			}
		}

		*inputs = append(*inputs, input)
	}
}

func (outputs *Outputs) CalculateOutputs(inputs any) {

	testCases := len(inputs.(Inputs))

	for i := 0; i < testCases; i++ {

		var output Output

		output.matrix = restoreMatrix(inputs.(Inputs)[i].rowSum, inputs.(Inputs)[i].colSum)
		output.m = len(output.matrix)
		if output.m > 0 {
			output.n = len(output.matrix[0])
		}

		*outputs = append(*outputs, output)
	}
}

func (outputs *Outputs) PrintOutputs() {

	for _, output := range *outputs {

		for i := 0; i < output.m; i++ {

			var line strings.Builder

			for j := 0; j < output.n; j++ {

				line.WriteString(strconv.Itoa(output.matrix[i][j]))

				if j < output.n-1 {
					line.WriteByte(' ')
				}
			}

			fmt.Println(line.String())
		}
	}
}
