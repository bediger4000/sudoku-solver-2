package board

import (
	"fmt"
)

// Sudoku board
// A board is 81 cells, in a 9x9 grid.
// When solved, each cell holds a single
// digit 1-9 inclusive. Until it's solved,
// it has up to 9 possiblities.
// Row and Col constitute <x,y> coords in
// the 9x9 grid. There are 9, 3x3 sub-grids
// called "blocks", numbered like this:
// 0 1 2
// 3 4 5
// 6 7 8

type Cell struct {
	Row      int
	Col      int
	Block    int
	Possible []int // logically possible digits
	Solved   bool
	Value    int
}

type Board [9][9]Cell

// Block number indexed by <x,y> cell
// coordinates. This array used to initialize
// the 9x9 grid of cells, but Cell.Block is
// used all over the place.
var blockNumber = [9][9]int{
	[9]int{0, 0, 0, 1, 1, 1, 2, 2, 2},
	[9]int{0, 0, 0, 1, 1, 1, 2, 2, 2},
	[9]int{0, 0, 0, 1, 1, 1, 2, 2, 2},
	[9]int{3, 3, 3, 4, 4, 4, 5, 5, 5},
	[9]int{3, 3, 3, 4, 4, 4, 5, 5, 5},
	[9]int{3, 3, 3, 4, 4, 4, 5, 5, 5},
	[9]int{6, 6, 6, 7, 7, 7, 8, 8, 8},
	[9]int{6, 6, 6, 7, 7, 7, 8, 8, 8},
	[9]int{6, 6, 6, 7, 7, 7, 8, 8, 8},
}

/*
var positionByBlock = [2][9][9]int{
	[2][9]int{{0, 0}, {0, 1}, {0, 3}, {1, 0}, {1, 1}, {1, 3}, {2, 0}, {2, 1}, {2, 3}},
	[2][9]int{{0, 3}, {0, 4}, {0, 5}, {1, 3}, {1, 4}, {1, 5}, {2, 3}, {2, 4}, {2, 5}},

	[2][9]int{{0, 0}, {0, 1}, {0, 3}, {1, 0}, {1, 1}, {1, 3}, {2, 0}, {2, 1}, {2, 3}},
	[2][9]int{{0, 0}, {0, 1}, {0, 3}, {1, 0}, {1, 1}, {1, 3}, {2, 0}, {2, 1}, {2, 3}},
	[2][9]int{{0, 0}, {0, 1}, {0, 3}, {1, 0}, {1, 1}, {1, 3}, {2, 0}, {2, 1}, {2, 3}},
	[2][9]int{{0, 0}, {0, 1}, {0, 3}, {1, 0}, {1, 1}, {1, 3}, {2, 0}, {2, 1}, {2, 3}},
	[2][9]int{{0, 0}, {0, 1}, {0, 3}, {1, 0}, {1, 1}, {1, 3}, {2, 0}, {2, 1}, {2, 3}},
	[2][9]int{{0, 0}, {0, 1}, {0, 3}, {1, 0}, {1, 1}, {1, 3}, {2, 0}, {2, 1}, {2, 3}},
	[2][9]int{{0, 0}, {0, 1}, {0, 3}, {1, 0}, {1, 1}, {1, 3}, {2, 0}, {2, 1}, {2, 3}},
}
*/

func New() *Board {
	var bd Board
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			bd[i][j].Row = i
			bd[i][j].Col = j
			bd[i][j].Possible = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			bd[i][j].Block = blockNumber[i][j]
		}
	}
	return &bd
}

func Copy(bd *Board) *Board {
	n := New()
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			copy(n[i][j].Possible, bd[i][j].Possible)
			n[i][j].Possible = n[i][j].Possible[:len(bd[i][j].Possible)]
			n[i][j].Solved = bd[i][j].Solved
			n[i][j].Value = bd[i][j].Value
		}
	}
	return n
}

func Compare(bd1, bd2 *Board, verbose bool) bool {
	identical := true
	for rowNo := 0; rowNo < 9; rowNo++ {
		for colNo := 0; colNo < 9; colNo++ {
			c1 := ((*bd1)[rowNo][colNo])
			c2 := ((*bd2)[rowNo][colNo])
			if c1.Row != c2.Row {
				if verbose {
					fmt.Printf("%d,%d: board 1 row %d, board 2 row %d\n",
						rowNo, colNo,
						c1.Row, c2.Row,
					)
				}
				identical = false
			}
			if c1.Col != c2.Col {
				if verbose {
					fmt.Printf("%d,%d: board 1 column %d, board 2 column %d\n",
						rowNo, colNo,
						c1.Col, c2.Col,
					)
				}
				identical = false
			}
			if c1.Block != c2.Block {
				if verbose {
					fmt.Printf("%d,%d: board 1 block %d, board 2 block %d\n",
						rowNo, colNo,
						c1.Block, c2.Block,
					)
				}
				identical = false
			}
			if c1.Solved != c2.Solved {
				if verbose {
					fmt.Printf("%d,%d: board 1 solved %v, board 2 solved %v\n",
						rowNo, colNo,
						c1.Solved, c2.Solved,
					)
				}
				identical = false
			}
			if c1.Value != c2.Value {
				if verbose {
					fmt.Printf("%d,%d: board 1 solved value %d, board 2 solved value %d\n",
						rowNo, colNo,
						c1.Value, c2.Value,
					)
				}
				identical = false
			}
			if len(c1.Possible) != len(c2.Possible) {
				identical = false
				if verbose {
					fmt.Printf("board 1 possible: %v\n", c1.Possible)
					fmt.Printf("board 2 possible: %v\n", c2.Possible)
				}
			} else {
				for i := 0; i < len(c1.Possible); i++ {
					if c1.Possible[i] != c2.Possible[i] {
						identical = false
						if verbose {
							fmt.Printf("%d,%d: board 1 has possible values %v, board 2 has possible values %v\n",
								rowNo, colNo,
								c1.Possible, c2.Possible,
							)
							break
						}
					}
				}
			}
		}
	}

	return identical
}
