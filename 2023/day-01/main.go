package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var Digits = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

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
		for i, r := range line {
			if r >= '0' && r <= '9' {
				digits += string(r)
			} else {
				for k, v := range Digits {
					if strings.HasPrefix(line[i:], k) {
						digits += strconv.Itoa(v)
						continue
					}
				}
			}
		}
		if len(digits) == 0 {
			log.Println("Line doesn't contain any digits: ", line)
			continue
		}
		d := string(digits[0]) + string(digits[len(digits)-1])
		num, err := strconv.Atoi(d)
		if err != nil {
			panic(err)
		}
		sum += num
	}
	fmt.Printf("Sum is: %d \n", sum)

}
