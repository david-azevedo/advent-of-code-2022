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
	file, err := os.Open("./day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	arrangement := map[int][]string{
		1: {"D", "L", "J", "R", "V", "G", "F"},
		2: {"T", "P", "M", "B", "V", "H", "J", "S"},
		3: {"V", "H", "M", "F", "D", "G", "P", "C"},
		4: {"M", "D", "P", "N", "G", "Q"},
		5: {"J", "L", "H", "N", "F"},
		6: {"N", "F", "V", "Q", "D", "G", "T", "Z"},
		7: {"F", "D", "B", "L"},
		8: {"M", "J", "B", "S", "V", "D", "N"},
		9: {"G", "L", "D"},
	}

	for scanner.Scan() {
		parseInstruction(scanner.Text(), arrangement)
	}

	printCode(arrangement)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parseInstruction(instruction string, arrangement map[int][]string) {
	// Example: move 3 from 4 to 6
	tokens := strings.Split(instruction, " ")
	a, s, d := tokens[1], tokens[3], tokens[5]
	amount, _ := strconv.Atoi(a)
	source, _ := strconv.Atoi(s)
	destination, _ := strconv.Atoi(d)

	/* Part 1
	for i := 0; i < amount; i++ {
		arrangement[destination] = append(arrangement[destination], arrangement[source][len(arrangement[source])-1])
		arrangement[source] = arrangement[source][:len(arrangement[source])-1]
	}
	*/

	// Part 2
	arrangement[destination] = append(arrangement[destination], arrangement[source][len(arrangement[source])-amount:]...)
	arrangement[source] = arrangement[source][:len(arrangement[source])-amount]
}

func printCode(arrangement map[int][]string) {
	for i := 1; i < 10; i++ {
		fmt.Print(arrangement[i][len(arrangement[i])-1])
	}
	fmt.Println()
}
