package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

type Direction [2]int

var (
	North Direction = [2]int{-1, 0}
	South Direction = [2]int{1, 0}
	West  Direction = [2]int{0, -1}
	East  Direction = [2]int{0, 1}
)

type player struct {
	p pair
}

func move(board [][]rune, p *player, move Direction) {
	pos := p.p
	for board[pos.x][pos.y] == '@' || board[pos.x][pos.y] == 'O' {
		m := pair{move[0], move[1]}
		pos.add(m)
	}

	if board[pos.x][pos.y] == '#' {
		return
	}

	m := pair{-move[0], -move[1]}
	pos.add(m)

	for board[pos.x][pos.y] == '@' || board[pos.x][pos.y] == 'O' {
		board[pos.x+move[0]][pos.y+move[1]] = board[pos.x][pos.y]
		if pos == p.p {
			break
		}
		pos.add(m)
	}
	board[p.p.x][p.p.y] = '.'
	m = pair{move[0], move[1]}
	p.p.add(m)
}

func part1(board [][]rune, p player, moves []Direction) {
	for _, d := range moves {
		move(board, &p, d)
	}

	s := 0

	for i, r := range board {
		for j, c := range r {
			if c == 'O' {
				s += i*100 + j
			}
			fmt.Print(string(c))
		}
		fmt.Println()
	}

	fmt.Printf("Part 1 solution: %d\n", s)
}

func blocked(board [][]rune, pos pair, d Direction, side bool) bool {
	if d == East || d == West {
		side = false
	}
	if side {
		block := false
		switch board[pos.x][pos.y] {
		case '[':
			block = blocked(board, pair{pos.x, pos.y + 1}, d, false)
			break
		case ']':
			block = blocked(board, pair{pos.x, pos.y - 1}, d, false)
			break
		}
		if block {
			return true
		}
	}

	m := pair{d[0], d[1]}
	pos.add(m)
	switch board[pos.x][pos.y] {
	case '.':
		return false
	case '#':
		return true
	case ']':
		return blocked(board, pos, d, true)
	case '[':
		return blocked(board, pos, d, true)
	}
	fmt.Print("fail", string(board[pos.x][pos.y]))
	return false
}

func move_stuff(board [][]rune, pos pair, d Direction, side bool) {
	if d == East || d == West {
		side = false
	}
	if side {
		switch board[pos.x][pos.y] {
		case '[':
			move_stuff(board, pair{pos.x, pos.y + 1}, d, false)
			break
		case ']':
			move_stuff(board, pair{pos.x, pos.y - 1}, d, false)
			break
		}
	}

	pn := pos
	m := pair{d[0], d[1]}
	pn.add(m)
	switch board[pn.x][pn.y] {
	case ']':
		move_stuff(board, pn, d, true)
		break
	case '[':
		move_stuff(board, pn, d, true)
		break
	}
	board[pn.x][pn.y] = board[pos.x][pos.y]
	board[pos.x][pos.y] = '.'
}

func move_wide(board [][]rune, p *player, move Direction) {
	if blocked(board, p.p, move, true) {
		return
	}

	move_stuff(board, p.p, move, true)
	m := pair{move[0], move[1]}
	p.p.add(m)
}

func part2(board [][]rune, p player, moves []Direction) {
	for _, d := range moves {
		move_wide(board, &p, d)
	}

	s := 0

	for i, r := range board {
		for j, c := range r {
			if c == '[' {
				s += i*100 + j
			}
			fmt.Print(string(c))
		}
		fmt.Println()
	}

	fmt.Printf("Part 2 solution: %d\n", s)
}

func main() {
	content, err := os.Open(os.Args[1])
	defer content.Close()
	if err != nil {
		log.Fatal(err)
	}

	board := [][]rune{}
	board2 := [][]rune{}
	p := player{}
	p2 := player{}

	scanner := bufio.NewScanner(content)
	row := 0
	c := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		t := []rune{}
		t2 := []rune{}
		c = 0
		for _, r := range line {
			if r == '@' {
				p = player{pair{row, c}}
				p2 = player{pair{row, c * 2}}
			}
			t = append(t, r)
			switch r {
			case 'O':
				t2 = append(t2, '[')
				t2 = append(t2, ']')
				break
			case '.':
				t2 = append(t2, '.')
				t2 = append(t2, '.')
				break
			case '@':
				t2 = append(t2, '@')
				t2 = append(t2, '.')
				break
			case '#':
				t2 = append(t2, '#')
				t2 = append(t2, '#')
				break
			}
			c++
		}

		board = append(board, t)
		board2 = append(board2, t2)
		row++
	}

	d := []Direction{}
	md := map[rune]Direction{'^': North, '>': East, 'v': South, '<': West}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		for _, r := range line {
			d = append(d, md[r])
		}
	}

	part1(board, p, d)
	part2(board2, p2, d)
}
