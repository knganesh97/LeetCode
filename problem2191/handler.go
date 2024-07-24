package problem2191

import (
	"LeetCode/constants"
	"LeetCode/util"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const INPUT_FILE_PATH = "problem2191\\sampleInputs.txt"

type Input struct {
	mapping []int
	nums    []int
}

type Inputs []Input

type Output struct {
	nums []int
}

type Outputs []Output

/**
Input Format :
t (number of test cases)
n (length of nums array) (for each test case)
n space separated numbers representing the elements of nums array (for each test case)
10 space separated digits representing the mapping (for each test case)
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

		//scan for length of nums array
		if !scanner.Scan() {
			log.Println(constants.NO_TOKEN_TO_SCAN_ERROR, "length of nums array")
		}
		scan = scanner.Text()

		n, err := strconv.Atoi(scan)
		if err != nil {
			log.Println(err)
		}

		input.nums = make([]int, n)

		//scan for nums array
		if !scanner.Scan() {
			log.Println(constants.NO_TOKEN_TO_SCAN_ERROR, "nums array")
		}
		scan = scanner.Text()
		nums := strings.Split(scan, " ")
		for j := range input.nums {
			input.nums[j], err = strconv.Atoi(nums[j])
			if err != nil {
				log.Println(err)
			}
		}

		//scan for mapping
		if !scanner.Scan() {
			log.Println(constants.NO_TOKEN_TO_SCAN_ERROR, "mapping")
		}
		scan = scanner.Text()
		mapping := strings.Split(scan, " ")
		for j := range mapping {
			digit, err := strconv.Atoi(mapping[j])
			if err != nil {
				log.Println(err)
			}
			input.mapping = append(input.mapping, digit)
		}

		*inputs = append(*inputs, input)
	}
}

func (outputs *Outputs) CalculateOutputs(inputs any) {
	for _, input := range inputs.(Inputs) {
		var output Output
		output.nums = sortJumbled(input.mapping, input.nums)
		*outputs = append(*outputs, output)
	}
}

func (outputs *Outputs) PrintOutputs() {
	for _, output := range *outputs {
		util.PrintArrayAsText(output.nums)
	}
}