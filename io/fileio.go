package io

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(path string) []string {
	return ReadFileWithMaxLength(path, 65535)
}

func ReadFileWithMaxLength(path string, maxLength int) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	buffer := make([]byte, maxLength)

	scanner.Buffer(buffer, maxLength)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func SplitIntoSections(path string) [][]string {
	lines := ReadFile(path)
	var sections [][]string
	var section []string
	for _, line := range lines {
		if len(line) == 0 {
			if len(section) != 0 {
				sections = append(sections, section)
				section = make([]string, 0)
			}
		} else {
			section = append(section, line)
		}
	}
	if len(section) != 0 {
		sections = append(sections, section)
	}
	return sections
}
