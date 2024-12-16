package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
)

type pq_e struct {
	value    player
	priority int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*pq_e

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*pq_e)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *pq_e, value player, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

type pair struct {
	x, y int
}

func (p *pair) add(p2 pair) {
	p.x += p2.x
	p.y += p2.y
}

func (p *pair) add_d(p2 Direction) {
	p.x += p2[0]
	p.y += p2[1]
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
	d Direction
}

func part1(board [][]rune, p player) {
	visited := map[player]int{}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	visited[p] = 0
	pq.Push(&pq_e{value: p, priority: 0})

	dirs := []Direction{North, East, South, West}

	res := 0

	for len(pq) > 0 {
		pla := heap.Pop(&pq).(*pq_e).value

		if board[pla.p.x][pla.p.y] == 'E' {
			res = visited[pla]
			break
		}

		p2 := pla
		p2.p.add_d(p2.d)

		if board[p2.p.x][p2.p.y] != '#' {
			cost, vis := visited[p2]

			if !vis || cost > visited[pla]+1 {
				visited[p2] = visited[pla] + 1
				ne := &pq_e{value: p2,
					priority: visited[p2]}
				heap.Push(&pq, ne)
			}
		}

		for _, d := range dirs {
			p2 = pla
			p2.d = d
			cost, vis := visited[p2]
			if !vis || cost > visited[pla]+1000 {
				visited[p2] = visited[pla] + 1000
				ne := &pq_e{value: p2,
					priority: visited[p2]}
				heap.Push(&pq, ne)
			}
		}
	}

	// for i := range board {
	// 	for j := range board[0] {
	// 		p1 := player{pair{i, j}, North}
	// 		p2 := player{pair{i, j}, South}
	// 		p3 := player{pair{i, j}, East}
	// 		p4 := player{pair{i, j}, West}
	// 		fmt.Print(board[i][j], min(visited[p1], visited[p2], visited[p3], visited[p4]), " ")
	// 	}
	// 	fmt.Println()
	// }

	fmt.Printf("Part 1 solution: %d\n", res)
}

func part2(board [][]rune, p player) {
	visited := map[player]int{}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	visited[p] = 0
	pq.Push(&pq_e{value: p, priority: 0})

	dirs := []Direction{North, East, South, West}

	res := 0
	e := pair{}

	for len(pq) > 0 {
		pla := heap.Pop(&pq).(*pq_e).value

		if board[pla.p.x][pla.p.y] == 'E' {
			if res == 0 {
				res = visited[pla]
			}
			e = pla.p
		}

		p2 := pla
		p2.p.add_d(p2.d)

		if board[p2.p.x][p2.p.y] != '#' {
			cost, vis := visited[p2]

			if !vis || cost > visited[pla]+1 {
				visited[p2] = visited[pla] + 1
				ne := &pq_e{value: p2,
					priority: visited[p2]}
				heap.Push(&pq, ne)
			}
		}

		for _, d := range dirs {
			p2 = pla
			p2.d = d
			cost, vis := visited[p2]
			if !vis || cost > visited[pla]+1000 {
				visited[p2] = visited[pla] + 1000
				ne := &pq_e{value: p2,
					priority: visited[p2]}
				heap.Push(&pq, ne)
			}
		}
	}

	fmt.Printf("Part 1 solution: %d\n", res)
}

func main() {
	content, err := os.Open(os.Args[1])
	defer content.Close()
	if err != nil {
		log.Fatal(err)
	}

	board := [][]rune{}
	p := player{}

	scanner := bufio.NewScanner(content)
	row := 0
	c := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		t := []rune{}
		c = 0
		for _, r := range line {
			if r == 'S' {
				p = player{pair{row, c}, East}
			}
			t = append(t, r)
			c++
		}

		board = append(board, t)
		row++
	}

	part1(board, p)
	part2(board, p)
}
