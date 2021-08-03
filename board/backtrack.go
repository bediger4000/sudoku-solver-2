package board

import "os"

func BackTrackSolution(bd *Board) {
	solution, foundit := backTrackSolution(*bd)
	if foundit {
		solution.Print(os.Stdout)
	}
}

func backTrackSolution(bd Board) (Board, bool) {
	for rowNo := 0; rowNo < 9; rowNo++ {
		for colNo := 0; colNo < 9; colNo++ {
			if !bd[rowNo][colNo].Solved {
				for _, digit := range bd[rowNo][colNo].Possible {
					// set digit as bd[][].Value
					bd[rowNo][colNo].Value = digit
					bd[rowNo][colNo].Solved = true

					// erase all possibilities this affects
					erasures := (&bd).erasePossibilities(rowNo, colNo, bd[rowNo][colNo].Block, digit)

					if len(erasures) > 0 {
						// check to see if this is a solution
						if valid, complete := (&bd).ValidAndComplete(); complete {
							return bd, valid
						}

						// recurse
						if solvedBd, solution := backTrackSolution(bd); solution {
							return solvedBd, true
						}

						// reset all the erased possibilities
						(&bd).replaceEliminations(erasures)
					}

					// reset bd[][].Value
					bd[rowNo][colNo].Value = 0
					bd[rowNo][colNo].Solved = false
				}
			}
		}
	}
	return bd, false
}

func (bd *Board) replaceEliminations(eliminations [][3]int) {
	for idx := range eliminations {
		row, col, digit := eliminations[idx][0], eliminations[idx][1], eliminations[idx][2]
		bd[row][col].Possible = append(bd[row][col].Possible, digit)
	}
}

// erasePossibilities erases all instances of digitEliminate
// from other squares in rowEliminate, colEliminate and blockEliminate,
// returning a slice of [3]int: {row, col, digit}, each representing
// an erased digit in square [row][col]
func (bd *Board) erasePossibilities(rowEliminate, colEliminate, blockEliminate, digitEliminate int) [][3]int {
	eliminations := make([][3]int, 0)
	for col := 0; col < 9; col++ {
		if bd[rowEliminate][col].Solved {
			continue
		}
		if col == colEliminate {
			continue
		}
		if bd.erase(rowEliminate, col, digitEliminate) {
			eliminations = append(eliminations, [3]int{rowEliminate, col, digitEliminate})
		}
	}
	for row := 0; row < 9; row++ {
		if bd[row][colEliminate].Solved {
			continue
		}
		if row == rowEliminate {
			continue
		}
		if bd.erase(row, colEliminate, digitEliminate) {
			eliminations = append(eliminations, [3]int{row, colEliminate, digitEliminate})
		}
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
				if bd.erase(row, col, digitEliminate) {
					eliminations = append(eliminations, [3]int{row, col, digitEliminate})
				}
			}
		}
	}

	return eliminations
}

// erase will eliminate at most 1 digit from
// the bd[row][col].Possible slice
func (bd *Board) erase(row, col, digitEliminate int) bool {
	for idx, digit := range bd[row][col].Possible {
		if digit == digitEliminate {
			bd[row][col].Possible = append(bd[row][col].Possible[:idx], bd[row][col].Possible[idx+1:]...)
			return true
		}
	}
	return false
}
