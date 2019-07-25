package blockastar

import (
	"fmt"
	"os"
)

// Grid Type
type Grid struct {
	R      int
	C      int
	K      int
	Cells  [][]*Cell
	Blocks [][]*Block
}

// Print grid data
func (g Grid) Print() {
	for i := 0; i < g.R; i++ {
		for j := 0; j < g.C; j++ {
			if g.Cells[i][j].isEmpty {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
	fmt.Println()

	for i := 0; i < g.R/g.K; i++ {
		for j := 0; j < g.C/g.K; j++ {
			block := g.Blocks[i][j]
			fmt.Println(block)
			for ii := 0; ii < g.K; ii++ {
				for jj := 0; jj < g.K; jj++ {
					cell := block.cells[ii][jj]
					fmt.Println("    ", cell, " -> ", cell.sides)
				}
			}
		}
	}

}

// BuildGridFromMaps return Grid generated from map file with size block K
func BuildGridFromMaps(path string, K int) Grid {
	var (
		err             error
		height, width   int
		direction, line string
		di              = [8]int{-1, 1, 0, 0, -1, 1, -1, 1}
		dj              = [8]int{0, 0, -1, 1, -1, -1, 1, 1}
	)

	// redirect input from file
	os.Stdin, err = os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// read data from map file and build cells
	fmt.Scanf("type %s\n", &direction)
	fmt.Scanf("height %d\n", &height)
	fmt.Scanf("width %d\n", &width)
	fmt.Scanf("map\n")
	R := ((height + K - 1) / K) * K
	C := ((width + K - 1) / K) * K
	cells := make([][]*Cell, R)
	for i := 0; i < R; i++ {
		cells[i] = make([]*Cell, C)
		fmt.Scanf("%s\n", &line)
		for j := 0; j < C; j++ {
			isEmpty := false
			if (i < height) && (j < width) {
				isEmpty = (line[j] == '.' || line[j] == 'S')
			}
			cells[i][j] = &Cell{
				index:   -1,
				x:       i,
				y:       j,
				g:       INF,
				isEmpty: isEmpty,
			}
		}
	}

	// build cells neighbours in his block
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if cells[i][j].isEmpty {
				if (direction == "tile") || (direction == "octile") {
					for d := 0; d < 4; d++ {
						ii, jj := i+di[d], j+dj[d]
						if (ii >= 0) && (ii < R) && (jj >= 0) && (jj < C) {
							if (ii/K == i/K) && (jj/K == j/K) {
								if cells[ii][jj].isEmpty {
									cells[i][j].sides = append(cells[i][j].sides, cells[ii][jj])
								}
							}
						}
					}
				}
				if direction == "octile" {
					for d := 4; d < 8; d++ {
						ii, jj := i+di[d], j+dj[d]
						if (ii >= 0) && (ii < R) && (jj >= 0) && (jj < C) {
							if (ii/K == i/K) && (jj/K == j/K) {
								if cells[i][jj].isEmpty && cells[ii][j].isEmpty && cells[ii][jj].isEmpty {
									cells[i][j].sides = append(cells[i][j].sides, cells[ii][jj])
								}
							}
						}
					}
				}
			}
		}
	}

	// build blocks
	blocks := make([][]*Block, R/K)
	for i := 0; i < R/K; i++ {
		blocks[i] = make([]*Block, C/K)
		for j := 0; j < C/K; j++ {
			blocks[i][j] = &Block{
				index:     -1,
				x:         i,
				y:         j,
				K:         K,
				heapvalue: INF,
				ingress:   map[*Cell]void{},
			}
			blocks[i][j].cells = make([][]*Cell, K)
			for ii := 0; ii < K; ii++ {
				blocks[i][j].cells[ii] = make([]*Cell, K)
				for jj := 0; jj < K; jj++ {
					cells[i*K+ii][j*K+jj].block = blocks[i][j]
					blocks[i][j].cells[ii][jj] = cells[i*K+ii][j*K+jj]
				}
			}
		}
	}

	// build blocks neighbours
	for i := 0; i < R/K; i++ {
		for j := 0; j < C/K; j++ {
			if (direction == "tile") || (direction == "octile") {
				for d := 0; d < 4; d++ {
					ii, jj := i+di[d], j+dj[d]
					if (ii >= 0) && (ii < R/K) && (jj >= 0) && (jj < C/K) {
						neighbour := blocks[ii][jj]
						egress := make([]Egress, 0)
						for l := 0; l < K; l++ {
							var (
								ein, eout, einki, einka, eoutki, eoutka *Cell
								out                                     []*Cell
							)
							switch d {
							case 0:
								ein = cells[i*K][j*K+l]
								eout = cells[i*K-1][j*K+l]
								if l > 0 {
									einki = cells[i*K][j*K+l-1]
									eoutki = cells[i*K-1][j*K+l-1]
								}
								if l < K-1 {
									einka = cells[i*K][j*K+l+1]
									eoutka = cells[i*K-1][j*K+l+1]
								}
							case 1:
								ein = cells[(i+1)*K-1][j*K+l]
								eout = cells[(i+1)*K][j*K+l]
								if l > 0 {
									einki = cells[(i+1)*K-1][j*K+l-1]
									eoutki = cells[(i+1)*K][j*K+l-1]
								}
								if l < K-1 {
									einka = cells[(i+1)*K-1][j*K+l+1]
									eoutka = cells[(i+1)*K][j*K+l+1]
								}
							case 2:
								ein = cells[i*K+l][j*K]
								eout = cells[i*K+l][j*K-1]
								if l > 0 {
									einki = cells[i*K+l-1][j*K]
									eoutki = cells[i*K+l-1][j*K-1]
								}
								if l < K-1 {
									einka = cells[i*K+l+1][j*K]
									eoutka = cells[i*K+l+1][j*K-1]
								}
							case 3:
								ein = cells[i*K+l][(j+1)*K-1]
								eout = cells[i*K+l][(j+1)*K]
								if l > 0 {
									einki = cells[i*K+l-1][(j+1)*K-1]
									eoutki = cells[i*K+l-1][(j+1)*K]
								}
								if l < K-1 {
									einka = cells[i*K+l+1][(j+1)*K-1]
									eoutka = cells[i*K+l+1][(j+1)*K]
								}
							}
							if ein.isEmpty && eout.isEmpty {
								out = append(out, eout)
							}
							if l > 0 {
								if ein.isEmpty && eout.isEmpty && einki.isEmpty && eoutki.isEmpty {
									out = append(out, eoutki)
								}
							}
							if l < K-1 {
								if ein.isEmpty && eout.isEmpty && einka.isEmpty && eoutka.isEmpty {
									out = append(out, eoutka)
								}
							}
							egress = append(egress, Egress{ein, out})
						}
						blocks[i][j].sides = append(blocks[i][j].sides, struct {
							neighbour *Block
							egress    []Egress
						}{
							neighbour: neighbour,
							egress:    egress,
						})
					}
				}
			}
			if direction == "octile" {
				for d := 4; d < 8; d++ {
					ii, jj := i+di[d], j+dj[d]
					if (ii >= 0) && (ii < R/K) && (jj >= 0) && (jj < C/K) {
						neighbour := blocks[ii][jj]
						egress := make([]Egress, 0)
						var ein, eout, eside1, eside2 *Cell
						switch d {
						case 4:
							ein = cells[i*K][j*K]
							eside1 = cells[i*K][j*K-1]
							eside2 = cells[i*K-1][j*K]
							eout = cells[i*K-1][j*K-1]
						case 5:
							ein = cells[(i+1)*K-1][j*K]
							eside1 = cells[(i+1)*K][j*K]
							eside2 = cells[(i+1)*K-1][j*K-1]
							eout = cells[(i+1)*K][j*K-1]
						case 6:
							ein = cells[i*K][(j+1)*K-1]
							eside1 = cells[i*K-1][(j+1)*K-1]
							eside2 = cells[i*K][(j+1)*K]
							eout = cells[i*K-1][(j+1)*K]
						case 7:
							ein = cells[(i+1)*K-1][(j+1)*K-1]
							eside1 = cells[(i+1)*K][(j+1)*K-1]
							eside2 = cells[(i+1)*K-1][(j+1)*K]
							eout = cells[(i+1)*K][(j+1)*K]
						}
						if ein.isEmpty && eout.isEmpty && eside1.isEmpty && eside2.isEmpty {
							egress = append(egress, Egress{ein, []*Cell{eout}})
						}
						blocks[i][j].sides = append(blocks[i][j].sides, struct {
							neighbour *Block
							egress    []Egress
						}{
							neighbour: neighbour,
							egress:    egress,
						})
					}
				}

			}
		}
	}

	return Grid{
		R:      R,
		C:      C,
		K:      K,
		Cells:  cells,
		Blocks: blocks,
	}
}
