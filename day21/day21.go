package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var keypad = []string{
	"789",
	"456",
	"123",
	"X0A",
}

var arrows = []string{
	"X^A",
	"<v>",
}

var paths_keypad map[[2]byte][3]int
var paths_arrows map[[2]byte][3]int

func calculate_paths_kb(board []string) map[[2]byte][3]int {
	res := map[[2]byte][3]int{}
	for i := range board {
		for j := range board[0] {

			for k := range board {
				for l := range board[0] {
					if j == 0 && k == 3 {
						res[[2]byte{board[i][j], board[k][l]}] = [3]int{k - i, l - j, 2}
						continue
					}
					if l == 0 && i == 3 {
						res[[2]byte{board[i][j], board[k][l]}] = [3]int{k - i, l - j, 1}
						continue
					}
					res[[2]byte{board[i][j], board[k][l]}] = [3]int{k - i, l - j, 3}
				}
			}
		}
	}
	return res
}

func calculate_paths_arr(board []string) map[[2]byte][3]int {
	res := map[[2]byte][3]int{}
	for i := range board {
		for j := range board[0] {

			for k := range board {
				for l := range board[0] {
					if j == 0 && k == 0 {
						res[[2]byte{board[i][j], board[k][l]}] = [3]int{k - i, l - j, 2}
						continue
					}
					if l == 0 && i == 0 {
						res[[2]byte{board[i][j], board[k][l]}] = [3]int{k - i, l - j, 1}
						continue
					}
					res[[2]byte{board[i][j], board[k][l]}] = [3]int{k - i, l - j, 3}
				}
			}
		}
	}
	return res
}

type mem struct {
	s     string
	level int
}

var memory = map[mem]int{}

func expand(s string, curr int, levels int) int {
	v, seen := memory[mem{s, curr}]
	if seen {
		return v
	}

	res := 0
	prev := 'A'

	if curr == levels {
		memory[mem{s, curr}] = len(s)
		return len(s)
	}

	for _, c := range s {
		a := [3]int{}
		if curr == 0 {
			a = paths_keypad[[2]byte{byte(prev), byte(c)}]
		} else {
			a = paths_arrows[[2]byte{byte(prev), byte(c)}]
		}

		best := 0
		if (a[2] & 1) > 0 {
			b := ""
			for i := 0; i < a[0]; i++ {
				b += "v"
			}
			for i := 0; i > a[0]; i-- {
				b += "^"
			}
			for i := 0; i < a[1]; i++ {
				b += ">"
			}
			for i := 0; i > a[1]; i-- {
				b += "<"
			}
			b += "A"
			c := expand(b, curr+1, levels)
			if best == 0 || best > c {
				best = c
			}
		}
		if (a[2] & 2) > 0 {
			b := ""
			for i := 0; i < a[1]; i++ {
				b += ">"
			}
			for i := 0; i > a[1]; i-- {
				b += "<"
			}
			for i := 0; i < a[0]; i++ {
				b += "v"
			}
			for i := 0; i > a[0]; i-- {
				b += "^"
			}
			b += "A"
			c := expand(b, curr+1, levels)
			if best == 0 || best > c {
				best = c
			}
		}

		res += best
		prev = c
	}

	memory[mem{s, curr}] = res
	return res
}

func part1(codes []string) {
	memory = map[mem]int{}
	res := 0

	paths_keypad = calculate_paths_kb(keypad)
	paths_arrows = calculate_paths_arr(arrows)

	for _, code := range codes {
		new_code := expand(code, 0, 3)
		num, _ := strconv.Atoi(code[:len(code)-1])
		fmt.Println(new_code, num)
		res += new_code * num
	}

	fmt.Printf("Part 1 solution: %d\n", res)
}

func part2(codes []string) {
	memory = map[mem]int{}
	res := 0

	paths_keypad = calculate_paths_kb(keypad)
	paths_arrows = calculate_paths_arr(arrows)

	for _, code := range codes {
		new_code := expand(code, 0, 26)
		num, _ := strconv.Atoi(code[:len(code)-1])
		fmt.Println(new_code, num)
		res += new_code * num
	}

	fmt.Printf("Part 2 solution: %d\n", res)
}

func main() {
	content, err := os.Open(os.Args[1])
	defer content.Close()
	if err != nil {
		log.Fatal(err)
	}

	codes := []string{}

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		codes = append(codes, line)
	}

	part1(codes)
	part2(codes)
}
