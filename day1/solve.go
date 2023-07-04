package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Solve() {
	file, _ := os.Open("./day1/input.txt")
	defer file.Close()
	scnr := bufio.NewScanner(file)
	var calo []int
	curCal := 0
	for scnr.Scan() {
		line := scnr.Text()
		if line == "" {
			calo = append(calo, curCal)
			curCal = 0
		} else {
			val, _ := strconv.Atoi(line)
			curCal += val
		}
	}
	sort.Ints(calo)
	sum := calo[len(calo)-3] + calo[len(calo)-2] + calo[len(calo)-1]

	fmt.Println(sum)
}
