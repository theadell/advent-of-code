package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

var chars = []rune{'#', '/', '?'}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer inputFile.Close()
	grid := make([][]rune, 0)
	sum := 0
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
					sum += d
				}
			}
		}
	}
	fmt.Println(sum)

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
