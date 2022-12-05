package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	counts := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			counts = append(counts, count)
			count = 0
		} else {
			lineInt, _ := strconv.Atoi(line)
			count += lineInt
		}
	}

	sort.Ints(counts)

	fmt.Println("Solution 1: %i", counts[len(counts)-1])
	fmt.Println("Solution 2: %i", counts[len(counts)-1]+counts[len(counts)-2]+counts[len(counts)-3])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
