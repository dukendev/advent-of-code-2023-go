package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "/Users/sanjeetyadav/dev/adventOfCode2023/advent-of-code-2023-go/day2/input.txt"
	lines, err := ReadLines(filePath)
	check(err)
	fmt.Println(lines)
}

func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
