package board

import "fmt"

type pos struct {
	row int
	col int
}

var blocks = [9][9]pos{
	// block 0
	[9]pos{
		{0, 0}, {0, 1}, {0, 2},
		{1, 0}, {1, 1}, {1, 2},
		{2, 0}, {2, 1}, {2, 2},
	},
	// block 1
	[9]pos{
		{0, 3}, {0, 4}, {0, 5},
		{1, 3}, {1, 4}, {1, 5},
		{2, 3}, {2, 4}, {2, 5},
	},
	// block 2
	[9]pos{
		{0, 6}, {0, 7}, {0, 8},
		{1, 6}, {1, 7}, {1, 8},
		{2, 6}, {2, 7}, {2, 8},
	},

	// block 3
	[9]pos{
		{3, 0}, {3, 1}, {3, 2},
		{4, 0}, {4, 1}, {4, 2},
		{5, 0}, {5, 1}, {5, 2},
	},
	// block 4
	[9]pos{
		{3, 3}, {3, 4}, {3, 5},
		{4, 3}, {4, 4}, {4, 5},
		{5, 3}, {5, 4}, {5, 5},
	},
	// block 5
	[9]pos{
		{3, 6}, {3, 7}, {3, 8},
		{4, 6}, {4, 7}, {4, 8},
		{5, 6}, {5, 7}, {5, 8},
	},

	// block 6
	[9]pos{
		{6, 0}, {6, 1}, {6, 2},
		{7, 0}, {7, 1}, {7, 2},
		{8, 0}, {8, 1}, {8, 2},
	},
	// block 7
	[9]pos{
		{6, 3}, {6, 4}, {6, 5},
		{7, 3}, {7, 4}, {7, 5},
		{8, 3}, {8, 4}, {8, 5},
	},
	// block 8
	[9]pos{
		{6, 6}, {6, 7}, {6, 8},
		{7, 6}, {7, 7}, {7, 8},
		{8, 6}, {8, 7}, {8, 8},
	},
}

func (bd *Board) PointingElimination(announce bool) int {
	eliminated := 0
	for blockNo := 0; blockNo < 9; blockNo++ {
		var cells [10][]*Cell

		for _, p := range blocks[blockNo] {
			c := &(bd[p.row][p.col])
			if c.Solved {
				continue
			}
			for _, poss := range c.Possible {
				cells[poss] = append(cells[poss], c)
			}

			for poss := 0; poss < 9; poss++ {
				if len(cells[poss]) != 2 {
					continue
				}
				// there's 2 cells each with i as a possiblie value
				c0 := cells[poss][0]
				c1 := cells[poss][1]
				if c0.Row == c1.Row {
					fmt.Printf("Cells <%d,%d> & <%d,%d>, common poss %d, on row %d\n",
						c0.Row, c0.Col, c1.Row, c1.Col, poss, c0.Row,
					)
					row := c0.Row
					for col := 0; col < 9; col++ {
						if bd[row][col].Solved {
							continue
						}
						if bd[row][col].Block == blockNo {
							continue
						}
						m := bd.SpliceOut(row, col, poss)
						if announce && m > 0 {
							fmt.Printf("eliminated %d by pointing at <%d,%d>\n", poss, row, col)
						}
						eliminated += m
					}

					continue
				}
				if c0.Col == c1.Col {
					fmt.Printf("Cells <%d,%d> & <%d,%d>, common poss %d, on col %d\n",
						c0.Row, c0.Col, c1.Row, c1.Col, poss, c0.Col,
					)
					col := c0.Col
					for row := 0; row < 9; row++ {
						if bd[row][col].Solved {
							continue
						}
						if bd[row][col].Block == blockNo {
							continue
						}
						m := bd.SpliceOut(row, col, poss)
						if announce && m > 0 {
							fmt.Printf("eliminated %d by pointing at <%d,%d>\n", poss, row, col)
						}
						eliminated += m
					}
					continue
				}
			}
		}
	}
	return eliminated
}
