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

func PrepareFileForWriting(path string) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	os.Truncate(path, 0)
}

func AppendLine(line string, path string) {
	f, err := os.OpenFile(path, os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(line + "\n")
	if err != nil {
		panic(err)
	}
}
func AppendLines(lines []string, path string) {
	f, err := os.OpenFile(path, os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, line := range lines {
		_, err = f.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}
}
