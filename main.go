package main

import (
	problem "LeetCode/problem2191"
)

func main() {
	var inputs problem.Inputs
	var outputs problem.Outputs
	inputs.ScanInputs()
	outputs.CalculateOutputs(inputs)
	outputs.PrintOutputs()
}
