package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	score := 0
	for scanner.Scan() {
		score += overlap(scanner.Text())
	}

	fmt.Println(score)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func overlap(line string) int {
	pairs := strings.Split(line, ",")
	pair1 := strings.Split(pairs[0], "-")
	pair2 := strings.Split(pairs[1], "-")
	lowerBound1, _ := strconv.Atoi(pair1[0])
	upperBound1, _ := strconv.Atoi(pair1[1])
	lowerBound2, _ := strconv.Atoi(pair2[0])
	upperBound2, _ := strconv.Atoi(pair2[1])

	if upperBound1 < lowerBound2 {
		return 0
	}

	if upperBound2 < lowerBound1 {
		return 0
	}

	return 1
}
