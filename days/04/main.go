package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

type Pos struct {
	X, Y int
}

func (p Pos) Reverse() Pos {
	return Pos{X: -(p.X), Y: -(p.Y)}
}

const path = "./"

var word = []string{"X", "M", "A", "S"}
var directions = []Pos{
	{X: -1, Y: 1},  // Up - Left
	{X: 0, Y: 1},   // Up
	{X: 1, Y: 1},   // Up - Right
	{X: 1, Y: 0},   // Right
	{X: 1, Y: -1},  // Down - Right
	{X: 0, Y: -1},  // Down
	{X: -1, Y: -1}, // Down - Left
	{X: -1, Y: 0},  // Left
}

var xmasDirections = []Pos{
	{X: -1, Y: 1},  // Up - Left {x: 1, y: -1}
	{X: 1, Y: 1},   // Up - Right {x: -1, y: -1}
	{X: 1, Y: -1},  // Down - Right {x: -1, y: 1}
	{X: -1, Y: -1}, // Down - Left {x: 1, y: 1}
}

func checkWord(
	table [][]string,
	pos Pos,
	dir Pos,
	save []string,
	wordIndex int,
	counter *int,
) {
	if pos.X < 0 || pos.Y < 0 {
		return
	}

	if pos.X >= len(table[0]) || pos.Y >= len(table) {
		return
	}

	if table[pos.Y][pos.X] != word[wordIndex] {
		return
	}

	// Save value
	save = append(save, table[pos.Y][pos.X])

	// Check if we have full word
	if reflect.DeepEqual(word, save) {
		*counter += 1
		return
	}

	newPos := Pos{X: pos.X + dir.X, Y: pos.Y + dir.Y}
	checkWord(table, newPos, dir, save, wordIndex+1, counter)
}

func part1(table [][]string) {
	counter := 0
	for y, row := range table {
		for x, col := range row {
			if col == word[0] {
				for _, dir := range directions {
					checkWord(table, Pos{X: x, Y: y}, dir, []string{}, 0, &counter)
				}
			}
		}
	}
	fmt.Println("Result: ", counter)
}

func check2(table [][]string, pos, pos2 Pos) bool {
	if pos.X < 0 || pos.Y < 0 || pos2.X < 0 || pos2.Y < 0 {
		return false
	}

	if pos.X >= len(table[0]) ||
		pos.Y >= len(table) ||
		pos2.X >= len(table[0]) ||
		pos2.Y >= len(table) {
		return false
	}

	if table[pos.Y][pos.X] == "M" && table[pos2.Y][pos2.X] == "S" {
		if table[pos.Y][pos2.X] == "M" && table[pos2.Y][pos.X] == "S" {
			return true
		}

		if table[pos2.Y][pos.X] == "M" && table[pos.Y][pos2.X] == "S" {
			return true
		}

		return false
	}

	return false
}

func part2(table [][]string) {
	counter := 0
	for y, row := range table {
		for x, col := range row {
			if col == "A" {
				for _, dir := range xmasDirections {
					mirrored := dir.Reverse()
					if check2(
						table,
						Pos{X: x + dir.X, Y: y + dir.Y},
						Pos{X: x + mirrored.X, Y: y + mirrored.Y},
					) {
						counter++
						break
					}
				}
			}
		}
	}

	fmt.Println("Part2: ", counter)
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
	table := [][]string{}
	for s.Scan() {
		row := strings.Split(s.Text(), "")
		table = append(table, row)
	}

	part1(table)
	part2(table)
}
