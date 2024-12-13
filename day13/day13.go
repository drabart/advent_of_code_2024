package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

type machine struct {
	a, b, prize pair
}

func solve_machine(m machine) int {
	for at := 0; ; at++ {
		if at*m.a.x > m.prize.x {
			break
		}
		bt := m.prize.x - m.a.x*at
		if bt%m.b.x != 0 {
			continue
		}
		bt /= m.b.x
		if m.a.y*at+m.b.y*bt == m.prize.y {
			return at*3 + bt
		}
	}
	return 0
}

func part1(machines []machine) {
	s := 0

	for _, m := range machines {
		s += solve_machine(m)
	}

	fmt.Printf("Part 1 solution: %d\n", s)
}

func part2(machines []machine) {
	s := 0

	for _, m := range machines {
		m.prize.add(pair{10000000000000, 10000000000000})
		s += solve_machine(m)
	}

	fmt.Printf("Part 2 solution: %d\n", s)
}

func main() {
	content, err := os.Open(os.Args[1])
	defer content.Close()
	if err != nil {
		log.Fatal(err)
	}

	machines := []machine{}

	scanner := bufio.NewScanner(content)
	matcher := regexp.MustCompile("X\\+(.+), Y\\+(.+)|X\\=(.+), Y\\=(.+)")

	state := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		matches := matcher.FindStringSubmatch(line)
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		if x == 0 {
			x, _ = strconv.Atoi(matches[3])
			y, _ = strconv.Atoi(matches[4])
		}
		if state == 0 {
			m := machine{pair{x, y}, pair{0, 0}, pair{0, 0}}
			machines = append(machines, m)
			state = 1
		} else if state == 1 {
			machines[len(machines)-1].b = pair{x, y}
			state = 2
		} else {
			machines[len(machines)-1].prize = pair{x, y}
			state = 0
		}
	}

	part1(machines)
	part2(machines)
}
