package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Solve() {
	file, _ := os.Open("./day1/input.txt")
	defer file.Close()
	scnr := bufio.NewScanner(file)
	maxCal := 0
	curCal := 0
	for scnr.Scan() {
		line := scnr.Text()
		if line == "" {
			if curCal > maxCal {
				maxCal = curCal
			}
			curCal = 0
		} else {
			val, _ := strconv.Atoi(line)
			curCal += val
		}
	}
	fmt.Println(maxCal)
}
