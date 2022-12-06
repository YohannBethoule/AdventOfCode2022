package day5

import (
	"AdventOfCode/pkg"
	"fmt"
	"github.com/golang-collections/collections/stack"
	"strconv"
	"strings"
)

func Day5() {
	numberOfStacks, maxNumberOfCrates := getNumberOfStacksAndCrates()
	var stacks []stack.Stack

	for i := 0; i < numberOfStacks; i++ {
		stacks = append(stacks, createStack(i, maxNumberOfCrates))
	}
	stacks = moves(stacks)
	for _, stack := range stacks {
		fmt.Print(stack.Peek())
	}
}

func getNumberOfStacksAndCrates() (int, int) {
	scan := pkg.ReadFile("./day5/input")

	numberOfCrates := 0

	for scan.Scan() {
		line := strings.TrimSpace(scan.Text())
		if _, err := strconv.Atoi(line[0:1]); err != nil {
			numberOfCrates++
			continue
		}
		return len(strings.ReplaceAll(line, " ", "")), numberOfCrates
	}

	return 0, 0
}

func createStack(index int, maxNumberOfCrates int) stack.Stack {
	scan := pkg.ReadFile("./day5/input")

	stack := stack.New()
	var crates []string

	for i := 0; i < maxNumberOfCrates; i++ {
		scan.Scan()
		line := scan.Text()
		crate := line[index*4 : index*4+3]
		if crate[1:2] == " " {
			continue
		}
		crates = append([]string{crate[1:2]}, crates...)
	}

	for _, c := range crates {
		stack.Push(c)
	}

	return *stack
}

func moves(stacks []stack.Stack) []stack.Stack {
	scan := pkg.ReadFile("./day5/input")

	for i := 0; i < 10; i++ {
		//skip 10 first lines
		scan.Scan()
	}

	for scan.Scan() {
		line := scan.Text()
		instruction := strings.Split(line, " ")
		nbCrates, _ := strconv.Atoi(instruction[1])
		from, _ := strconv.Atoi(instruction[3])
		to, _ := strconv.Atoi(instruction[5])

		moveByPack(stacks, nbCrates, from, to)
	}

	return stacks
}

// star1
func moveOneByOne(stacks []stack.Stack, nbCrates int, from int, to int) []stack.Stack {
	for i := 0; i < nbCrates; i++ {
		c := stacks[from-1].Pop()
		stacks[to-1].Push(c)
	}
	return stacks
}

// star2
func moveByPack(stacks []stack.Stack, nbCrates int, from int, to int) []stack.Stack {
	var crates []interface{}
	for i := 0; i < nbCrates; i++ {
		crates = append([]interface{}{stacks[from-1].Pop()}, crates...)
	}
	for _, c := range crates {
		stacks[to-1].Push(c)
	}
	return stacks
}
