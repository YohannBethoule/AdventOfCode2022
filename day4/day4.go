package day4

import (
	"AdventOfCode/pkg"
	"fmt"
	"strconv"
	"strings"
)

func Day4Star1() {
	scan := pkg.ReadFile("./day4/input")

	total := 0

	for scan.Scan() {
		line := scan.Text()
		assignements := strings.Split(line, ",")
		assignement1 := strings.Split(assignements[0], "-")
		a1Lower, _ := strconv.Atoi(assignement1[0])
		a1Upper, _ := strconv.Atoi(assignement1[1])
		assignement2 := strings.Split(assignements[1], "-")
		a2Lower, _ := strconv.Atoi(assignement2[0])
		a2Upper, _ := strconv.Atoi(assignement2[1])

		if (a1Lower >= a2Lower && a1Upper <= a2Upper) || a2Lower >= a1Lower && a2Upper <= a1Upper {
			total++
		}
	}

	fmt.Printf("Total pairs fully containing : %d\n", total)
}

func Day4Star2() {
	scan := pkg.ReadFile("./day4/input")

	total := 0

	for scan.Scan() {
		line := scan.Text()
		assignements := strings.Split(line, ",")
		assignement1 := strings.Split(assignements[0], "-")
		a1Lower, _ := strconv.Atoi(assignement1[0])
		a1Upper, _ := strconv.Atoi(assignement1[1])
		assignement2 := strings.Split(assignements[1], "-")
		a2Lower, _ := strconv.Atoi(assignement2[0])
		a2Upper, _ := strconv.Atoi(assignement2[1])

		if (a1Lower >= a2Lower && a1Lower <= a2Upper) || (a1Upper >= a2Lower && a1Upper <= a2Upper) || (a2Lower >= a1Lower && a2Lower <= a1Upper) || (a2Upper >= a1Lower && a2Upper <= a1Upper) {
			total++
		}
	}

	fmt.Printf("Total pairs fully containing : %d\n", total)
}
