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
	winMap := map[byte]byte{
		'X': 'Z',
		'Y': 'X',
		'Z': 'Y',
	}
	for scnr.Scan() {
		line := scnr.Text()
		op, me := line[0], line[2]
		op += 'X' - 'A'
		score += scMap[me]
		if winMap[me] == op {
			score += 6
		} else if me == op {
			score += 3
		}
	}
	fmt.Println(score)
}
