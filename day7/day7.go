package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type operator_func func(int, int) int

func possible(sum int, csum int, numbers []int, operators []operator_func) bool {
	if len(numbers) == 0 {
		return sum == csum
	}
	if csum > sum {
		return false
	}
	p := false
	for _, operator := range operators {
		p = p || possible(sum, operator(csum, numbers[0]), numbers[1:], operators)
		if p {
			break
		}
	}
	return p
}

type TestSample struct {
	sum     int
	numbers []int
}

func plus(a, b int) int {
	return a + b
}

func mult(a, b int) int {
	return a * b
}

func concat(a, b int) int {
	return a*int(math.Pow10(int(math.Log10(float64(b))+1))) + b
}

func part1(samples []TestSample) {
	found := 0

	for _, sample := range samples {
		if possible(sample.sum, sample.numbers[0], sample.numbers[1:], []operator_func{plus, mult}) {
			found += sample.sum
		}
	}

	fmt.Printf("Part 1 solution: %d\n", found)
}

func part2(samples []TestSample) {
	found := 0

	for _, sample := range samples {
		if possible(sample.sum, sample.numbers[0], sample.numbers[1:], []operator_func{plus, mult, concat}) {
			found += sample.sum
		}
	}

	fmt.Printf("Part 2 solution: %d\n", found)
}

func main() {
	content, err := os.Open(os.Args[1])
	defer content.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(content)

	samples := []TestSample{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		split := strings.Split(line, ": ")
		sum, _ := strconv.Atoi(split[0])
		numbers_r := strings.Split(split[1], " ")
		numbers := []int{}
		for _, a := range numbers_r {
			b, _ := strconv.Atoi(a)
			numbers = append(numbers, b)
		}
		samples = append(samples, TestSample{sum, numbers})
	}

	part1(samples)
	part2(samples)
}
