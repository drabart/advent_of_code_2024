package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func part1(rocks []int) {
	for i := 0; i < 25; i++ {
		rc := []int{}
		for _, v := range rocks {
			if v == 0 {
				rc = append(rc, 1)
				continue
			}
			lg := (int(math.Log10(float64(v))) + 1)
			if lg%2 == 0 {
				rc = append(rc, v/int(math.Pow10(lg/2)))
				rc = append(rc, v%int(math.Pow10(lg/2)))
				continue
			}
			rc = append(rc, v*2024)
		}
		rocks = rc
	}
	fmt.Printf("Part 1 solution: %d\n", len(rocks))
}

func part2(rocks []int) {
	r2 := map[int]int{}

	for _, r := range rocks {
		r2[r] += 1
	}

	for i := 0; i < 75; i++ {
		rc := map[int]int{}
		for k, v := range r2 {
			if k == 0 {
				rc[1] += v
				continue
			}
			lg := (int(math.Log10(float64(k))) + 1)
			if lg%2 == 0 {
				rc[k/int(math.Pow10(lg/2))] += v
				rc[k%int(math.Pow10(lg/2))] += v
				continue
			}
			rc[k*2024] += v
		}
		r2 = rc
		fmt.Println(len(r2), i)
	}

	sum := 0
	for _, v := range r2 {
		sum += v
	}

	fmt.Printf("Part 2 solution: %d\n", sum)
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

	r2 := make([]int, len(rocks))
	copy(r2, rocks)

	part1(r2)
	part2(rocks)
}
