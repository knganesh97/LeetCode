package main

import (
	problem "LeetCode/problem912"
)

func main() {
	var inputs problem.Inputs
	var outputs problem.Outputs
	inputs.ScanInputs()
	outputs.CalculateOutputs(inputs)
	outputs.PrintOutputs()
}
