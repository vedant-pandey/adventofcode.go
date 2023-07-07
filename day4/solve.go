package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() {
  file, _ := os.Open("./day4/input.txt")
  defer file.Close()
  scnr := bufio.NewScanner(file)
  soln := 0
  for scnr.Scan() {
    line := scnr.Text()
    elves := strings.Split(line, ",")
    elveRanges := [][]string{strings.Split(elves[0], "-"), strings.Split(elves[1], "-")}
    fSt, _ := strconv.Atoi(elveRanges[0][0])
    fEn, _ := strconv.Atoi(elveRanges[0][1])
    sSt, _ := strconv.Atoi(elveRanges[1][0])
    sEn, _ := strconv.Atoi(elveRanges[1][1])
    if (fSt >= sSt && fEn <= sEn) || (sSt >= fSt && sEn <= fEn) {
      soln += 1
    }
  }
  fmt.Println(soln)
}

