package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := readLines("./day2/day2-test-input.txt")
	check(err)
	fmt.Print(data)
}

func readLines(path string) ([]string, error) {
	dat, err := os.Open(path)
	check(err)
	defer dat.Close()

	var lines []string

	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		fmt.Printf("line: %s\n", scanner.Text())
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines, err
}
