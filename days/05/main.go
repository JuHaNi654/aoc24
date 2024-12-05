package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

const path = "./"

func validate(rules map[int][]int, updates [][]int) {
	result := 0
	result2 := 0
	invalidUpdates := [][]int{}

	for _, update := range updates {
		isValidUpdate := true
		for i, value := range update {
			rest := update[i+1:]
			if len(rest) != 0 {
				list, _ := rules[value]
				for _, j := range rest {
					if !slices.Contains(list, j) {
						isValidUpdate = false
						break
					}
				}
			}
		}

		if isValidUpdate {
			result += update[len(update)/2]
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	fmt.Println("Part 1:", result)

	for _, update := range invalidUpdates {
		i := 0
		for i < len(update) {
			value := update[i]
			rest := update[i+1:]
			list, _ := rules[value]
			rerun := false

			if len(rest) != 0 {
				for j, val := range rest {
					if !slices.Contains(list, val) {
						update[i] = val
						update[i+j+1] = value
						rerun = true

						break
					}
				}
			}
			if !rerun {
				i++
			}
		}

		result2 += update[len(update)/2]
	}

	fmt.Println("Part 2:", result2)
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

	rules := make(map[int][]int)
	updates := [][]int{}
	handlingRules := true

	s := bufio.NewScanner(file)
	for s.Scan() {
		// Assume that we are switching handling updates after first empty
		// line
		if len(s.Text()) == 0 {
			handlingRules = false
			continue
		}

		if handlingRules { // Parse rules
			values := strings.Split(s.Text(), "|")
			from, _ := strconv.Atoi(values[0])
			to, _ := strconv.Atoi(values[1])

			rule, ok := rules[from]
			if ok {
				rule = append(rule, to)
				rules[from] = rule
			} else {
				rules[from] = []int{to}
			}

		} else { // Parse updates
			values := strings.Split(s.Text(), ",")
			update := []int{}
			for _, v := range values {
				valAsInt, _ := strconv.Atoi(v)
				update = append(update, valAsInt)
			}

			updates = append(updates, update)
		}
	}

	validate(rules, updates)
}
