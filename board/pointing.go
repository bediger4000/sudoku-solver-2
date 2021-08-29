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
		block := blocks[blockNo]

		m := make(map[int][]pos)

		for pi := 0; pi < 9; pi++ {
			if bd[block[pi].row][block[pi].col].Solved {
				continue
			}
			for _, possible := range bd[block[pi].row][block[pi].col].Possible {
				m[possible] = append(m[possible], block[pi])
			}
		}

		for possible, positions := range m {
			if len(positions) == 2 {
				if positions[0].row == positions[1].row {
					rowEliminate := positions[0].row
					if announce {
						fmt.Printf("row %d pointing elimination %ds because <%d,%d> and <%d,%d>\n",
							rowEliminate,
							possible,
							positions[0].row, positions[0].col,
							positions[1].row, positions[1].col,
						)
					}

					for colNo := 0; colNo < 9; colNo++ {
						if rowEliminate == positions[0].row && colNo == positions[0].col {
							continue
						}
						if rowEliminate == positions[1].row && colNo == positions[1].col {
							continue
						}
						x := bd.SpliceOut(rowEliminate, colNo, possible)
						if announce && x == 1 {
							fmt.Printf("Eliminated %d at <%d,%d> due to pointing\n",
								possible, rowEliminate, colNo)
						}
						eliminated += x
					}
				}
				if positions[0].col == positions[1].col {
					colEliminate := positions[0].col
					if announce {
						fmt.Printf("col %d pointing elimination of %ds because <%d,%d> and <%d,%d>\n",
							colEliminate,
							possible,
							positions[0].row, positions[0].col,
							positions[1].row, positions[1].col,
						)
					}

					for rowNo := 0; rowNo < 9; rowNo++ {
						if rowNo == positions[0].row && colEliminate == positions[0].col {
							continue
						}
						if rowNo == positions[1].row && colEliminate == positions[1].col {
							continue
						}
						x := bd.SpliceOut(rowNo, colEliminate, possible)
						if announce && x == 1 {
							fmt.Printf("Eliminated %d at <%d,%d> due to pointing\n",
								possible, rowNo, colEliminate)
						}
						eliminated += x
					}
				}
			}
		}
	}
	return eliminated
}
