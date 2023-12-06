package utils

import (
	"bufio"
	"os"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(path string) (fileLines []string) {

	readFile, err := os.Open(path)
	CheckError(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	return fileLines
}
