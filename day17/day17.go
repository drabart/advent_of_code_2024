package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type pair struct {
	x, y int
}

func (p *pair) add(p2 pair) {
	p.x += p2.x
	p.y += p2.y
}

func (p *pair) mul(p2 int) {
	p.x *= p2
	p.y *= p2
}

type command_t int

const (
	adv command_t = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

var match_func = map[command_t]func(regs []int, value command_t){
	adv: a_div,
	bxl: b_xor,
	bst: b_st,
	jnz: j_nz,
	bxc: b_xorc,
	out: out_p,
	bdv: b_div,
	cdv: c_div,
}

type regs_t int

const (
	A regs_t = iota
	B
	C
)

func get_combo(regs []int, combo command_t) int {
	switch combo {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		return int(combo)
	case 4:
		return regs[A]
	case 5:
		return regs[B]
	case 6:
		return regs[C]
	default:
		fmt.Print("reserved!")
		return -1
	}
}

var pcnt int = 0

func a_div(regs []int, value command_t) {
	regs[A] /= (1 << get_combo(regs, value))
}

func b_div(regs []int, value command_t) {
	regs[B] = regs[A] / (1 << get_combo(regs, value))
}

func c_div(regs []int, value command_t) {
	regs[C] = regs[A] / (1 << get_combo(regs, value))
}

func b_xor(regs []int, value command_t) {
	regs[B] = regs[B] ^ int(value)
}

func b_st(regs []int, value command_t) {
	regs[B] = get_combo(regs, value) % 8
}

func j_nz(regs []int, value command_t) {
	if regs[A] == 0 {
		return
	}
	pcnt = int(value) - 2
}

func b_xorc(regs []int, value command_t) {
	regs[B] ^= regs[C]
}

func out_p(regs []int, value command_t) {
	fmt.Print(get_combo(regs, value)%8, ",")
}

func part1(regs []int, commands []command_t) {
	fmt.Printf("Part 1 solution: ")

	pcnt = 0
	for true {
		if pcnt+1 >= len(commands) {
			break
		}
		cmd := commands[pcnt]
		val := commands[pcnt+1]
		match_func[cmd](regs, val)
		pcnt += 2
	}

	fmt.Println()
}

func check_region(start, n int, commands []command_t) int {
	regs := []int{0, 0, 0}

	for a := start; a < start+n; a++ {
		regs[A] = a
		regs[B] = 0
		regs[C] = 0
		fail := false
		it := 0
		pcnt2 := 0

		for true {
			if pcnt2+1 >= len(commands) {
				break
			}
			cmd := commands[pcnt2]
			val := commands[pcnt2+1]
			if cmd == out {
				// match_func[cmd](regs, val)
				if it == len(commands) {
					fail = true
					// fmt.Println("len1")
					break
				}
				if get_combo(regs, val)%8 != int(commands[it]) {
					fail = true
					// fmt.Println("inc")
					break
				}
				it++
			} else if cmd == jnz {
				if regs[A] != 0 {
					pcnt2 = int(val) - 2
				}
			} else {
				match_func[cmd](regs, val)
			}
			pcnt2 += 2
		}

		if it != len(commands) {
			fail = true
		}

		if !fail {
			return a
		}
	}
	fmt.Println("checked: ", start, start+n)
	return -1
}

func part2(regs []int, commands []command_t) {
	num_workers := 100
	results := make(chan int, num_workers)
	var lowestHit int = -1
	var lowestHitMu sync.Mutex
	var active int = 0
	var activeMu sync.Mutex

	findHits := func(start, end int) {
		hit := check_region(start, end, commands)
		if hit != -1 {
			lowestHitMu.Lock()
			if lowestHit == -1 || hit < lowestHit {
				lowestHit = hit
			}
			lowestHitMu.Unlock()

			select {
			case results <- hit:
			default:
			}
		}
		activeMu.Lock()
		active--
		activeMu.Unlock()
	}

	// don't have time to do something more elaborate
	rl := 10000000
	// calculated first 12 bit by hand on paper, so that brute needs to find the 36 remaining ones
	hardcoded := 0b111_010_110_000
	h2 := hardcoded
	l2 := 0
	for h2 > 0 {
		l2++
		h2 >>= 1
	}
	rem := (len(commands)*3 - l2)

	for i := (hardcoded) << rem; ; i += rl {
		lowestHitMu.Lock()
		if lowestHit != -1 {
			lowestHitMu.Unlock()
			break
		}
		lowestHitMu.Unlock()
		for {
			activeMu.Lock()
			if active < num_workers {
				activeMu.Unlock()
				break
			}
			activeMu.Unlock()
		}

		go findHits(i, rl)
		activeMu.Lock()
		active++
		activeMu.Unlock()

		if i >= (hardcoded+(1<<l2))<<rem {
			fmt.Print("Fuck")
			break
		}
	}

	for {
		activeMu.Lock()
		if active == 0 {
			activeMu.Unlock()
			break
		}
		activeMu.Unlock()
	}

	fmt.Printf("Part 2 solution: %d\n", lowestHit)
}

func main() {
	content, err := os.Open(os.Args[1])
	defer content.Close()
	if err != nil {
		log.Fatal(err)
	}

	regs := []int{0, 0, 0}
	commands := []command_t{}

	scanner := bufio.NewScanner(content)
	for i := 0; i < 3; i++ {
		scanner.Scan()
		s := scanner.Text()[12:]
		v, _ := strconv.Atoi(s)
		regs[i] = v
	}
	scanner.Scan()
	scanner.Scan()
	s := scanner.Text()[9:]
	c := strings.Split(s, ",")
	for _, comm := range c {
		com, _ := strconv.Atoi(comm)
		commands = append(commands, command_t(com))
	}

	part1(regs, commands)
	part2(regs, commands)
}
