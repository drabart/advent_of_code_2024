package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func part1(list1 []int32, list2 []int32) {
	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	var sum int32 = 0
	for i := 0; i < len(list1); i++ {
		sum += int32(math.Abs(float64(list1[i] - list2[i])))
	}

	fmt.Printf("Part 1 solution: %d\n", sum)
}

func part2(list1 []int32, list2 []int32) {
	var sum int32 = 0

	count := make(map[int32]int32)

	for i := 0; i < len(list2); i++ {
		_, ok := count[list2[i]]
		if !ok {
			count[list2[i]] = 0
		}
		count[list2[i]] += 1
	}

	for i := 0; i < len(list1); i++ {
		value := count[list1[i]]
		sum += list1[i] * value
	}

	fmt.Printf("Part 2 solution: %d\n", sum)
}

func main() {
	content, err := os.Open(os.Args[1])
	defer content.Close()
	if err != nil {
		log.Fatal(err)
	}

	list1 := []int32{}
	list2 := []int32{}

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		var a, b int32
		fmt.Sscanf(scanner.Text(), "%d %d", &a, &b)

		list1 = append(list1, a)
		list2 = append(list2, b)
	}

	part1(list1, list2)
	part2(list1, list2)
}
