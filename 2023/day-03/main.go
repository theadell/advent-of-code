package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"unicode"
)

type Pos struct {
	x, y int
}
type Num struct {
	pos     Pos
	ndigits int
}

func main() {
	inputFile, err := os.Open("input.txt")
	partNumers := map[Pos]int{}
	if err != nil {
		log.Fatal(err.Error())
	}
	defer inputFile.Close()
	grid := make([][]rune, 0)
	sum := 0
	gearRatios := 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	height, width := len(grid), len(grid[0])
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r := grid[y][x]
			if unicode.IsDigit(r) {
				start := Pos{x: x, y: y}
				nx := x + 1
				for nx < width && unicode.IsDigit(grid[y][nx]) {
					nx++
				}
				x = nx - 1
				n := Num{pos: start, ndigits: nx - start.x}
				if IsPartNumber(n, grid) {
					d, err := strconv.Atoi(string(grid[y][start.x:nx]))
					if err != nil {
						log.Fatalf("Failed to parse as number. Num %v, %s \n", n, err.Error())
					}
					// Store each position of the part number
					for i := start.x; i <= x; i++ {
						partNumers[Pos{x: i, y: y}] = d
					}
					sum += d
				}
			}
		}
	}

	for y, row := range grid {
		for x, r := range row {
			if r == '*' {
				adjacentPartNumbers := make([]int, 0)
				adjPos := getAdjacentPositions(x, y, width, height)

				for _, pos := range adjPos {
					if num, exists := partNumers[pos]; exists {

						if !slices.Contains(adjacentPartNumbers, num) {
							adjacentPartNumbers = append(adjacentPartNumbers, num)
						}
					}
				}
				if len(adjacentPartNumbers) == 2 {
					// Multiply the two part numbers to get the gear ratio
					gearRatios += adjacentPartNumbers[0] * adjacentPartNumbers[1]
				}
			}
		}

	}

	fmt.Println(sum)
	fmt.Println(gearRatios)
}

func contains(slice []int, element int) bool {
	for _, el := range slice {
		if el == element {
			return true
		}
	}
	return false
}
func IsPartNumber(n Num, grid [][]rune) bool {
	x, y, l := n.pos.x, n.pos.y, n.ndigits

	for i := 0; i < l; i++ {
		if x+i >= len(grid[0]) {
			break
		}

		if y > 0 && isSymbol(grid[y-1][x+i]) {
			return true
		}
		if y < len(grid)-1 && isSymbol(grid[y+1][x+i]) {
			return true
		}
	}

	if x > 0 {
		if isSymbol(grid[y][x-1]) {
			return true
		}
		if y > 0 && isSymbol(grid[y-1][x-1]) {
			return true
		}
		if y < len(grid)-1 && isSymbol(grid[y+1][x-1]) {
			return true
		}
	}

	if x+l < len(grid[0]) {
		if isSymbol(grid[y][x+l]) {
			return true
		}
		if y > 0 && isSymbol(grid[y-1][x+l]) {
			return true
		}
		if y < len(grid)-1 && isSymbol(grid[y+1][x+l]) {
			return true
		}
	}

	return false
}
func isSymbol(r rune) bool {
	return !(r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' || r >= '0' && r <= '9' || r == '.')
}

func getAdjacentPositions(x, y, width, height int) []Pos {
	var positions []Pos

	// Directions: top, bottom, left, right, and diagonals
	dx := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	dy := []int{-1, -1, -1, 0, 0, 1, 1, 1}

	for i := 0; i < 8; i++ {
		adjX, adjY := x+dx[i], y+dy[i]
		if adjX >= 0 && adjX < width && adjY >= 0 && adjY < height {
			positions = append(positions, Pos{x: adjX, y: adjY})
		}
	}

	return positions
}
