package main

import (
	problem "LeetCode/problem1380"
)

const inputFilePath = "problem1380\\sampleInputs.txt"

func main() {
	var inputs problem.Inputs
	var outputs problem.Outputs
	inputs.ScanInputs(inputFilePath)
	outputs.CalculateOutputs(inputs)
	outputs.PrintOutputs()
}
