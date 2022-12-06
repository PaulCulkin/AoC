package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadLine(day string) ([]string, error) {
	if day == "" {
		return []string{}, errors.New("empty day")
	}

	filePath := fmt.Sprintf("%v/data.txt", day)
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanRunes)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines, nil
}

func ReadLines(day string) ([]string, error) {
	if day == "" {
		return []string{}, errors.New("empty day")
	}

	filePath := fmt.Sprintf("%v/data.txt", day)
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines, nil
}
