package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	sum := 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, ":")
		if len(lineParts) != 2 {
			log.Fatal("invalid input")
		}
		game := strings.Split(lineParts[1], "|")
		if len(game) != 2 {
			log.Fatal("Invalid input")
		}
		cardStr, winningStr := game[0], game[1]
		card := make([]int, 0)
		for _, cardNumber := range strings.Fields(cardStr) {
			num, err := strconv.Atoi(cardNumber)
			if err != nil {
				log.Fatal(err.Error())
			}
			card = append(card, num)
		}
		winningSet := make([]int, 0)
		for _, winningNum := range strings.Fields(winningStr) {
			num, err := strconv.Atoi(winningNum)
			if err != nil {
				log.Fatal(err.Error())
			}
			winningSet = append(winningSet, num)
		}
		winningSet = unique(winningSet)
		numOfWinningCards := 0
		for _, winningNumber := range winningSet {
			if slices.Contains(card, winningNumber) {
				numOfWinningCards++
			}
		}
		if numOfWinningCards > 0 {
			sum += int(math.Pow(2, float64(numOfWinningCards-1)))
		}
	}
	fmt.Println(sum)
}

func unique[S ~[]E, E comparable](s S) S {
	uniqueMap := make(map[E]struct{})
	var result S

	for _, v := range s {
		if _, exists := uniqueMap[v]; !exists {
			uniqueMap[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}
