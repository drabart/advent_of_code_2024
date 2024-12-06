package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Direction [2]int

var (
	North Direction = [2]int{-1, 0}
	South Direction = [2]int{1, 0}
	West  Direction = [2]int{0, -1}
	East  Direction = [2]int{0, 1}
)

type Piece int

const (
	Free Piece = iota
	Blocked
	Outside
	Seen
)

type player struct {
	row    int
	column int
	dir    Direction
}

func check_spot(board [][]Piece, row int, column int) Piece {
	if row < 0 || row >= len(board) || column < 0 || column >= len(board[0]) {
		return Outside
	}

	return board[row][column]
}

func part1(board [][]Piece, p player) {
	found := 0

	rotate_right := map[Direction]Direction{
		North: East,
		East:  South,
		South: West,
		West:  North,
	}

	visited := map[player]bool{}

	for true {
		if !visited[player{p.row, p.column, North}] {
			visited[player{p.row, p.column, North}] = true
			found += 1
		}

		new_spot := check_spot(board, p.row+p.dir[0], p.column+p.dir[1])
		if new_spot == Outside {
			break
		}
		if new_spot == Blocked {
			p.dir = rotate_right[p.dir]
		}
		if new_spot == Free {
			p.row += p.dir[0]
			p.column += p.dir[1]
		}
	}

	fmt.Printf("Part 1 solution: %d\n", found)
}

func cycles(board [][]Piece, p player, visited map[player]bool) bool {
	rotate_right := map[Direction]Direction{
		North: East,
		East:  South,
		South: West,
		West:  North,
	}

	for true {
		if !visited[p] {
			visited[p] = true
		} else {
			return true
		}

		new_spot := check_spot(board, p.row+p.dir[0], p.column+p.dir[1])
		if new_spot == Outside {
			break
		}
		if new_spot == Blocked {
			p.dir = rotate_right[p.dir]
		}
		if new_spot == Free {
			p.row += p.dir[0]
			p.column += p.dir[1]
		}
	}

	return false
}

func part2(board [][]Piece, p player) {
	found := 0

	rotate_right := map[Direction]Direction{
		North: East,
		East:  South,
		South: West,
		West:  North,
	}

	visited := map[player]bool{}
	valid_blocks := map[[2]int]bool{}

	for true {
		if !visited[p] {
			visited[p] = true
		}

		new_spot := check_spot(board, p.row+p.dir[0], p.column+p.dir[1])
		if new_spot == Outside {
			break
		}
		if new_spot == Blocked {
			p.dir = rotate_right[p.dir]
		}
		if new_spot == Free {
			if !valid_blocks[[2]int{}] &&
				!(visited[player{p.row + p.dir[0], p.column + p.dir[1], North}] ||
					visited[player{p.row + p.dir[0], p.column + p.dir[1], South}] ||
					visited[player{p.row + p.dir[0], p.column + p.dir[1], East}] ||
					visited[player{p.row + p.dir[0], p.column + p.dir[1], West}]) {

				board[p.row+p.dir[0]][p.column+p.dir[1]] = Blocked

				pc := p
				pc.dir = rotate_right[p.dir]
				visited_c := make(map[player]bool)
				for k, v := range visited {
					visited_c[k] = v
				}

				if cycles(board, pc, visited_c) {
					found += 1
				}

				board[p.row+p.dir[0]][p.column+p.dir[1]] = Free
				valid_blocks[[2]int{p.row + p.dir[0], p.column + p.dir[1]}] = true
			}

			p.row += p.dir[0]
			p.column += p.dir[1]
		}
	}

	fmt.Printf("Part 1 solution: %d\n", found)
}

func main() {
	content, err := os.Open(os.Args[1])
	defer content.Close()
	if err != nil {
		log.Fatal(err)
	}

	arrow_to_direction := map[rune]Direction{
		'^': North,
		'v': South,
		'<': West,
		'>': East,
	}
	board := [][]Piece{}
	var p player

	scanner := bufio.NewScanner(content)

	row := 0
	column := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		board = append(board, make([]Piece, len(line)))

		for _, c := range line {
			if c == '.' {
				board[row][column] = Free
			} else if c == '#' {
				board[row][column] = Blocked
			} else {
				board[row][column] = Free
				p = player{row: row, column: column, dir: arrow_to_direction[c]}
			}
			column += 1
		}

		row += 1
		column = 0
	}

	part1(board, p)
	part2(board, p)
}
