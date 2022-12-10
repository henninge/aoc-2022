package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Processor struct {
	regX  int
	cycle int

	screen [6][40]bool

	SignalSum int
}

func NewProcessor() *Processor {
	return &Processor{regX: 1}
}

func (p *Processor) drawPixel() {
	row := ((p.cycle - 1) / 40) % 6
	col := (p.cycle - 1) % 40
	if col >= p.regX-1 && col <= p.regX+1 {
		p.screen[row][col] = true
	}
}

func (p *Processor) tick() {
	p.cycle++
	if (p.cycle-20)%40 == 0 {
		p.SignalSum += p.regX * p.cycle
	}
	p.drawPixel()
}

func (p *Processor) AddX(val int) {
	p.tick()
	p.tick()
	p.regX += val
}

func (p *Processor) Noop() {
	p.tick()
}

func (p *Processor) Run(prog []string) {
	for _, line := range prog {
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "addx":
			val, _ := strconv.Atoi(parts[1])
			p.AddX(val)
		case "noop":
			p.Noop()
		}
	}
}

func (p *Processor) Display() {
	for _, row := range p.screen {
		for _, pixel := range row {
			if pixel {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
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

func main() {
	prog, err := getInput("input.txt")
	if err != nil {
		panic(err)
	}
	processor := NewProcessor()
	processor.Run(prog)
	fmt.Printf("SignalSum: %d\n", processor.SignalSum)
	processor.Display()
}
