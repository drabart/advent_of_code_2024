package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Antena rune

type position struct {
	x, y int
}

func (p position) Add(p2 position) position {
	p.x += p2.x
	p.y += p2.y
	return p
}

func (p position) Sub(p2 position) position {
	p.x -= p2.x
	p.y -= p2.y
	return p
}

func check_spot(board [][]Antena, p position) bool {
	if p.x < 0 || p.x >= len(board) || p.y < 0 || p.y >= len(board[0]) {
		return false
	}

	return true
}

func part1(board [][]Antena) {
	antennas := map[Antena][]position{}

	for i, row := range board {
		for j, field := range row {
			if field != '.' {
				antennas[field] = append(antennas[field], position{i, j})
			}
		}
	}

	hotspots := map[position]bool{}

	for _, antenna := range antennas {
		for i, a := range antenna {
			for j, b := range antenna {
				if i == j {
					continue
				}
				v := a.Sub(b)
				if check_spot(board, a.Add(v)) {
					hotspots[a.Add(v)] = true
				}
			}
		}
	}

	fmt.Printf("Part 1 solution: %d\n", len(hotspots))
}

func gcd(p position) position {
	pc := p
	for pc.y != 0 {
		pc.x, pc.y = pc.y, pc.x%pc.y
	}
	p.x /= pc.x
	p.y /= pc.x
	return p
}

func part2(board [][]Antena) {
	antennas := map[Antena][]position{}

	for i, row := range board {
		for j, field := range row {
			if field != '.' {
				antennas[field] = append(antennas[field], position{i, j})
			}
		}
	}

	hotspots := map[position]bool{}

	for _, antenna := range antennas {
		for i, a := range antenna {
			for j, b := range antenna {
				if i == j {
					continue
				}
				v := gcd(a.Sub(b))
				c := a.Add(v)
				for check_spot(board, c) {
					hotspots[c] = true
					c = c.Add(v)
				}
				c = a.Sub(v)
				for check_spot(board, c) {
					hotspots[c] = true
					c = c.Sub(v)
				}
			}
		}
	}

	fmt.Printf("Part 2 solution: %d\n", len(hotspots))
}

func main() {
	content, err := os.Open(os.Args[1])
	defer content.Close()
	if err != nil {
		log.Fatal(err)
	}

	board := [][]Antena{}

	scanner := bufio.NewScanner(content)

	row := 0
	column := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		board = append(board, make([]Antena, len(line)))

		for _, c := range line {
			board[row][column] = Antena(c)
			column += 1
		}

		row += 1
		column = 0
	}

	part1(board)
	part2(board)
}
