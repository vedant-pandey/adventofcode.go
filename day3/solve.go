package day3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Solve() {
  file, _ := os.Open("./day3/input.txt")
  defer file.Close()
  scnr := bufio.NewScanner(file)
  prio := map[byte]int{}
  var i byte
  for i = 'a'; i < 'z'; i++ {
    prio[i] = int(i) - int('a') + 1
  }
  for i = 'A'; i < 'Z'; i++ {
    prio[i] = int(i) - int('A') + 27
  }
  score := int64(0)
  j := 0
  for scnr.Scan() {
    line := scnr.Text()
    rightPart := line[len(line)/2:]
    for i := 0; i < len(line)/2; i++ {
      if strings.Contains(rightPart, string(line[i])) {
        score += int64(prio[line[i]])
        break
      }
    }
    j++
  }
  fmt.Println(score)
}
