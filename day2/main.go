package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	filePath := "/Users/sanjeetyadav/dev/adventOfCode2023/advent-of-code-2023-go/day2/input.txt"
	lines, err := ReadLines(filePath)
	check(err)

	game := make(map[string]string)

	for _, s := range lines {
		gameValue := strings.Split(s, ":")
		game[gameValue[0]] = gameValue[1]
	}
	solvePart1(game)
	solvePart2(game)

}

func solvePart2(game map[string]string) {
	sumOfPowers := 0
	for _, v := range game {
		sets := ReadSets(v)
		maxR := -1
		maxG := -1
		maxB := -1
		for _, s := range sets {
			maxR = max(maxR, cubeCount(s)[0])
			maxG = max(maxG, cubeCount(s)[1])
			maxB = max(maxB, cubeCount(s)[2])
		}
		sumOfPowers += setPower(maxR, maxG, maxB)
	}
	fmt.Println(sumOfPowers)
}

func setPower(a, b, c int) int {
	return a * b * c
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solvePart1(game map[string]string) {
	validGames := []int{}

	for k, v := range game {
		sets := ReadSets(v)
		isValid := true
		for _, s := range sets {
			fmt.Printf("calling validate for game %s and set %s\n", k, s)
			isValid = isValid && validateSet(s)
			fmt.Printf("%t validate for game %s and set %s\n", isValid, k, s)
		}
		if isValid {
			g_num, e := strconv.Atoi(strings.Split(k, " ")[1])
			check(e)
			validGames = append(validGames, g_num)
			fmt.Printf("game %s and set %d\n", k, g_num)
		}
	}

	slices.Sort(validGames)

	ans := 0
	for _, n := range validGames {
		ans += n
	}

	fmt.Println(ans)

}

func cubeCount(cubes string) [3]int {
	cubeValues := strings.Split(cubes, ",")
	rgb := [3]int{0, 0, 0}
	for _, c := range cubeValues {
		countColor := strings.Split(strings.TrimSpace(c), " ")
		num, err := strconv.Atoi(strings.TrimSpace(countColor[0]))
		check(err)
		if countColor[1] == "green" {
			rgb[1] = num
		} else if countColor[1] == "blue" {
			rgb[2] = num
		} else {
			rgb[0] = num
		}
	}
	return rgb
}

func validateSet(cubes string) bool {
	cubeValues := strings.Split(cubes, ",")
	fmt.Println(cubeValues)

	isValid := true
	for _, c := range cubeValues {
		countColor := strings.Split(strings.TrimSpace(c), " ")
		num, err := strconv.Atoi(strings.TrimSpace(countColor[0]))
		check(err)
		fmt.Printf("validating.. %s with num : %d\n ", countColor[1], num)
		if countColor[1] == "green" {
			isValid = isValid && num <= 13
		} else if countColor[1] == "blue" {
			isValid = isValid && num <= 14
		} else {
			isValid = isValid && num <= 12
		}
		if !isValid {
			return false
		}
	}
	return isValid
}

func ReadSets(game string) []string {
	return strings.Split(game, ";")
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
