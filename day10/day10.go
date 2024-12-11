package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var n int
var m int

func coordinates_valid(x int, y int) bool {
	return x >= 0 && y >= 0 && x < n && y < m
}

var h int
var helper [][][]int
var helper2 [][]int

func iterate_over(board [][]int, number int, f func([][]int, int, int, int, int)) {
	dirs := [4][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

	for row := range board {
		for column := range board[row] {
			if board[row][column] != number {
				continue
			}
			for _, dir := range dirs {
				f(board, row, column, row+dir[0], column+dir[1])
			}
		}
	}
}

func removeDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func part1(board [][]int) {
	n = len(board)
	m = len(board[0])

	helper = make([][][]int, n)
	for i := range helper {
		helper[i] = make([][]int, m)
	}

	f := func(board [][]int, row, column, nr, nc int) {
		if !coordinates_valid(nr, nc) {
			return
		}
		if board[row][column] == 0 {
			helper[row][column] = []int{row*n + column}
		}
		if board[nr][nc] == board[row][column]+1 {
			helper[nr][nc] = append(helper[nr][nc], helper[row][column]...)
			helper[nr][nc] = removeDuplicate(helper[nr][nc])
		}
	}

	for i := 0; i < 9; i++ {
		iterate_over(board, i, f)
	}

	g := func(board [][]int, row, column, _, _ int) {
		h += len(helper[row][column])
		helper[row][column] = []int{}
	}

	iterate_over(board, 9, g)

	fmt.Printf("Part 1 solution: %d\n", h)
}

func part2(board [][]int) {
	n = len(board)
	m = len(board[0])

	h = 0
	helper2 = make([][]int, n)
	for i := range helper {
		helper2[i] = make([]int, m)
	}

	f := func(board [][]int, row, column, nr, nc int) {
		if !coordinates_valid(nr, nc) {
			return
		}
		if board[row][column] == 0 {
			helper2[row][column] = 1
		}
		if board[nr][nc] == board[row][column]+1 {
			helper2[nr][nc] += helper2[row][column]
		}
	}

	for i := 0; i < 9; i++ {
		iterate_over(board, i, f)
	}

	g := func(board [][]int, row, column, _, _ int) {
		h += helper2[row][column]
		helper2[row][column] = 0
	}

	iterate_over(board, 9, g)

	fmt.Printf("Part 2 solution: %d\n", h)
}

func main() {
	content, err := os.Open(os.Args[1])
	defer content.Close()
	if err != nil {
		log.Fatal(err)
	}

	board := [][]int{}

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		t := []int{}

		for _, r := range line {
			a, _ := strconv.Atoi(string(r))
			t = append(t, a)
		}

		board = append(board, t)
	}

	part1(board)
	part2(board)
}
