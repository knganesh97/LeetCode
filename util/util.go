package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

type InputScanner interface {
	ScanInputs(filePath string)
}

type OutputPrinter interface {
	CalculateOutputs(any)
	PrintOutputs()
}
