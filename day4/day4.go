package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func coordinates_valid(x int, y int, n int, m int) bool {
	return x >= 0 && y >= 0 && x < n && y < m
}

func check_xmas(lines []string, row int, column int) int {
	match := "XMAS"
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	number_found := 0
	for _, direction := range directions {
		found := true
		for i := 0; i < len(match); i++ {
			if !coordinates_valid(row+direction[0]*i, column+direction[1]*i, len(lines), len(lines[0])) {
				found = false
				break
			}
			if lines[row+direction[0]*i][column+direction[1]*i] != match[i] {
				found = false
				break
			}
		}
		if found {
			number_found++
		}
	}
	return number_found
}

func check_mas(lines []string, row int, column int) int {
	match := []string{"M.M", ".A.", "S.S"}
	directions := [][]int{{1, 1, 1}, {-1, 1, 1}, {-1, -1, 1}, {1, -1, 1}, {1, 1, 0}, {-1, 1, 0}, {-1, -1, 0}, {1, -1, 0}}
	number_found := 0
	for _, direction := range directions {
		found := true
		for i := 0; i < len(match); i++ {
			for j := 0; j < len(match[0]); j++ {
				if !coordinates_valid(row+direction[0]*i, column+direction[1]*j, len(lines), len(lines[0])) {
					found = false
					break
				}
				if match[i*direction[2]+j*(1-direction[2])][i*(1-direction[2])+j*direction[2]] != '.' &&
					lines[row+direction[0]*i][column+direction[1]*j] != match[i*direction[2]+j*(1-direction[2])][i*(1-direction[2])+j*direction[2]] {
					found = false
					break
				}
			}
		}
		if found {
			number_found++
		}
	}
	return number_found
}

func part1(lines []string) {
	found := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			found += check_xmas(lines, i, j)
		}
	}

	fmt.Printf("Part 1 solution: %d\n", found)
}

func part2(lines []string) {
	found := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			found += check_mas(lines, i, j)
		}
	}

	fmt.Printf("Part 2 solution: %d\n", found/2)
}

func main() {
	content, err := os.Open(os.Args[1])
	defer content.Close()
	if err != nil {
		log.Fatal(err)
	}

	lines := []string{}

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	part1(lines)
	part2(lines)
}
