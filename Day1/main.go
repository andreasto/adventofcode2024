package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type instruction struct {
	Left  []int
	Right []int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var inputs instruction

	for scanner.Scan() {
		fixed := strings.ReplaceAll(scanner.Text(), "   ", ",")
		left, right := strings.Split(fixed, ",")[0], strings.Split(fixed, ",")[1]
		parsedLeft, _ := strconv.Atoi(left)
		parsedRight, _ := strconv.Atoi(right)

		inputs.Left = append(inputs.Left, parsedLeft)
		inputs.Right = append(inputs.Right, parsedRight)
	}

	sort.Slice(inputs.Left, func(i, j int) bool {
		return inputs.Left[i] < inputs.Left[j]
	})

	sort.Slice(inputs.Right, func(i, j int) bool {
		return inputs.Right[i] < inputs.Right[j]
	})

	log.Printf("Day1 Task 1: %v", answerPartOne(inputs))
	log.Printf("Day1 Task 2: %v", answerPartTwo(inputs))
}

func answerPartOne(inputs instruction) int {
	result := []int{}
	for i := 0; i < len(inputs.Left); i++ {
		var diff int
		if inputs.Left[i] > inputs.Right[i] {
			diff = inputs.Left[i] - inputs.Right[i]
		} else if inputs.Left[i] < inputs.Right[i] {
			diff = inputs.Right[i] - inputs.Left[i]
		} else {
			diff = 0
		}
		result = append(result, diff)
	}

	final := 0
	for i := 0; i < len(result); i++ {
		final += result[i]
	}

	return final
}

func answerPartTwo(inputs instruction) int {
	result := []int{}
	for i := 0; i < len(inputs.Left); i++ {
		numberOfHits := 0
		for j := 0; j < len(inputs.Right); j++ {
			if inputs.Left[i] == inputs.Right[j] {
				numberOfHits++
			}
		}
		result = append(result, inputs.Left[i]*numberOfHits)
	}

	final := 0
	for i := 0; i < len(result); i++ {
		final += result[i]
	}

	return final
}
