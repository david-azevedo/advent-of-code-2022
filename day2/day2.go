package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	score := 0
	for scanner.Scan() {
		score += calculateScore(scanner.Text())
	}

	fmt.Println("Score:", score)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

var ScorePoints = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

var LoseStrat = map[string]int{
	"A": 3,
	"B": 1,
	"C": 2,
}

var WinStrat = map[string]int{
	"A": 2,
	"B": 3,
	"C": 1,
}

var DrawStrat = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

func calculateScore(line string) int {
	plays := strings.Fields(line)
	opponent, outcome := plays[0], plays[1]

	score := ScorePoints[outcome]

	switch outcome {
	case "X":
		score += LoseStrat[opponent]
	case "Y":
		score += DrawStrat[opponent]
	case "Z":
		score += WinStrat[opponent]
	}

	return score
}
