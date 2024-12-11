package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func l10(a int) int {
	b := 0
	for a > 0 {
		b++
		a /= 10
	}
	return b
}

var pow10 = []int{
	1,
	10,
	100,
	1000,
	10000,
	100000,
	1000000,
	10000000,
	100000000,
}

var skip map[int]map[int]int = make(map[int]map[int]int)
var skip_step = 5

func compute_ss(a, step int) map[int]int {
	r2 := map[int]int{a: 1}
	rc := make(map[int]int)

	for i := 0; i < step; i++ {
		for k := range rc {
			delete(rc, k)
		}

		for k, v := range r2 {
			if k == 0 {
				rc[1] += v
				continue
			}
			lg := l10(k)
			if lg%2 == 0 {
				rc[k/pow10[lg/2]] += v
				rc[k%pow10[lg/2]] += v
				continue
			}
			rc[k*2024] += v
		}

		r2, rc = rc, r2
	}

	if step == skip_step {
		skip[a] = r2
	}
	return r2
}

func solve(r2 map[int]int, n int) int {
	rc := make(map[int]int)

	for n > 0 {
		for k := range rc {
			delete(rc, k)
		}

		for k, v := range r2 {
			skip_map, found := skip[k]

			if !found {
				skip_map = compute_ss(k)
			}

			for kk, vvv := range skip_map {
				rc[kk] += vvv * v
			}
		}

		r2, rc = rc, r2
	}

	sum := 0
	for _, v := range r2 {
		sum += v
	}

	return sum
}

func part1(rocks []int) {
	r2 := map[int]int{}

	for _, r := range rocks {
		r2[r] += 1
	}

	fmt.Printf("Part 1 solution: %d\n", solve(r2, 25))
}

func part2(rocks []int) {
	r2 := map[int]int{}

	for _, r := range rocks {
		r2[r] += 1
	}

	fmt.Printf("Part 2 solution: %d\n", solve(r2, 75))
}

func part3(rocks []int) {
	r2 := map[int]int{}

	for _, r := range rocks {
		r2[r] += 1
	}

	fmt.Printf("Part 3 solution: %d\n", solve(r2, 2000))
}

func main() {
	file, err := os.Open(os.Args[1])
	defer file.Close()
	if err != nil {
		fmt.Print("Invalid arguments")
	}

	s := bufio.NewScanner(file)

	s.Scan()
	l := s.Text()
	ls := strings.Split(l, " ")

	rocks := []int{}

	for _, n := range ls {
		nn, _ := strconv.Atoi(n)
		rocks = append(rocks, nn)
	}

	part1(rocks)
	part2(rocks)
	part3(rocks)
}
