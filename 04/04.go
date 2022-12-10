package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parseSets(line string) (s1, s2 *Set) {
	pat := regexp.MustCompile("(.+)-(.+),(.+)-(.+)")
	match := pat.FindStringSubmatch(line)
	s1 = &Set{}
	s2 = &Set{}
	val1, _ := strconv.Atoi(match[1])
	val2, _ := strconv.Atoi(match[2])
	s1.AddRange(val1, val2)
	val3, _ := strconv.Atoi(match[3])
	val4, _ := strconv.Atoi(match[4])
	s2.AddRange(val3, val4)

	return
}

var containsCounter, intersectsCounter int

func processLine(line string) error {
	s1, s2 := parseSets(line)
	if s1.Contains(s2) || s2.Contains(s1) {
		containsCounter++
	}
	if s1.Intersects(s2) {
		intersectsCounter++
	}

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
	processInput("input.txt", processLine)
	fmt.Printf("Contains: %d\n", containsCounter)
	fmt.Printf("Intersects: %d\n", intersectsCounter)
}
