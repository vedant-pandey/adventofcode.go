package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	file, _ := os.Open("./day5/small.txt")
	defer file.Close()
	scnr := bufio.NewScanner(file)
	stacks := [][]string{}
	lines := []string{}
	/*

	  1 -> 0 , 2
	  2 -> 4 , 6

	*/
	for scnr.Scan() {
		line := scnr.Text()
		if len(line) == 0 {
			stacks = buildStack(lines[:len(lines)-1])
			lines = nil
		} else {
			lines = append(lines, line)
		}
	}
	fmt.Println(stacks)
	soln := doCalculation(stacks, lines)
	fmt.Println(soln)
}

func doCalculation(stacks [][]string, lines []string) string {
	for _, line := range lines {
		words := strings.Split(line, " ")
		cnt, _ := strconv.Atoi(words[1])
		src, _ := strconv.Atoi(words[3])
		dst, _ := strconv.Atoi(words[5])
		for i := 0; i < cnt; i++ {
      fmt.Println(stacks)
			val := stacks[src - 1][0]
			stacks[src - 1] = stacks[src - 1][1:]
			stacks[dst - 1] = append(stacks[dst - 1], val)
		}
	}
	soln := ""
	for _, stack := range stacks {
		soln += string(stack[0])
	}
	return soln
}

func buildStack(lines []string) [][]string {
	var stacks [][]string
	for i := 0; i <= len(lines[0])+1; i += 4 {
		stacks = append(stacks, []string{})
	}
	for _, line := range lines {
		for i := 0; 4*i < len(line); i++ {
			if line[4*i] == '[' {
				stacks[i] = append(stacks[i], string(line[4*i+1]))
			}
		}
	}
  return stacks[:len(stacks)-1]
}
