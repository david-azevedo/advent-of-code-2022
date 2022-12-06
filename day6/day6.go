package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		chars := strings.Split(scanner.Text(), "")
		processData(chars)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func processData(data []string) {
	// uniques := 4 // Part 1
	uniques := 14 // Part 2
	lastSeen := make([]string, 0, uniques)
	counter := 0
	for _, char := range data {
		counter++
		seen := false
		for index, v := range lastSeen {
			if char == v {
				lastSeen = lastSeen[index+1:]
				lastSeen = append(lastSeen, char)
				seen = true
				break
			}
		}

		if !seen {
			lastSeen = append(lastSeen, char)
		}

		if len(lastSeen) == uniques {
			fmt.Println(counter)
			return
		}
	}
}
