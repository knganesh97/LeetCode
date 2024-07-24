package problem1530

import "fmt"

const INPUT_FILE_PATH = "problem1530\\sampleInputs.txt"

type Input struct {
	root     *TreeNode
	distance int
}

type Inputs []Input

type Output struct {
	pairs int
}

type Outputs []Output

func (input *Input) ScanInput() { //to be updated
	input.root = &TreeNode{1, nil, nil}
	input.distance = 1
}

func (output *Output) CalculateOutput(input any) {
	output.pairs = countPairs(input.(Input).root, input.(Input).distance)
}

func (output *Output) PrintOutput() {
	fmt.Println(output.pairs)
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
