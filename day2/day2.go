package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var movePts = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var equivalentMoves = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
	"X": "A",
	"Y": "B",
	"Z": "C",
}

func Day2Star1() {
	file, err := os.Open("./day2/input")
	if err != nil {
		log.Fatal(err)
	}
	scan := bufio.NewScanner(file)

	score := 0

	for scan.Scan() {
		move := scan.Text()
		enemyMove := string(move[0])
		playerMove := string(move[2])

		score += movePts[playerMove] + getWinPoints(enemyMove, playerMove)
	}

	fmt.Printf("Score : %d", score)
}

func getWinPoints(enemyMove string, playerMove string) int {
	if enemyMove == equivalentMoves[playerMove] {
		return 3
	}
	if playerMove == "X" && enemyMove == "C" {
		return 6
	}
	if playerMove == "Y" && enemyMove == "A" {
		return 6
	}
	if playerMove == "Z" && enemyMove == "B" {
		return 6
	}
	return 0
}

func Day2Star2() {
	file, err := os.Open("./day2/input")
	if err != nil {
		log.Fatal(err)
	}
	scan := bufio.NewScanner(file)

	score := 0

	for scan.Scan() {
		move := scan.Text()
		enemyMove := string(move[0])
		output := string(move[2])

		playerMove := choseMove(output, enemyMove)

		score += movePts[playerMove] + getWinPoints(enemyMove, playerMove)
	}

	fmt.Printf("Score : %d", score)
}

func choseMove(output string, enemyMove string) string {
	if output == "Y" {
		return equivalentMoves[enemyMove]
	}
	switch enemyMove {
	case "A":
		if output == "Z" {
			return "Y"
		} else {
			return "Z"
		}
	case "B":
		if output == "Z" {
			return "Z"
		} else {
			return "X"
		}
	case "C":
		if output == "Z" {
			return "X"
		} else {
			return "Y"
		}
	}
	log.Panic("We shouldn't be here")
	return ""
}
