package board

import (
	"fmt"
	"os"
)

func (bd *Board) printPly(ply int, phrase string) {
	fmt.Printf("ply %d: %s\n", ply, phrase)
	bd.Print(os.Stdout)
	openCount := 0
	for rowNo := 0; rowNo < 9; rowNo++ {
		for colNo := 0; colNo < 9; colNo++ {
			if !bd[rowNo][colNo].Solved {
				fmt.Printf("  <%d,%d> can hold %v\n", rowNo, colNo, bd[rowNo][colNo].Possible)
				openCount++
			}
		}
	}
	fmt.Printf("ply %d, %d unsolved cells\n", ply, openCount)
}

func BackTrackSolution(bd *Board) {
	bd.printPly(-1, "start backtracking")
	backTrackSolution(0, bd)
	fmt.Println("===")
}

func backTrackSolution(ply int, bd *Board) {
	if valid, complete := bd.ValidAndComplete(); complete {
		if valid {
			// a solution
			fmt.Println("---")
			bd.Print(os.Stdout)
			return
		}
		// complete, but invalid
		return
	}
	for rowNo := 0; rowNo < 9; rowNo++ {
		for colNo := 0; colNo < 9; colNo++ {
			if bd[rowNo][colNo].Solved {
				continue
			}
			if len(bd[rowNo][colNo].Possible) == 0 {
				// no possibilities: this position is invalid
				return
			}
			for _, digit := range bd[rowNo][colNo].Possible {
				// set digit as bd[][].Value
				bd[rowNo][colNo].Value = digit
				bd[rowNo][colNo].Solved = true

				// erase all possibilities this affects
				erasures := bd.erasePossibilities(rowNo, colNo, bd[rowNo][colNo].Block, digit)
				erasedToInvalid := false
				if len(erasures) > 0 {
				FOUNDINVALID:
					for r := 0; r < 9; r++ {
						for c := 0; c < 9; c++ {
							if bd[r][c].Solved {
								continue
							}
							if len(bd[r][c].Possible) == 0 {
								// erasures made an invalid position
								erasedToInvalid = true
								break FOUNDINVALID
							}
						}
					}
				}
				if erasedToInvalid {
					bd.replaceEliminations(erasures)
					bd[rowNo][colNo].Value = 0
					bd[rowNo][colNo].Solved = false
					continue
				}

				// check to see if this is a solution
				if valid, complete := bd.ValidAndComplete(); complete {
					if valid {
						fmt.Println("***")
						bd.Print(os.Stdout)
					}
				} else {
					// recurse
					backTrackSolution(ply+1, bd)
				}

				// reset all the erased possibilities
				if len(erasures) > 0 {
					bd.replaceEliminations(erasures)
				}

				// reset bd[][].Value
				bd[rowNo][colNo].Value = 0
				bd[rowNo][colNo].Solved = false
			}
		}
	}
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
