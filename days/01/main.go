package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const path = "./"

func parseLine(s string) (int, int) {
	values := strings.Split(s, "   ")
	l, _ := strconv.Atoi(values[0])
	r, _ := strconv.Atoi(values[1])
	return l, r
}

func part1(l, r []int) {
	sort.Ints(l)
	sort.Ints(r)
	var result int

	for i, v := range l {
		distance := v - r[i]
		if distance < 0 {
			result -= distance
			continue
		}

		result += distance
	}

	fmt.Println("Result: ", result)
}

func part2(l, r []int) {
	var result int

	for _, v := range l {
		repeat := 0
		for _, v2 := range r {
			if v == v2 {
				repeat++
			}
		}

		result += (v * repeat)
	}

	fmt.Println("Result: ", result)
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
	leftCol := []int{}
	rightCol := []int{}
	for s.Scan() {
		l, r := parseLine(s.Text())
		leftCol = append(leftCol, l)
		rightCol = append(rightCol, r)
	}

	part1(leftCol, rightCol)
	part2(leftCol, rightCol)
}
