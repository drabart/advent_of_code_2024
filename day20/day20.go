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

func (p *pair) sub_d(p2 Direction) {
	p.x -= p2[0]
	p.y -= p2[1]
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

func valid_coords(n, m, r, c int) bool {
	return r >= 0 && c >= 0 && r < n && c < m
}

func part1(board [][]rune, p player) {
	visited := map[player]int{}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	visited[p] = 0
	heap.Push(&pq, &pq_e{value: p, priority: 0})

	for len(pq) > 0 {
		pla := heap.Pop(&pq).(*pq_e).value

		for _, d := range dirs {
			p2 := pla
			p2.p.add_d(d)

			if !valid_coords(len(board), len(board[0]), p2.p.x, p2.p.y) || board[p2.p.x][p2.p.y] != '.' {
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

	for i := range board {
		for j := range board[0] {
			if board[i][j] == '#' {
				visited[player{pair{i, j}}] = 21372137
			}
			// fmt.Printf("%8d ", visited[player{pair{i, j}}])
		}
		// fmt.Println()
	}

	res := map[int]int{}

	for i := range board {
		for j := range board[0] {
			for _, d1 := range dirs {
				for _, d2 := range dirs {
					p0 := pair{i, j}
					p1 := p0
					p1.add_d(d1)
					p2 := p1
					p2.add_d(d2)
					if !valid_coords(len(board), len(board[0]), p1.x, p1.y) ||
						!valid_coords(len(board), len(board[0]), p2.x, p2.y) {
						continue
					}
					if board[p0.x][p0.y] != '.' || board[p2.x][p2.y] != '.' {
						continue
					}
					res[visited[player{p2}]-visited[player{p0}]-2]++
				}
			}
		}
	}

	sum := 0
	for k, v := range res {
		// if k > 0 {
		// 	fmt.Println(k, v)
		// }
		if k >= 100 {
			sum += v
		}
	}

	fmt.Printf("Part 1 solution: %d\n", sum)
}

var vv map[[2]pair]bool = map[[2]pair]bool{}
var dirs = []Direction{North, East, South, West}
var dfs_vis map[pair]bool
var nn = map[int]int{}

func bfs(board [][]rune, dist map[player]int, p pair) {
	dfs_vis = map[pair]bool{}
	b := []pair{}
	a := []pair{p}

	for i := 0; i < 20; i++ {
		for len(a) > 0 {
			pla := a[0]
			a = a[1:]

			for _, d := range dirs {
				p2 := pla
				p2.add_d(d)

				if dfs_vis[p2] {
					continue
				}
				if !valid_coords(len(board), len(board[0]), p2.x, p2.y) {
					continue
				}
				if board[p2.x][p2.y] != '#' {
					nn[dist[player{p2}]-dist[player{p}]-i]++
					if dist[player{p2}]-dist[player{p}]-i >= 100 {
						vv[[2]pair{p2, p}] = true
					}
				}

				dfs_vis[p2] = true
				b = append(b, p2)
			}
		}
		a = b
		b = []pair{}
	}

}

func part2(board [][]rune, p player) {
	visited := map[player]int{}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	visited[p] = 0
	heap.Push(&pq, &pq_e{value: p, priority: 0})

	for len(pq) > 0 {
		pla := heap.Pop(&pq).(*pq_e).value

		for _, d := range dirs {
			p2 := pla
			p2.p.add_d(d)

			if !valid_coords(len(board), len(board[0]), p2.p.x, p2.p.y) || board[p2.p.x][p2.p.y] != '.' {
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

	for i := range board {
		for j := range board[0] {
			if board[i][j] == '#' {
				visited[player{pair{i, j}}] = 21372137
			}
			// fmt.Printf("%8d ", visited[player{pair{i, j}}])
		}
		// fmt.Println()
	}

	for i := range board {
		for j := range board[0] {
			if board[i][j] == '#' {
				continue
			}
			bfs(board, visited, pair{i, j})
		}
	}

	s := 0
	for k, v := range nn {
		if k >= 100 {
			// fmt.Println(k, v)
			s += v
		}
	}
	// fmt.Println(nn)

	fmt.Printf("Part 2 solution: %d\n", len(vv))
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
			if r == 'E' {
				p = player{pair{row, c}}
				t = append(t, '.')
				c++
				continue
			}
			if r == 'S' {
				t = append(t, '.')
				c++
				continue
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
