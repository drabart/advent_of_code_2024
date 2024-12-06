package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func correct(rules [][]int, problem []int) bool {
	ok := true

	rest := map[int][]int{}

	for _, rule := range rules {
		rest[rule[1]] = append(rest[rule[1]], rule[0])
	}

	seen := map[int]bool{}
	restricted := map[int]bool{}

	for _, v := range problem {
		if restricted[v] {
			ok = false
			break
		}

		if !seen[v] {
			seen[v] = true
			for _, r := range rest[v] {
				restricted[r] = true
			}
		}
	}

	return ok
}

func fix(rules [][]int, problem []int) []int {
	rest := map[int][]int{}

	for _, rule := range rules {
		rest[rule[1]] = append(rest[rule[1]], rule[0])
	}

	// toposort might be more efficient, but removing inversions is fast enough
	ok := false
	for !ok {
		needs_inversion := []int{}
		seen := map[int]bool{}
		restricted := map[int]int{}
		ok = true

		for i, v := range problem {
			x, rs := restricted[v]
			if rs {
				ok = false
				needs_inversion = append(needs_inversion, i, x)
				break
			}

			if !seen[v] {
				seen[v] = true
				for _, r := range rest[v] {
					_, rs = restricted[r]
					if !rs {
						restricted[r] = i
					}
				}
			}
		}

		if !ok {
			problem[needs_inversion[0]], problem[needs_inversion[1]] =
				problem[needs_inversion[1]], problem[needs_inversion[0]]
		}
	}

	return problem
}

func part1(rules [][]int, problems [][]int) {
	found := 0

	for _, problem := range problems {
		if correct(rules, problem) {
			found += problem[len(problem)/2]
		}
	}

	fmt.Printf("Part 1 solution: %d\n", found)
}

func part2(rules [][]int, problems [][]int) {
	found := 0

	for _, problem := range problems {
		if !correct(rules, problem) {
			found += fix(rules, problem)[len(problem)/2]
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

	rules := [][]int{}
	problems := [][]int{}

	scanner := bufio.NewScanner(content)

	ip1 := true

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			ip1 = false
			continue
		}
		if ip1 {
			rules = append(rules, []int{0, 0})
			fmt.Sscanf(line, "%d|%d", &rules[len(rules)-1][0], &rules[len(rules)-1][1])
		} else {
			parts := strings.Split(line, ",")

			l := []int{}
			for _, part := range parts {
				val, _ := strconv.Atoi(part)
				l = append(l, val)
			}
			problems = append(problems, l)
		}
	}

	part1(rules, problems)
	part2(rules, problems)
}
