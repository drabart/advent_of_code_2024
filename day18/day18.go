package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func (p *pair) sub_d(p2 Direction) {
	p.x -= p2[0]
	p.y -= p2[1]
}

func (p *pair) mul(p2 int) {
	p.x *= p2
	p.y *= p2
}

var n int
var m int

func coordinates_valid(x int, y int) bool {
	return x >= 0 && y >= 0 && x < n && y < m
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

func part1(board [][]rune, p player, e pair) {
	visited := map[player]int{}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	visited[p] = 0
	heap.Push(&pq, &pq_e{value: p, priority: 0})
	dirs := []Direction{North, East, South, West}

	res := 2137

	for len(pq) > 0 {
		pla := heap.Pop(&pq).(*pq_e).value
		// fmt.Println(pla)

		if pla.p.x == e.x && pla.p.y == e.y {
			res = visited[pla]
			break
		}

		for _, d := range dirs {
			p2 := pla
			p2.p.add_d(d)

			if !coordinates_valid(p2.p.x, p2.p.y) || board[p2.p.x][p2.p.y] == 1 {
				continue
			}

			cost, vis := visited[p2]
			if !vis || cost > visited[pla]+1 {
				visited[p2] = visited[pla] + 1
				ne := &pq_e{value: p2,
					priority: visited[p2]}
				heap.Push(&pq, ne)
			}
		}
	}

	fmt.Printf("Part 1 solution: %d\n", res)
}

func check(board [][]rune, p player, e pair) bool {
	visited := map[player]int{}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	visited[p] = 0
	heap.Push(&pq, &pq_e{value: p, priority: 0})
	dirs := []Direction{North, East, South, West}

	res := 0

	for len(pq) > 0 {
		pla := heap.Pop(&pq).(*pq_e).value
		// fmt.Println(pla)

		if pla.p.x == e.x && pla.p.y == e.y {
			res = visited[pla]
			break
		}

		for _, d := range dirs {
			p2 := pla
			p2.p.add_d(d)

			if !coordinates_valid(p2.p.x, p2.p.y) || board[p2.p.x][p2.p.y] == 1 {
				continue
			}

			cost, vis := visited[p2]
			if !vis || cost > visited[pla]+1 {
				visited[p2] = visited[pla] + 1
				ne := &pq_e{value: p2,
					priority: visited[p2]}
				heap.Push(&pq, ne)
			}
		}
	}

	return res != 0
}

func part2(coords []pair, p player, e pair) {
	pp := 0
	kk := len(coords)
	res := -1
	for pp <= kk {
		mid := (pp + kk) / 2

		board := make([][]rune, n)
		for i := 0; i < n; i++ {
			board[i] = make([]rune, m)
		}

		for i := 0; i < mid; i++ {
			board[coords[i].x][coords[i].y] = 1
		}

		if check(board, p, e) {
			res = mid
			pp = mid + 1
		} else {
			kk = mid - 1
		}
	}

	fmt.Printf("Part 2 solution: %d,%d\n", coords[res].x, coords[res].y)
}

func main() {
	content, err := os.Open(os.Args[1])
	defer content.Close()
	if err != nil {
		log.Fatal(err)
	}

	p := player{pair{0, 0}}
	e := pair{}

	scanner := bufio.NewScanner(content)

	scanner.Scan()
	s := scanner.Text()
	ss := strings.Split(s, ",")
	n, _ = strconv.Atoi(ss[0])
	m, _ = strconv.Atoi(ss[1])
	n++
	m++
	e = pair{n - 1, m - 1}
	scanner.Scan()
	s = scanner.Text()
	v, _ := strconv.Atoi(s)

	coords := []pair{}

	board := make([][]rune, n)
	for i := 0; i < n; i++ {
		board[i] = make([]rune, m)
	}

	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			break
		}
		ss := strings.Split(s, ",")
		x, _ := strconv.Atoi(ss[0])
		y, _ := strconv.Atoi(ss[1])

		coords = append(coords, pair{x, y})

		v--
		if v > 0 {
			board[y][x] = 1
		}
	}

	part1(board, p, e)
	part2(coords, p, e)
}
