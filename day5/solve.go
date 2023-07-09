package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	file, _ := os.Open("./day5/input.txt")
	defer file.Close()
	scnr := bufio.NewScanner(file)
	stacks := [][]string{}
	lines := []string{}
	for scnr.Scan() {
		line := scnr.Text()
		if len(line) == 0 {
			stacks = buildStack(lines[:len(lines)-1])
			lines = nil
		} else {
			lines = append(lines, line)
		}
	}
	soln := doCalculation2(stacks, lines)
	fmt.Println(soln)
}

func doCalculation2(stacks [][]string, lines []string) string {
	for _, line := range lines {
		words := strings.Split(line, " ")
		cnt, _ := strconv.Atoi(words[1])
		srcInd, _ := strconv.Atoi(words[3])
		dstInd, _ := strconv.Atoi(words[5])
		srcInd--
		dstInd--
		src := stacks[srcInd]
		dst := stacks[dstInd]
		pushRange := src[len(src)-cnt:]
		src = src[:len(src)-cnt]
		dst = append(dst, pushRange...)
    stacks[srcInd] = src
    stacks[dstInd] = dst
	}
	soln := ""
	for _, stack := range stacks {
		soln += string(stack[len(stack)-1])
	}
	return soln

}

func doCalculation(stacks [][]string, lines []string) string {
	for _, line := range lines {
		words := strings.Split(line, " ")
		cnt, _ := strconv.Atoi(words[1])
		src, _ := strconv.Atoi(words[3])
		dst, _ := strconv.Atoi(words[5])
		for i := 0; i < cnt; i++ {
			val := stacks[src-1][len(stacks[src-1])-1]
			stacks[src-1] = stacks[src-1][:len(stacks[src-1])-1]
			stacks[dst-1] = append(stacks[dst-1], val)
		}
	}
	soln := ""
	for _, stack := range stacks {
		soln += string(stack[len(stack)-1])
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
				stacks[i] = append([]string{string(line[4*i+1])}, stacks[i]...)
			}
		}
	}
	return stacks[:len(stacks)-1]
}
