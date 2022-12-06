package day1

import (
	"AdventOfCode/pkg"
	"fmt"
	"log"
	"sort"
	"strconv"
)

func Day1Star1And2() {
	scan := pkg.ReadFile("./day1/input")

	var elves []int

	elf := 0
	max := 0

	for scan.Scan() {
		if scan.Text() == "" {
			elves = append(elves, elf)
			if elf > max {
				max = elf
			}
			elf = 0
			continue
		}
		cal, err := strconv.Atoi(scan.Text())
		if err != nil {
			log.Fatal(err)
		}
		elf += cal
	}

	fmt.Printf("Most carrying elf calories : %d\n", max)

	sort.Ints(elves)
	total := 0
	for i := 1; i <= 3; i++ {
		total += elves[len(elves)-i]
	}

	fmt.Println(elves)
	fmt.Printf("Most carrying 3 elves calories : %d\n", total)
}

func Day1Star2() {
	scan := pkg.ReadFile("./day1/input")

	var elves []int

	elf := 0
	max := 0

	for scan.Scan() {
		if scan.Text() == "" {
			elves = append(elves, elf)
			if elf > max {
				max = elf
			}
			elf = 0
			continue
		}
		cal, err := strconv.Atoi(scan.Text())
		if err != nil {
			log.Fatal(err)
		}
		elf += cal
	}

	sort.Ints(elves)
	total := 0
	for i := 1; i <= 3; i++ {
		total += elves[len(elves)-i]
	}

	fmt.Println(elves)
}
