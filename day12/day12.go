package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var n int
var m int

func coordinates_valid(x int, y int) bool {
	return x >= 0 && y >= 0 && x < n && y < m
}

var area int
var peri int
var it int

func dfs(x, y int, board [][]rune, vis [][]int) {
	area++
	vis[x][y] = it

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, dir := range dirs {
		nx, ny := x+dir[0], y+dir[1]
		if !coordinates_valid(nx, ny) || board[nx][ny] != board[x][y] {
			peri++
		}
		if !coordinates_valid(nx, ny) {
			continue
		}
		if vis[nx][ny] != 0 {
			continue
		}
		if board[nx][ny] != board[x][y] {
			continue
		}
		dfs(nx, ny, board, vis)
	}
}

func part1(board [][]rune) {
	vis := make([][]int, len(board))
	for i := range vis {
		vis[i] = make([]int, len(board[0]))
	}

	n = len(board)
	m = len(board[0])
	s := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if vis[i][j] != 0 {
				continue
			}
			it++
			area, peri = 0, 0
			dfs(i, j, board, vis)
			s += area * peri
		}
	}

	fmt.Printf("Part 1 solution: %d\n", s)
}

func count_sides(board [][]int) int {
	c := 0

	pe := false
	for j := 0; j < m; j++ {
		if !pe && board[0][j] == it {
			c++
			pe = true
		}
		if pe && board[0][j] != it {
			pe = false
		}
	}

	for i := 1; i < n; i++ {
		pe = false
		for j := 0; j < m; j++ {
			if !pe && (board[i][j] == it && board[i-1][j] != board[i][j]) {
				c++
				pe = true
			}
			if pe && (board[i][j] != it || board[i-1][j] == board[i][j]) {
				pe = false
			}
		}
	}
	for i := 0; i < n-1; i++ {
		pe = false
		for j := 0; j < m; j++ {
			if !pe && (board[i][j] == it && board[i+1][j] != board[i][j]) {
				c++
				pe = true
			}
			if pe && (board[i][j] != it || board[i+1][j] == board[i][j]) {
				pe = false
			}
		}
	}

	pe = false
	for j := 0; j < m; j++ {
		if !pe && board[n-1][j] == it {
			c++
			pe = true
		}
		if pe && board[n-1][j] != it {
			pe = false
		}
	}

	pe = false
	for j := 0; j < n; j++ {
		if !pe && board[j][0] == it {
			c++
			pe = true
		}
		if pe && board[j][0] != it {
			pe = false
		}
	}

	for i := 1; i < m; i++ {
		pe = false
		for j := 0; j < n; j++ {
			if !pe && (board[j][i] == it && board[j][i-1] != board[j][i]) {
				c++
				pe = true
			}
			if pe && (board[j][i] != it || board[j][i-1] == board[j][i]) {
				pe = false
			}
		}
	}
	for i := 0; i < m-1; i++ {
		pe = false
		for j := 0; j < n; j++ {
			if !pe && (board[j][i] == it && board[j][i+1] != board[j][i]) {
				c++
				pe = true
			}
			if pe && (board[j][i] != it || board[j][i+1] == board[j][i]) {
				pe = false
			}
		}
	}

	pe = false
	for j := 0; j < m; j++ {
		if !pe && board[j][m-1] == it {
			c++
			pe = true
		}
		if pe && board[j][m-1] != it {
			pe = false
		}
	}

	return c
}

func part2(board [][]rune) {
	vis := make([][]int, len(board))
	for i := range vis {
		vis[i] = make([]int, len(board[0]))
	}

	n = len(board)
	m = len(board[0])
	s := 0
	it = 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if vis[i][j] != 0 {
				continue
			}
			it++
			area = 0
			dfs(i, j, board, vis)

			sides := 0

			sides += count_sides(vis)

			s += area * sides
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

	board := [][]rune{}

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		t := []rune{}

		for _, r := range line {
			t = append(t, r)
		}

		board = append(board, t)
	}

	part1(board)
	part2(board)
}
