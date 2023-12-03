package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()
		digits := ""
		for _, r := range line {
			if r >= '0' && r <= '9' {
				digits += string(r)
			}
		}
		if len(digits) == 0 {
			log.Println("Line doesn't contain any digits: ", line)
			continue
		}
		num, err := strconv.Atoi(string(digits[0]) + string(digits[len(digits)-1]))
		if err != nil {
			panic(err)
		}
		sum += num
	}
	fmt.Printf("Sum is: %d \n", sum)

}
