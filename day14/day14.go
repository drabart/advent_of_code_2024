package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var w, h int

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
	p, v pair
}

func (m *machine) add_v(x int) {
	v := m.v
	v.mul(x)
	m.p.add(v)

	m.p.x %= w
	m.p.x += w
	m.p.x %= w

	m.p.y %= h
	m.p.y += h
	m.p.y %= h
}

func solve_machine(m *machine) {
	v := m.v
	v.mul(100)
	m.p.add(v)

	m.p.x %= w
	m.p.x += w
	m.p.x %= w

	m.p.y %= h
	m.p.y += h
	m.p.y %= h
}

func part1(machines []machine) {
	s := [4]int{}

	for _, m := range machines {
		solve_machine(&m)
		if m.p.x < w/2 {
			if m.p.y < h/2 {
				s[0]++
			} else if m.p.y > h/2 {
				s[1]++
			}
		} else if m.p.x > w/2 {
			if m.p.y < h/2 {
				s[2]++
			} else if m.p.y > h/2 {
				s[3]++
			}
		}
	}

	fmt.Printf("Part 1 solution: %d\n", s[0]*s[1]*s[2]*s[3])
}

func part2(machines []machine) {
	s := 0

	for i := range machines {
		machines[i].add_v(12)
	}

	for x := 0; x <= 101; x++ {
		s++
		t := make([][]string, h)
		for i := 0; i < h; i++ {
			t[i] = make([]string, w)
			for j := 0; j < w; j++ {
				t[i][j] = " "
			}
		}
		for i := range machines {
			t[machines[i].p.y][machines[i].p.x] = "#"
		}

		fmt.Println(s)
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				fmt.Print(t[i][j])
			}
			fmt.Println()
		}

		fmt.Println()

		for i := range machines {
			machines[i].add_v(103)
		}
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
	matcher := regexp.MustCompile("p=(.+),(.+) v=(.+),(.+)")

	scanner.Scan()
	s := scanner.Text()
	s2 := strings.Split(s, " ")
	w, _ = strconv.Atoi(s2[0])
	h, _ = strconv.Atoi(s2[1])

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		matches := matcher.FindStringSubmatch(line)
		m := machine{}
		m.p.x, _ = strconv.Atoi(matches[1])
		m.p.y, _ = strconv.Atoi(matches[2])
		m.v.x, _ = strconv.Atoi(matches[3])
		m.v.y, _ = strconv.Atoi(matches[4])
		machines = append(machines, m)
	}

	part1(machines)
	part2(machines)
}
