package day3

import (
	"AdventOfCode/pkg"
	"fmt"
	"strings"
)

func Day3Star1() {
	scan := pkg.ReadFile("./day3/input")
	priority := 0

	for scan.Scan() {
		bag := scan.Text()

		comp1 := bag[:(len(bag) / 2)]
		comp2 := bag[len(bag)/2:]

		fmt.Println("First half : " + comp1)
		fmt.Println("Second half : " + comp2)

		priority += analyzeBag1(comp1, comp2)
	}

	fmt.Printf("Total priority : %d", priority)
}

func analyzeBag1(comp1 string, comp2 string) int {
	for _, char1 := range comp1 {
		for _, char2 := range comp2 {
			if char1 == char2 {
				return getPriority(char1)
			}
		}
	}
	return 0
}

func Day3Star2() {
	scan := pkg.ReadFile("./day3/input")
	priority := 0

	for scan.Scan() {
		bag1 := scan.Text()
		scan.Scan()
		bag2 := scan.Text()
		scan.Scan()
		bag3 := scan.Text()

		priority += analyzeBag2(bag1, bag2, bag3)
	}

	fmt.Printf("Total priority : %d", priority)
}

func analyzeBag2(bag1 string, bag2 string, bag3 string) int {
	fmt.Println("------------------------------------------------------")
	fmt.Println(bag1)
	fmt.Println(bag2)
	fmt.Println(bag3)
	for _, char1 := range bag1 {
		for _, char2 := range bag2 {
			for _, char3 := range bag3 {
				if char1 == char2 && char2 == char3 {
					fmt.Println(getPriority(char1))
					fmt.Println("------------------------------------------------------")
					return getPriority(char1)
				}
			}

		}
	}
	return 0
}

func getPriority(char int32) int {
	str := "abcdefghijklmnopqrstuvwxyz"
	str += strings.ToUpper(str)
	return strings.IndexRune(str, char) + 1
}
