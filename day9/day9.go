package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func part1(board []int) {
	b2 := []int{}
	file := true
	id := 0

	for _, a := range board {
		if file {
			for j := 0; j < a; j++ {
				b2 = append(b2, id)
			}

			id += 1
		} else {
			for j := 0; j < a; j++ {
				b2 = append(b2, -1)
			}
		}

		file = !file
	}

	it := 0
	for i := len(b2) - 1; i >= 0; i-- {
		if b2[i] != -1 {
			for b2[it] != -1 {
				it++
			}
			if it >= i {
				break
			}
			b2[it], b2[i] = b2[i], -1
		}
	}

	s := 0

	for i, b := range b2 {
		if b != -1 {
			s += i * b
		}
	}

	fmt.Printf("Part 1 solution: %d\n", s)
}

func part2(board []int) {
	b2 := [][3]int{}
	id := 0
	for i := range board {
		b2 = append(b2, [3]int{board[i], i % 2, id})
		if i%2 == 0 {
			id++
		}
	}

	for i := len(b2) - 1; i >= 0; i-- {
		if b2[i][1] == 1 {
			continue
		}
		for j := range b2 {
			if j >= i {
				break
			}
			if b2[j][1] == 0 {
				continue
			}
			if b2[j][0] >= b2[i][0] {
				b2[j][0] -= b2[i][0]
				b2 = slices.Insert(b2, j, b2[i])
				i++
				b2[i][1] = 1
				break
			}
		}
		i = min(i, len(b2)-1)
	}

	s := 0

	it := 0
	for _, b := range b2 {
		if b[1] == 1 {
			it += b[0]
			continue
		}
		s += (it + it + b[0] - 1) * b[0] / 2 * b[2]
		it += b[0]
	}

	fmt.Printf("Part 2 solution: %d\n", s)
}

func main() {
	content, err := os.Open(os.Args[1])
	defer content.Close()
	if err != nil {
		log.Fatal(err)
	}

	board := []int{}

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		for _, c := range line {
			a, _ := strconv.Atoi(string(c))
			board = append(board, a)
		}
	}

	part1(board)
	part2(board)
}
