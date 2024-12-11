package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("testInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
	}

}
