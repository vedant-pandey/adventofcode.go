package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	file, _ := os.Open("./day7/small.txt")
	defer file.Close()
	scnr := bufio.NewScanner(file)
	// soln := -1
	root := newTree()
	ptr := root
	// lsMode := false
	dirCnt := 0
	for scnr.Scan() {
		line := scnr.Text()
		words := strings.Split(line, " ")
		if words[0] == "$" {
			// lsMode = false
			if words[1] == "cd" {
				if words[2] == "/" {
					ptr = root
				} else {
					ptr = ptr.children[words[2]]
				}
			} else {
				// lsMode = true
			}
		} else {
			// if !lsMode {
			// 	panic("hey that's illegal")
			// }
			curChild := newTree()
			if words[0] == "dir" {
				curChild.isDir = true
				dirCnt++
			} else {
				curSize, _ := strconv.Atoi(words[0])
				curChild.isDir = false
				curChild.size = uint32(curSize)
			}
			ptr.addChild(curChild, words[1])
		}
	}
	dirMap := make(map[string]uint32, dirCnt)
	recursiveCalcSize(root, dirMap, "/")
  for k, v := range dirMap {
    fmt.Println(k,v)
  }
}

func recursiveCalcSize(node Tree, dirMap map[string]uint32, name string) {
	// if node == Tree{} {
	// 	return
	// }
	size := uint32(0)
	for childName, child := range node.children {
		if child.isDir {
			recursiveCalcSize(child, dirMap, childName)
      size += child.size
		} else {
			size += child.size
		}
	}
  node.size = size
  dirMap[name] = size
}

type Tree struct {
	children map[string]Tree
	size     uint32
	isDir    bool
}

func newTree() Tree {
  t := Tree{}
  t.children = map[string]Tree{}
  t.size = 0
  t.isDir = false
  return t
}

func (t Tree) addChild(child Tree, add string) {
  if t.children == nil {
    t.children = map[string]Tree{}
  }
	t.children[add] = child
}
