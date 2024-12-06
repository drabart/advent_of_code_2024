package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func part1(s string) {
	sum := 0

	matcher := regexp.MustCompile("mul\\(([0-9]{1,3}),([0-9]{1,3})\\)")
	matches := matcher.FindAllStringSubmatch(s, -1)

	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		sum += a * b
	}

	fmt.Printf("Part 1 solution: %d\n", sum)
}

func part2(s string) {
	sum := 0

	matcher := regexp.MustCompile("mul\\(([0-9]{1,3}),([0-9]{1,3})\\)|do\\(\\)|don't\\(\\)")
	matches := matcher.FindAllStringSubmatch(s, -1)

	enabled := true

	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if enabled {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			sum += a * b
		}
	}

	fmt.Printf("Part 2 solution: %d\n", sum)
}

func main() {
	content, err := os.Open(os.Args[1])
	defer content.Close()
	if err != nil {
		log.Fatal(err)
	}

	s := ""
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		ss := scanner.Text()
		s += ss + "\n"
	}

	part1(s)
	part2(s)
}
