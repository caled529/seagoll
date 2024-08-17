package main

type Grid map[int]map[int]bool

func (g Grid) NextState() (newGrid Grid) {
	living := [][2]int{}
	for x, col := range g {
		for y := range col {
			if n := g.countNeighbours(x, y); n == 2 || n == 3 {
				living = append(living, [2]int{x, y})
			}
			living = append(living, g.localNextState(x, y)...)
		}
	}
	newGrid = make(map[int]map[int]bool)
	for _, pos := range living {
		if _, ok := newGrid[pos[0]]; !ok {
			newGrid[pos[0]] = make(map[int]bool)
		}
	}
	for _, pos := range living {
		newGrid[pos[0]][pos[1]] = true
	}
	return newGrid
}

func (g Grid) localNextState(x, y int) (living [][2]int) {
	for i := x - 1; i < x+2; i++ {
		for j := y - 1; j < y+2; j++ {
			if i == x && j == y {
				continue
			}
			if !g[i][j] && g.countNeighbours(i, j) == 3 {
				living = append(living, [2]int{i, j})
			}
		}
	}
	return living
}

func (g Grid) countNeighbours(x, y int) (n int) {
	for i := x - 1; i < x+2; i++ {
		for j := y - 1; j < y+2; j++ {
			if g[i][j] && (i != x || j != y) {
				n++
			}
		}
	}
	return n
}

func (g Grid) ToggleAt(x, y int) {
	if g[x][y] {
		delete(g[x], y)
		if len(g[x]) == 0 {
			delete(g, x)
		}
	} else if _, ok := g[x]; ok {
		g[x][y] = true
	} else {
		g[x] = make(map[int]bool)
		g[x][y] = true
	}
}

func (g Grid) IsAlive(x, y int) bool {
	if g[x][y] {
		return true
	}
	return false
}

func (g Grid) BoundedView(x, y, width, height int) (s string) {
	if height <= 0 || width <= 0 {
		return ""
	}
	for j := range height {
		for i := range width {
			if _, ok := g[x+i][y+j]; ok {
				s += "██"
			} else {
				s += "  "
			}
		}
		s += "\n"
	}
	return s[:len(s)-1]
}
