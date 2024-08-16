package main

type Grid struct {
	tileMap map[int]map[int]bool
}

func (g *Grid) IsAliveAt(x, y int) bool {
	if g.tileMap[x][y] {
		return true
	}
	return false
}

func (g *Grid) ToggleAt(x, y int) {
	if g.tileMap[x][y] {
		delete(g.tileMap[x], y)
		if len(g.tileMap[x]) == 0 {
			delete(g.tileMap, x)
		}
	} else if _, ok := g.tileMap[x]; ok {
		g.tileMap[x][y] = true
	} else {
		g.tileMap[x] = make(map[int]bool)
		g.tileMap[x][y] = true
	}
}

func (g *Grid) NextGeneration() {
	newMap := make(map[int]map[int]bool)
	alive := [][2]int{}
	for x, col := range g.tileMap {
		for y := range col {
			for i := x - 1; i < x+2; i++ {
				for j := y - 1; j < y+2; j++ {
					if g.tileMap[i][j] {
						if n := g.neighboursAround(i, j); n == 2 || n == 3 {
							alive = append(alive, [2]int{i, j})
						}
					} else if g.neighboursAround(i, j) == 3 {
						alive = append(alive, [2]int{i, j})
					}
				}
			}
		}
	}
	for _, pos := range alive {
		newMap[pos[0]] = make(map[int]bool)
	}
	for _, pos := range alive {
		newMap[pos[0]][pos[1]] = true
	}
	g.tileMap = newMap
}

func (g *Grid) neighboursAround(x, y int) (n int) {
	for i := x - 1; i < x+2; i++ {
		for j := y - 1; j < y+2; j++ {
			if g.tileMap[i][j] && (i != x || j != y) {
				n++
			}
		}
	}
	return n
}

func (g *Grid) BoundedView(x, y, width, height int) (s string) {
	if height <= 0 || width <= 0 {
		return ""
	}
	for j := range height {
		for i := range width {
			if _, ok := g.tileMap[x+i][y+j]; ok {
				s += "██"
			} else {
				s += "  "
			}
		}
		s += "\n"
	}
	return s[:len(s)-1]
}
