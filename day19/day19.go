package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type pair struct {
	x, y int
}

func (p *pair) add(p2 pair) {
	p.x += p2.x
	p.y += p2.y
}

func (p *pair) mul(p2 int) {
	p.x *= p2
	p.y *= p2
}

func check_creation(towels []string, creation string) int {
	possible := make([]int, len(creation)+1)
	possible[0] = 1

	for i := 0; i < len(creation); i++ {
		if possible[i] == 0 {
			continue
		}
		for _, towel := range towels {
			if strings.HasPrefix(creation[i:], towel) {
				possible[i+len(towel)] += possible[i]
			}
		}
	}

	return possible[len(creation)]
}

func part1(towels []string, creations []string) {
	res := 0
	for _, creation := range creations {
		if check_creation(towels, creation) > 0 {
			res++
		}
	}

	fmt.Printf("Part 1 solution: %d\n", res)
}

func part2(towels []string, creations []string) {
	res := 0
	for _, creation := range creations {
		res += check_creation(towels, creation)
	}

	fmt.Printf("Part 2 solution: %d\n", res)
}

func main() {
	content, err := os.Open(os.Args[1])
	defer content.Close()
	if err != nil {
		log.Fatal(err)
	}

	towels := []string{}
	creations := []string{}

	scanner := bufio.NewScanner(content)

	scanner.Scan()
	s := scanner.Text()
	ss := strings.Split(s, ", ")
	for _, towel := range ss {
		towels = append(towels, towel)
	}

	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}
		creations = append(creations, s)
	}

	part1(towels, creations)
	part2(towels, creations)
}
