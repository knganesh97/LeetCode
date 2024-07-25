package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
* Interfaces
 */
type InputScanner interface {
	ScanInputs()
}

type OutputPrinter interface {
	CalculateOutputs(any)
	PrintOutputs()
}

/**
* Helper Functions
 */
func GetFileAsString(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
	}
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var fileAsString strings.Builder
	for scanner.Scan() {
		fileAsString.WriteString(scanner.Text())
		fileAsString.WriteByte('\n')
	}
	file.Close()
	return fileAsString.String()
}

func RemoveFirstLine(s string) string {
	i := strings.IndexByte(s, '\n')
	return s[i+1:]
}

func RemoveFirstWord(s string) string {
	i := strings.IndexByte(s, ' ')
	return s[i+1:]
}

func PrintArrayAsText(nums []int) {
	var text strings.Builder
	for i := range nums {
		text.WriteString(strconv.Itoa(nums[i]))
		if i < len(nums)-1 {
			text.WriteByte(' ')
		}
	}
	fmt.Println(text.String())
}
