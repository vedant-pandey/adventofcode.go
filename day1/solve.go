package day1

import (
	"bufio"
	"fmt"
	"os"
)

func Solve() {
	file, _ := os.OpenFile("input.txt", os.O_RDONLY, 0644)
	defer file.Close()
	scnr := bufio.NewScanner(file)
	for scnr.Scan() {
		fmt.Println(scnr.Text())
	}

}
