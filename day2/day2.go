package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func is_safe(list []int) (bool, int) {
	if len(list) <= 1 {
		return true, 1
	}
	increasing := list[0] < list[1]

	for i := 1; i < len(list); i++ {
		if increasing {
			if list[i-1] >= list[i] {
				return false, i
			}
		} else {
			if list[i-1] <= list[i] {
				return false, i
			}
		}

		diff := math.Abs(float64(list[i-1] - list[i]))
		if diff < 1 || 3 < diff {
			return false, i
		}
	}
	return true, 0
}

func is_safe_without_index(list []int, index int) bool {
	safe, _ := is_safe(append(append([]int{}, list[:index]...), list[index+1:]...))
	return safe
}

func part1(lists [][]int) {
	safe := 0

	for _, level := range lists {
		level_safe, _ := is_safe(level)
		if level_safe {
			safe += 1
		}
	}

	fmt.Printf("Part 1 solution: %d\n", safe)
}

func part2(lists [][]int) {
	safe := 0

	for _, level := range lists {
		level_safe, unsafe_index := is_safe(level)
		if level_safe {
			safe += 1
			continue
		}

		// Problem lies between unsafe_index - 1 and unsafe_index, so we try 2 configurations
		// we also check 0, as it might flip the direction
		if is_safe_without_index(level, 0) {
			safe += 1
			continue
		}

		if is_safe_without_index(level, unsafe_index-1) {
			safe += 1
			continue
		}

		if is_safe_without_index(level, unsafe_index) {
			safe += 1
			continue
		}
	}

	fmt.Printf("Part 2 solution: %d\n", safe)
}

func main() {
	content, err := os.Open(os.Args[1])
	defer content.Close()
	if err != nil {
		log.Fatal(err)
	}

	lists := [][]int{}

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		readings := strings.Fields(scanner.Text())

		list := []int{}
		for _, reading := range readings {
			num, err := strconv.Atoi(reading)
			if err != nil {
				continue
			}
			list = append(list, num)
		}

		lists = append(lists, list)
	}

	part1(lists)
	part2(lists)
}
