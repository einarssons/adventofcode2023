package utils

type Grid3D[K any] struct {
	G   [][][]K
	Min [3]int
	Max [3]int
}

var NeighborsStraight = [][3]int{{-1, 0, 0}, {1, 0, 0}, {0, -1, 0}, {0, 1, 0}, {0, 0, -1}, {0, 0, 1}}
var NeighborsAll [][3]int

func init() {
	for x := -1; x <= 1; x += 2 {
		for y := -1; y <= 1; y += 2 {
			for z := -1; z <= 1; z += 2 {
				NeighborsAll = append(NeighborsAll, [3]int{x, y, z})
			}
		}
	}
}

func NewGrid[K any](min, max [3]int) *Grid3D[K] {
	g := make([][][]K, 0, max[0]-min[0]+1)
	for x := min[0]; x <= max[0]; x++ {
		yw := make([][]K, 0, max[1]-min[1]+1)
		for y := min[1]; y <= max[1]; y++ {
			yw = append(yw, make([]K, max[2]-min[2]+1))
		}
		g = append(g, yw)
	}
	return &Grid3D[K]{
		G:   g,
		Min: min,
		Max: max,
	}
}

func (g *Grid3D[K]) Set(v K, x, y, z int) {
	g.G[x-g.Min[0]][y-g.Min[1]][z-g.Min[2]] = v
}

func (g *Grid3D[K]) Get(x, y, z int) K {
	return g.G[x-g.Min[0]][y-g.Min[1]][z-g.Min[2]]
}

func (g *Grid3D[K]) IsEdge(x, y, z int) bool {
	low := x == g.Min[0] || y == g.Min[1] || z == g.Min[2]
	high := x == g.Max[0] || y == g.Max[1] || z == g.Max[2]
	return low || high
}

func (g *Grid3D[K]) IsOut(x, y, z int) bool {
	if x < g.Min[0] || y < g.Min[1] || z < g.Min[2] {
		return true
	}
	if x > g.Max[0] || y > g.Max[1] || z > g.Max[2] {
		return true
	}
	return false
}

func (g *Grid3D[K]) Volume() int {
	f := 1
	for i := 0; i < 3; i++ {
		f *= g.Max[i] - g.Min[i] + 1
	}
	return f
}
