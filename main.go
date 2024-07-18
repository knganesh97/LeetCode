package main

import (
	"LeetCode/problem1530"
	"bufio"
	"fmt"
	"os"
)

func main() {
	startScan()
	var testCases int
	fmt.Scanln(&testCases)
	var inputs = make(problem1530.Inputs, testCases)
	var outputs = make(problem1530.Outputs, testCases)
	
	for i := range inputs {
		inputs[i].ScanInput()
		outputs[i].CalculateOutput(inputs[i])
	}

	for i := range outputs {
		outputs[i].PrintOutput()
	}
}

type GetInput interface {
	ScanInput()
}

type GetOutput interface {
	CalculateOutput(any)
	PrintOutput()
}

func startScan() {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
}