package day6

import (
	"AdventOfCode/pkg"
	"fmt"
	"strings"
)

func Day6Star1() {
	scan := pkg.ReadFile("./day6/input")
	scan.Scan()
	stream := scan.Text()

	//star1
	fmt.Printf("Start position of packet = %d\n", getStartOfPacket(stream, 4))
	//star2
	fmt.Printf("Start position of packet = %d\n", getStartOfPacket(stream, 14))
}

func getStartOfPacket(stream string, nbDigits int) int {
	for i := 0; i < len(stream)-nbDigits; i++ {
		chars := stream[i : i+nbDigits]
		right := true
		for j := 0; j < len(chars)-1; j++ {
			if strings.Count(chars, chars[j:j+1]) > 1 {
				right = false
			}
		}
		if right {
			return i + nbDigits
		}
	}
	return -1
}
