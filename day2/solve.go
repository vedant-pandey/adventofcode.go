package day2

import (
	"bufio"
	"fmt"
	"os"
)

func Solve() {
	file, _ := os.Open("./day2/input.txt")
	defer file.Close()
	scnr := bufio.NewScanner(file)
	var score int64
	scMap := map[byte]int64{
		'X': 1,
		'Y': 2,
		'Z': 3,
	}
	winRevMap := map[byte]byte{
		'Z': 'X',
		'X': 'Y',
		'Y': 'Z',
	}
	for scnr.Scan() {
		line := scnr.Text()
		op, me := line[0], line[2]
		op += 'X' - 'A'
		if me == 'Z' {
			score += 6 + scMap[winRevMap[op]]
		} else if me == 'Y' {
			score += 3 + scMap[op]
		} else {
			score += scMap[winRevMap[winRevMap[op]]]
		}
	}
	fmt.Println(score)
}
