package board

import (
	"fmt"
	"os"
)

type replacement struct {
	row      int
	col      int
	possible []int
}

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
	bd.printPly(ply, "enter backTrackSolution")
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
			possibleDigits := make([]int, len(bd[rowNo][colNo].Possible))
			copy(possibleDigits, bd[rowNo][colNo].Possible)
			for _, digit := range possibleDigits {
				// set digit as bd[][].Value
				fmt.Printf("\tply %d: set <%d,%d> to %d\n", ply, rowNo, colNo, digit)
				bd[rowNo][colNo].Value = digit
				bd[rowNo][colNo].Solved = true

				// erase all possibilities this affects
				erasures := bd.erasePossibilities(rowNo, colNo, bd[rowNo][colNo].Block, digit)
				fmt.Printf("\tply %d: %d erasures\n", ply, len(erasures))
				for _, p := range erasures {
					fmt.Printf("\t<%d,%d> had %v, has %v\n",
						p.row, p.col, p.possible, bd[rowNo][colNo].Possible)
				}

				// check to see if some square has no further possibilities
				erasedToInvalid := false
				if len(erasures) > 0 {
				FOUNDINVALID:
					for r := 0; r < 9; r++ {
						for c := 0; c < 9; c++ {
							if bd[r][c].Solved {
								continue
							}
							if len(bd[r][c].Possible) == 0 {
								fmt.Printf("\tply %d: <%d,%d>->%d, no possibilites at <%d,%d>\n", ply, rowNo, colNo, digit, r, c)
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

func (bd *Board) replaceEliminations(eliminations []*replacement) {
	for idx := range eliminations {
		e := eliminations[idx]
		bd[e.row][e.col].Possible = e.possible
	}
}

// erasePossibilities erases all instances of digitEliminate
// from other squares in rowEliminate, colEliminate and blockEliminate,
// returning a slice of []int: {row, col, original []int}. The
// erased digit in square [row][col] will be missing from slice
// bd[row][col].Possible. This returns the original .Possible element
func (bd *Board) erasePossibilities(rowEliminate, colEliminate, blockEliminate, digitEliminate int) []*replacement {
	var replace []*replacement
	for col := 0; col < 9; col++ {
		if bd[rowEliminate][col].Solved {
			continue
		}
		if col == colEliminate {
			continue
		}
		if possible := bd.erase(rowEliminate, col, digitEliminate); len(possible) > 0 {
			fmt.Printf("erased %d at <%d,%d>: %v\n",
				digitEliminate, rowEliminate, col, bd[rowEliminate][col].Possible)
			replace = append(replace, &replacement{rowEliminate, col, possible})
		}
	}
	for row := 0; row < 9; row++ {
		if bd[row][colEliminate].Solved {
			continue
		}
		if row == rowEliminate {
			continue
		}
		if possible := bd.erase(row, colEliminate, digitEliminate); len(possible) > 0 {
			fmt.Printf("erased %d at <%d,%d>: %v\n",
				digitEliminate, row, colEliminate, bd[row][colEliminate].Possible)
			replace = append(replace, &replacement{row, colEliminate, possible})
		}
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if row == rowEliminate && col == colEliminate {
				continue
			}
			if bd[row][col].Block != blockEliminate {
				continue
			}
			if bd[row][col].Solved {
				continue
			}
			if possible := bd.erase(row, col, digitEliminate); len(possible) > 0 {
				fmt.Printf("erased %d at <%d,%d>: %v\n",
					digitEliminate, row, col, bd[row][col].Possible)
				replace = append(replace, &replacement{row, col, possible})
			}
		}
	}

	return replace
}

func (bd *Board) erase(row, col, digitEliminate int) []int {
	l := len(bd[row][col].Possible)
	for i := 0; i < l; i++ {
		if bd[row][col].Possible[i] == digitEliminate {
			fmt.Printf("\terase %d from %v at <%d,%d>\n", bd[row][col].Possible[i], bd[row][col].Possible, row, col)
			// give it a new .Possible array missing digitEliminate
			newpossibles := make([]int, l-1)
			j := 0
			for k := 0; k < i; k++ {
				newpossibles[j] = bd[row][col].Possible[k]
				j++
			}
			for k := i + 1; k < l; k++ {
				newpossibles[j] = bd[row][col].Possible[k]
				j++
			}
			tmp := bd[row][col].Possible
			bd[row][col].Possible = newpossibles
			fmt.Printf("now has %v\n", bd[row][col].Possible)
			return tmp
		}
	}
	return nil
}

/*
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
*/
