package day6

import (
	"bufio"
	"fmt"
	"os"
)

func Solve() {
	file, _ := os.Open("./day6/input.txt")
	defer file.Close()
	scnr := bufio.NewScanner(file)
	soln := -1
	for scnr.Scan() {
		line := scnr.Text()
		wrd := [...]byte{line[0], line[1], line[2], line[3]}
		for i := 4; i < len(line); i++ {
			if checkUnique(&wrd) {
				soln = i - 1
				break
			}
			moveIt(&wrd, line[i])
		}
		if soln == -1 {
			soln = len(line)
		}
    fmt.Println(soln+1)
	}
}

func moveIt(wrd *[4]byte, c byte) {
	wrd[0] = wrd[1]
	wrd[1] = wrd[2]
	wrd[2] = wrd[3]
	wrd[3] = c
}

func checkUnique(wrd *[4]byte) bool {
	myMap := make(map[byte]bool, 4)
	for _, ch := range wrd {
		if myMap[ch] {
			return false
		}
		myMap[ch] = true
	}
	return true
}
