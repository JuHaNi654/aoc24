package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const path = "./"

type Direction string

var (
	Asc  Direction = "asc"
	Desc Direction = "desc"
)

func isValid(v1, v2 int) bool {
	result := v1 - v2
	if result >= 1 && result <= 3 {
		return true
	}

	return false
}

func isSafe(r []int, curr, next int, skippable bool, dir Direction) bool {
	if next >= len(r) {
		return true
	}

	var valid bool
	if dir == Asc {
		valid = isValid(r[next], r[curr])
	} else {
		valid = isValid(r[curr], r[next])
	}

	if !valid && !skippable {
		return false
	}

	if !valid && skippable && curr-1 >= 0 {
		if isSafe(r, curr-1, next, false, dir) {
			return isSafe(r, curr+1, next+1, skippable, dir)
		}
	}

	if valid {
		next += 1
		curr = next - 1
		return isSafe(r, curr, next, skippable, dir)
	}

	if curr == 0 && !valid && isSafe(r, curr+1, next+1, false, dir) {
		return true
	}

	return isSafe(r, curr, next+1, false, dir)
}

func part1(rows [][]int) {
	result := 0
	for _, r := range rows {
		if isSafe(r, 0, 1, true, Asc) || isSafe(r, 0, 1, true, Desc) {
			result++
		}
	}

	fmt.Println("Result:", result)
}

func main() {
	pwd, _ := os.Getwd()
	filename := os.Args[1]

	file, err := os.Open(filepath.Join(pwd, path, filename))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	levels := [][]int{}
	for s.Scan() {
		levels = append(levels, parseRow(s.Text()))
	}

	part1(levels)
}

func parseRow(r string) []int {
	level := []int{}

	for _, v := range strings.Split(r, " ") {
		vAsInt, _ := strconv.Atoi(v)
		level = append(level, vAsInt)
	}

	return level
}
