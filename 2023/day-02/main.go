package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Id       int
	ColorSet map[string]int
}

func (g *Game) Set(k, v string) {
	vint, err := strconv.Atoi(v)
	if err != nil {
		log.Fatal(err.Error())
	}
	count, _ := g.ColorSet[k]
	if vint > count {
		g.ColorSet[k] = vint
	}
}
func (g *Game) Valid() bool {
	if g.ColorSet["red"] <= 12 && g.ColorSet["green"] <= 13 && g.ColorSet["blue"] <= 14 {
		return true
	}
	return false
}
func (g *Game) Power() int {
	return g.ColorSet["red"] * g.ColorSet["green"] * g.ColorSet["blue"]
}
func NewGame(line string) Game {
	tokens := strings.Split(line, ":")
	if len(tokens) < 2 {
		log.Fatalf("invalid input: %s\n", line)
	}

	fields := strings.Fields(tokens[0])
	id, err := strconv.Atoi(fields[1])
	if err != nil {
		log.Fatalf("Failed to parse game id from tokens: %s \n", tokens[0])
	}
	game := Game{Id: id, ColorSet: map[string]int{}}

	cubeSet := strings.ReplaceAll(tokens[1], ";", ",")
	for _, colorCount := range strings.Split(cubeSet, ",") {
		fields := strings.Fields(colorCount)
		if len(fields) != 2 {
			log.Fatalf("Invalid set count %v at line: %s \n", fields, line)
		}
		game.Set(fields[1], fields[0])
	}
	return game
}

type CubeSet map[string]int

func main() {
	ans := 0
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		game := NewGame(line)
		ans += game.Power()
	}
	fmt.Println(ans)
}
