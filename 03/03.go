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

func calcSum(line string) error {
	center := len(line) / 2
	firstSet := make(map[int]bool)
	secondSet := make(map[int]bool)
	for i := 0; i < center; i++ {
		firstSet[charToPrio(rune(line[i]))] = true
		secondSet[charToPrio(rune(line[i+center]))] = true
	}
	for prio := range firstSet {
		fmt.Printf("%d,", prio)
		if _, ok := secondSet[prio]; ok {
			fmt.Printf(": %d\n", prio)
			prioSum += prio
			return nil
		}
	}
	fmt.Println()
	return nil
}

func processInput(filename string, processor func(string) error) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		err = processor(scanner.Text())
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	processInput("input.txt", calcSum)
	fmt.Printf("Sum: %d\n", prioSum)
}
