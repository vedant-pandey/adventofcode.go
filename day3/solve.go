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
  prio := map[byte]int{}
  var i byte
  for i = 'a'; i < 'z'; i++ {
    prio[i] = int(i) - int('a') + 1
  }
  for i = 'A'; i < 'Z'; i++ {
    prio[i] = int(i) - int('A') + 27
  }
  sum := uint(0)
  for scnr.Scan() {
    rucksack := scnr.Text()
		var bitSet uint64
		for i, item := range rucksack {
			prio := priority(item)
			mask := uint64(1) << prio
			if i < len(rucksack)/2 {
				bitSet |= mask
			} else if bitSet&mask == mask {
				sum += prio
				bitSet &^= mask 
			}
		}
  }
  fmt.Println(sum)
}

func priority(item rune) uint {
	if item >= 'a' && item <= 'z' {
		return uint(item-'a') + 1
	}
	return uint(item-'A') + 27
}
