package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	input = "L4, R2, R4, L5, L3, L1, R4, R5, R1, R3, L3, L2, L2, R5, R1, L1, L2, R2, R2, L5, R5, R5, L2, R1, R2, L2, L4, L1, R5, R2, R1, R1, L2, L3, R2, L5, L186, L5, L3, R3, L5, R4, R2, L5, R1, R4, L1, L3, R3, R1, L1, R4, R2, L1, L4, R5, L1, R50, L4, R3, R78, R4, R2, L4, R3, L4, R4, L1, R5, L4, R1, L2, R3, L2, R5, R5, L4, L1, L2, R185, L5, R2, R1, L3, R4, L5, R2, R4, L3, R4, L2, L5, R1, R2, L2, L1, L2, R2, L2, R1, L5, L3, L4, L3, L4, L2, L5, L5, R2, L3, L4, R4, R4, R5, L4, L2, R4, L5, R3, R1, L1, R3, L2, R2, R1, R5, L4, R5, L3, R2, R3, R1, R4, L4, R1, R3, L5, L1, L3, R2, R1, R4, L4, R3, L3, R3, R2, L3, L3, R4, L2, R4, L3, L4, R5, R1, L1, R5, R3, R1, R3, R4, L1, R4, R3, R1, L5, L5, L4, R4, R3, L2, R1, R5, L3, R4, R5, L4, L5, R2"
)

type dir int

const (
	north dir = iota // 0
	east             // 1
	south            // 2
	west             // 3
)

type point struct{ x, y int }

type pos struct {
	point
	d dir
}

func (p pos) travelDir(dir string) pos {
	s, d := split(dir)
	if s == "L" {
		return p.ccw().travel(d)
	}
	return p.cw().travel(d)
}

func (p pos) travel(dist int) pos {
	switch p.d {
	case north:
		return pos{point{p.x, p.y + dist}, p.d}
	case east:
		return pos{point{p.x + dist, p.y}, p.d}
	case south:
		return pos{point{p.x, p.y - dist}, p.d}
	default:
		return pos{point{p.x - dist, p.y}, p.d}
	}
}

func (p pos) cw() pos {
	p2 := p
	p2.d = (p2.d + 1) % 4
	return p2
}
func (p pos) ccw() pos {
	p2 := p
	if p2.d == north {
		p2.d = west
	} else {
		p2.d = (p2.d - 1) % 4
	}
	return p2
}

func (p point) dist() int {
	dist := p.x
	if dist < 0 {
		dist *= -1
	}
	if p.y > 0 {
		return dist + p.y
	}
	return dist - p.y
}

func main() {
	p := calcPos(strings.Split(input, ", "))
	fmt.Printf("%#v\n", p)
	fmt.Printf("dist=%d\n", p.dist())

	p = cacheVisits(strings.Split(input, ", "))
	fmt.Printf("%#v\n", p)
	fmt.Printf("dist=%d\n", p.dist())
}

func calcPos(dirs []string) point {
	p := pos{point{0, 0}, north}

	for _, dir := range dirs {
		p = p.travelDir(dir)
	}
	return p.point
}

// cacheVisits is used for part two)
func cacheVisits(dirs []string) point {
	p := pos{point{0, 0}, north}
	visits := map[point]struct{}{
		point{0, 0}: struct{}{},
	}

	for _, dir := range dirs {
		new := p.travelDir(dir)
		for _, pp := range pointsBetween(p.point, new.point) {
			if _, ok := visits[pp]; ok {
				return pp
			}
			visits[pp] = struct{}{}
		}
		p = new
	}
	return p.point

}

func split(s string) (string, int) {
	d := s[:1]
	n := s[1:]
	x, _ := strconv.Atoi(n)
	return d, x
}

func pointsBetween(p1, p2 point) []point {
	var points []point
	if p1.x == p2.x {
		if p1.y > p2.y {
			for i := p1.y - 1; i >= p2.y; i-- {
				points = append(points, point{p1.x, i})
			}
		} else {
			for i := p1.y + 1; i <= p2.y; i++ {
				points = append(points, point{p1.x, i})
			}
		}
	} else {
		if p1.x > p2.x {
			for i := p1.x - 1; i >= p2.x; i-- {
				points = append(points, point{i, p1.y})
			}
		} else {
			for i := p1.x + 1; i <= p2.x; i++ {
				points = append(points, point{i, p1.y})
			}
		}
	}
	return points
}
