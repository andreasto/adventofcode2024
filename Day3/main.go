package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)

	matches := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		m := re.FindAllString(line, -1)
		matches = append(matches, m...)
	}

	file, err = os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	log.Printf("Day3 Part1: %d\n", Part1(matches))

	re = regexp.MustCompile(`do\(\)|don't\(\)|mul\([0-9]{1,3},[0-9]{1,3}\)`)

	isEnabled := true
	total := 0

	inputString := ""
	for scanner.Scan() {
		inputString += scanner.Text()
	}

	matches = re.FindAllString(inputString, -1)

	for _, match := range matches {
		switch match {
		case "do()":
			isEnabled = true
		case "don't()":
			isEnabled = false
		default:
			if strings.HasPrefix(match, "mul") && isEnabled {
				var numbers = GetNumbers(match)
				total += numbers[0] * numbers[1]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Day3 Part2: %d\n", total)

}

func GetNumbers(match string) []int {
	re := regexp.MustCompile(`\d{1,3}`)
	numbers := re.FindAllString(match, -1)

	a, _ := strconv.Atoi(numbers[0])
	b, _ := strconv.Atoi(numbers[1])

	returnValue := []int{}
	returnValue = append(returnValue, a)
	returnValue = append(returnValue, b)

	return returnValue
}

func Part1(matches []string) int {
	totalCount := 0
	for _, match := range matches {
		match = strings.Replace(match, "mul(", "", -1)
		match = strings.Replace(match, ")", "", -1)

		nums := strings.Split(match, ",")

		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		totalCount += num1 * num2
	}

	return totalCount
}

func Part2(rows []string) int {

	return 0
}
