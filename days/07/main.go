package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Kind string

var (
	None     Kind = ""
	Add      Kind = "+"
	Multiply Kind = "*"
	Combine  Kind = "||"
)

type Node struct {
	Kind    Kind
	Value   int
	Sum     *Node
	Mult    *Node
	Combine *Node
}

type Calibration struct {
	Result int
	Values []int
}

func newCalibration(result int, values []int) *Calibration {
	return &Calibration{
		Result: result,
		Values: values,
	}
}

const path = "./"

func generateTree(level int, list []int, kind Kind) *Node {
	if level >= len(list) {
		return nil
	}

	node := &Node{Kind: kind, Value: list[level]}
	node.Sum = generateTree(level+1, list, Add)
	node.Mult = generateTree(level+1, list, Multiply)
	node.Combine = generateTree(level+1, list, Combine)
	return node
}

func walkTree(node *Node, result, expected int) bool {
	if node == nil {
		return result == expected
	}

	if node.Kind == None {
		result = node.Value
	}

	if node.Kind == Add {
		result += node.Value
	}

	if node.Kind == Multiply {
		result *= node.Value
	}

	if node.Kind == Combine {
		combined := fmt.Sprintf("%d%d", result, node.Value)
		result, _ = strconv.Atoi(combined)
	}

	return (walkTree(node.Sum, result, expected) || walkTree(node.Mult, result, expected) || walkTree(node.Combine, result, expected))
}

func part1(calibrations []*Calibration) {
	part1 := 0
	for _, c := range calibrations {
		tree := generateTree(0, c.Values, None)
		if walkTree(tree, 0, c.Result) {
			part1 += c.Result
		}
	}

	fmt.Println("Part 1:", part1)
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
	calibrations := []*Calibration{}
	for s.Scan() {
		split := strings.Split(s.Text(), ":")
		resultAsInt, _ := strconv.Atoi(split[0])
		values := []int{}
		tmp := strings.Trim(split[1], " ")
		for _, v := range strings.Split(tmp, " ") {
			vAsInt, _ := strconv.Atoi(v)
			values = append(values, vAsInt)
		}

		calibrations = append(calibrations, newCalibration(resultAsInt, values))
	}

	part1(calibrations)
}
