package day3

import (
	"bufio"
	"fmt"
	"os"
)

func Solve() {
  file, _ := os.Open("./day3/input.txt")
  defer file.Close()
  scnr := bufio.NewScanner(file)
  sum := uint(0)
  for scnr.Scan() {
    line := scnr.Text()
    leftMap := map[byte]bool{}
    for i := 0; i < len(line)/2; i++ {
      leftMap[line[i]] = true
    }
    for i := len(line)/2; i < len(line); i++ {
      if (leftMap[line[i]]) {
        sum += priority(line[i])
        break
      }
    }
  }
  fmt.Println(sum)
}

func priority(item byte) uint {
	if item >= 'a' && item <= 'z' {
		return uint(item-'a') + 1
	}
	return uint(item-'A') + 27
}
