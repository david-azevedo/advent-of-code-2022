package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rucksacks := make([][]string, 0)
	for scanner.Scan() {
		rucksacks = append(rucksacks, strings.Split(scanner.Text(), ""))
	}

	score := 0
	for i := 0; i < len(rucksacks); i += 3 {
		elf1, elf2, elf3 := rucksacks[i], rucksacks[i+1], rucksacks[i+2]
		common := intersect(elf1, elf2)
		common = intersect(common, elf3)
		score += decode(common[0])
	}

	fmt.Println("Score:", score)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func intersect(part1, part2 []string) []string {
	result := make([]string, 0)
	for _, p1 := range part1 {
		for _, p2 := range part2 {
			if p1 == p2 {
				result = append(result, p1)
			}
		}
	}

	return result
}

func decode(letter string) int {
	// A - 65 Z - 90  -----> 27 - 52
	// a - 97 z - 122 -----> 1 - 26

	number := int(letter[0])
	if 65 <= number && number <= 90 {
		number -= 38
	} else {
		number -= 96
	}
	return number
}
