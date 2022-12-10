package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processInput(lines []string) *Dir {
	var (
		root    *Dir = NewDir("/", nil)
		current *Dir
	)
	for _, line := range lines {
		if line == "" {
			return root
		}
		if line[0] == '$' {
			if line[2:4] == "cd" {
				if line[5] == '/' {
					current = root
				} else {
					name := line[5:]
					if name == ".." {
						current = current.Parent
					} else {
						entry := current.Entries[name]
						current = entry.(*Dir)
					}
				}
			}
			// ls is a noop
		} else {
			parts := strings.Split(line, " ")
			var entry Entry
			if parts[0] == "dir" {
				entry = NewDir(parts[1], current)
			} else {
				size, _ := strconv.Atoi(parts[0])
				entry = &File{Size: size, Name: parts[1]}
			}
			current.Entries[parts[1]] = entry
		}
	}
	return root
}

var smallSum int

func addUpSmall(dir *Dir) {
	dirSize := dir.GetSize()
	//fmt.Printf("d %s: %d\n", dir.Name, dirSize)
	if dirSize <= 100000 {
		smallSum += dirSize
	}
}

func getInput(filename string) ([]string, error) {
	content := []string{}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return content, nil
}

type Finder struct {
	TargetSize   int
	SmallestSize int
	SmallestName string
}

func (f *Finder) find(dir *Dir) {
	size := dir.GetSize()
	if size >= f.TargetSize && size < f.SmallestSize {
		f.SmallestSize = size
		f.SmallestName = dir.Name
	}
}

func main() {
	inp, _ := getInput("input.txt")
	root := processInput(inp)
	//root.Print(0)

	// Part 1
	root.Walk(addUpSmall)
	fmt.Printf("Total Small Size: %d\n", smallSum)

	// Part 2
	totalSize := root.GetSize()
	fmt.Printf("Total size: %d\n", totalSize)
	freeSpace := 70000000 - totalSize
	fmt.Printf("Free sopace: %d\n", freeSpace)
	neededSpace := 30000000 - freeSpace
	fmt.Printf("Needed Space: %d\n", neededSpace)

	finder := Finder{
		TargetSize:   neededSpace,
		SmallestSize: 70000000,
	}
	root.Walk(finder.find)
	fmt.Printf("Smallest dir: %s (size=%d)\n", finder.SmallestName, finder.SmallestSize)
}
