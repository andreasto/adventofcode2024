package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Levels []int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	instructions := []Instruction{}

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		parsedLevels := []int{}
		for _, s := range split {
			parsed, _ := strconv.Atoi(s)
			parsedLevels = append(parsedLevels, parsed)
		}

		instructions = append(instructions, Instruction{Levels: parsedLevels})
	}

	log.Printf("Day2, Part1: %v", answerPartOne(instructions))

	safeCount := 0
	for _, l := range instructions {
		safe := answerPartTwo(l)
		if safe {
			safeCount++
			continue
		}

		if !safe {
			// check if removing 1 value makes it safe
			levelsLength := len(l.Levels)
			for j := 0; j < levelsLength; j++ {
				copy := []int{}
				for k := 0; k < levelsLength; k++ {
					if k == j {
						continue
					}
					copy = append(copy, l.Levels[k])
				}
				if answerPartTwo(Instruction{Levels: copy}) {
					log.Printf("Safe: %v when index %v is removed", copy, j)
					safeCount++
					break
				}
			}
		}
	}
	log.Printf("Day2, Part2: %v", safeCount)
}

func answerPartTwo(instruction Instruction) bool {
	length := len(instruction.Levels)
	safe := false
	shouldIncrease := true
	for i, _ := range instruction.Levels {
		if i == length-1 {
			break
		}

		if i == 0 {
			if instruction.Levels[i] > instruction.Levels[i+1] {
				shouldIncrease = false
			} else if instruction.Levels[i] < instruction.Levels[i+1] {
				shouldIncrease = true
			} else {
				safe = false
				break
			}
		}

		currentLevel := instruction.Levels[i]
		nextLevel := instruction.Levels[i+1]

		safe = CheckSafe(currentLevel, nextLevel, shouldIncrease)

		if !safe {
			break
		}
	}

	return safe
}

func answerPartOne(instructions []Instruction) int {
	safeCount := 0
	for i := range instructions {
		shouldIncrease := true
		safe := true
		for j := range instructions[i].Levels {
			if j == 0 {
				if instructions[i].Levels[j] > instructions[i].Levels[j+1] {
					shouldIncrease = false
				} else if instructions[i].Levels[j] < instructions[i].Levels[j+1] {
					shouldIncrease = true
				} else {
					safe = false
					break
				}
			}

			if j == len(instructions[i].Levels)-1 {
				break
			}

			currentLevel := instructions[i].Levels[j]
			nextLevel := instructions[i].Levels[j+1]

			safe = CheckSafe(currentLevel, nextLevel, shouldIncrease)

			if !safe {
				break
			}
		}

		if safe {
			safeCount++
		}
	}

	return safeCount
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func CheckSafe(current, next int, increase bool) bool {
	if increase {
		if current > next || current == next || next-current > 3 {
			return false
		}
	} else {
		if current < next || current == next || current-next > 3 {
			return false
		}
	}

	return true
}
