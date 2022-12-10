package main

import (
	"bufio"
	"fmt"
	"os"
)

var prioSum int

func charToPrio(char rune) int {
	if char >= 'a' && char <= 'z' {
		return int(char-'a') + 1
	}
	if char >= 'A' && char <= 'Z' {
		return int(char-'A') + 27
	}
	return 0
}

func makeSet(line string) map[int]bool {
	prioSet := make(map[int]bool)
	for _, char := range line {
		prioSet[charToPrio(char)] = true
	}
	return prioSet
}

func calcSum(lines []string) error {
	counts := make(map[int]int)
	for _, line := range lines {
		for prio, _ := range makeSet(line) {
			counts[prio] += 1
		}
	}
	for prio, count := range counts {
		if count == 3 {
			prioSum += prio
			return nil
		}
	}
	return nil
}

func processInput(filename string, processor func([]string) error) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) == 3 {
			err = processor(lines)
			if err != nil {
				return err
			}
			lines = []string{}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	err := processInput("input.txt", calcSum)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf("Sum: %d\n", prioSum)
}
