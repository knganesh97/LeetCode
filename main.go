package main

import (
	"LeetCode/constants"
	problem "LeetCode/problem1605"
)

func main() {
	var inputs problem.Inputs
	var outputs problem.Outputs
	inputs.ScanInputs(constants.INPUT_FILE_PATH_1605)
	outputs.CalculateOutputs(inputs)
	outputs.PrintOutputs()
}
