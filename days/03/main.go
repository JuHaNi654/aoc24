package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

const path = "./"

func part1(input []byte) [][]int {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	data := re.FindAll(input, -1)
	list := [][]int{}
	for _, i := range data {
		group := []int{}
		number := ""
		for _, r := range i {
			if r >= 48 && r <= 57 {
				number += string(r)
			}

			if r == 44 || r == 41 {
				asInt, _ := strconv.Atoi(number)
				group = append(group, asInt)
				number = ""
			}
		}
		list = append(list, group)
	}

	return list
}

var active = true

func part2(input []byte) [][]int {
	list := [][]int{}
	re := regexp.MustCompile(`mul\(\d+,\d+\)|don't\(\)|do\(\)`)
	data := re.FindAll(input, -1)
	for _, i := range data {
		if string(i) == "don't()" {
			active = false
      continue
		}

		if string(i) == "do()" {
			active = true
      continue
		}

		if !active {
			continue
		}

		group := []int{}
		number := ""
		for _, r := range i {
			if r >= 48 && r <= 57 {
				number += string(r)
			}

			if r == 44 || r == 41 {
				asInt, _ := strconv.Atoi(number)
				group = append(group, asInt)
				number = ""
			}
		}
		list = append(list, group)
	}

	return list
}

func count(list [][]int) {
	var result int
	for _, g := range list {
		tmp := 1
		for _, i := range g {
			tmp *= i
		}

		result += tmp
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
	var data [][]int
	var data2 [][]int
	for s.Scan() {
		data = append(data, part1(s.Bytes())...)
		data2 = append(data2, part2(s.Bytes())...)
	}

	count(data)
	count(data2)
	// Add code
}
