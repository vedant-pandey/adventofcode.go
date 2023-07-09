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
		solve2(line, soln)
	}
}

func solve2(line string, soln int) {
	wrd := []byte{line[0], line[1], line[2], line[3], line[4], line[5], line[6], line[7], line[8], line[9], line[10], line[11], line[12], line[13]}
	for i := len(wrd); i < len(line); i++ {
		if checkUnique(&wrd) {
			soln = i - 1
			break
		}
		moveIt(&wrd, line[i])
	}
	if soln == -1 {
		soln = len(line)
	}
	fmt.Println(soln + 1)
}

func solve(line string, soln int) {
	wrd := []byte{line[0], line[1], line[2], line[3]}
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
	fmt.Println(soln + 1)
}

func moveIt(wrd *[]byte, c byte) {
	for i := 0; i < len(*wrd)-1; i++ {
		(*wrd)[i] = (*wrd)[i+1]
	}
	(*wrd)[len(*wrd)-1] = c
}

func checkUnique(wrd *[]byte) bool {
	myMap := make(map[byte]bool, len(*wrd))
	for _, ch := range *wrd {
		if myMap[ch] {
			return false
		}
		myMap[ch] = true
	}
	return true
}
