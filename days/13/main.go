package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var PART2_INCREASE = 10_000_000_000_000

const path = "./"

type Vector2 struct {
	X, Y int
}

func (v Vector2) Add(val int) Vector2 {
	v.X += val
	v.Y += val
	return v
}

func newVector2(values []string) Vector2 {
	xAsInt, _ := strconv.Atoi(values[0])
	yAsInt, _ := strconv.Atoi(values[1])

	return Vector2{
		X: xAsInt,
		Y: yAsInt,
	}
}

type Game struct {
	A     Vector2
	B     Vector2
	Prize Vector2
}

func getPrice(a, b int) int {
	return a*3 + b*1
}

func checkWinnings(game Game, part1 bool) int {
	tmpA := Vector2{
		X: game.A.X * game.B.Y,
		Y: game.A.Y * game.B.X,
	}
	tmpPrize := Vector2{
		X: game.Prize.X * game.B.Y,
		Y: game.Prize.Y * game.B.X,
	}

	a := (tmpPrize.X - tmpPrize.Y) / (tmpA.X - tmpA.Y)
	b := (game.Prize.X - (game.A.X * a)) / game.B.X

	if part1 && (a > 100 || b > 100) {
		return 0
	}

	if (game.A.X*a)+(game.B.X*b) == game.Prize.X &&
		(game.A.Y*a)+(game.B.Y*b) == game.Prize.Y {
		return getPrice(a, b)
	}

	return 0
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

	games := []Game{}
	s := bufio.NewScanner(file)
	s.Split(blockSplitFunc)
	for s.Scan() {
		games = append(games, handleInput(s.Bytes()))
	}

	result := 0
	result2 := 0
	for _, g := range games {
		result += checkWinnings(g, true)
	}

	for _, g := range games {
		g.Prize = g.Prize.Add(PART2_INCREASE)
		result2 += checkWinnings(g, false)
	}

	fmt.Println("Part1:", result)
	fmt.Println("Part2:", result2)
}

func blockSplitFunc(data []byte, atEOF bool) (int, []byte, error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		return i + 2, data[0:i], nil
	}

	if atEOF {
		return len(data), bytes.TrimSuffix(data, []byte("\n")), nil
	}

	return 0, nil, nil
}

func handleInput(data []byte) Game {
	game := Game{}
	re := regexp.MustCompile(`\d+`)

	lines := bytes.Split(data, []byte("\n"))
	for _, line := range lines {
		lineAsString := string(line)
		if strings.HasPrefix(lineAsString, "Button A:") {
			game.A = newVector2(re.FindAllString(lineAsString, -1))
		} else if strings.HasPrefix(lineAsString, "Button B:") {
			game.B = newVector2(re.FindAllString(lineAsString, -1))
		} else {
			game.Prize = newVector2(re.FindAllString(lineAsString, -1))
		}
	}

	return game
}
