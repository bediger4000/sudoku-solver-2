package board

func (bd *Board) EliminatePossibilities(rowEliminate, colEliminate, blockEliminate, digitEliminate int) {
	for col := 0; col < 9; col++ {
		if bd[rowEliminate][col].Solved {
			continue
		}
		if col == colEliminate {
			continue
		}
		bd.SpliceOut(rowEliminate, col, digitEliminate)
	}
	for row := 0; row < 9; row++ {
		if bd[row][colEliminate].Solved {
			continue
		}
		if row == rowEliminate {
			continue
		}
		bd.SpliceOut(row, colEliminate, digitEliminate)
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if row == rowEliminate && col == colEliminate {
				continue
			}
			if bd[row][col].Solved {
				continue
			}
			if bd[row][col].Block == blockEliminate {
				bd.SpliceOut(row, col, digitEliminate)
			}
		}
	}
}

// SpliceOut will eliminate at most 1 digit from
// the bd[row][col].Possible slice
func (bd *Board) SpliceOut(row, col, digitEliminate int) int {
	for idx, digit := range bd[row][col].Possible {
		if digit == digitEliminate {
			bd[row][col].Possible = append(bd[row][col].Possible[:idx], bd[row][col].Possible[idx+1:]...)
			return 1
		}
	}
	return 0
}
