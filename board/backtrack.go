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

// Track unique boards
var uniqueBoards = make(map[[81]byte]bool)
var uniqueBoardCount int
var invalidBoardCount int

// stringify makes an 81-byte slice with byte-numerical-values from
// the solved .Value of each Cell in a Board.
func (bd *Board) stringify() [81]byte {
	var key [81]byte
	for rowNo := 0; rowNo < 9; rowNo++ {
		for colNo := 0; colNo < 9; colNo++ {
			key[rowNo*9+colNo] = byte(bd[rowNo][colNo].Value)
		}
	}
	return key
}

// printUniqueBoards prints only the unique boards it gets passed.
func printUniqueBoards(bd *Board) {
	boardKey := bd.stringify()
	if !uniqueBoards[boardKey] {
		uniqueBoardCount++
		fmt.Printf("*** backtracking solution %d\n", uniqueBoardCount)
		bd.Print(os.Stdout)
		uniqueBoards[boardKey] = true
	}
}

func (bd *Board) printPly(ply int, phrase string) {
	fmt.Printf("ply %d: %s\n", ply, phrase)
	bd.Print(os.Stdout)
	openCount := 0
	possibleCombinations := 1
	for rowNo := 0; rowNo < 9; rowNo++ {
		for colNo := 0; colNo < 9; colNo++ {
			if !bd[rowNo][colNo].Solved {
				fmt.Printf("  <%d,%d> can hold %v\n", rowNo, colNo, bd[rowNo][colNo].Possible)
				openCount++
				possibleCombinations *= len(bd[rowNo][colNo].Possible)
			}
		}
	}
	fmt.Printf("ply %d, %d unsolved cells, might take %d attempts\n", ply, openCount, possibleCombinations)
}

func BackTrackSolution(bd *Board) {
	bd.printPly(-1, "start backtracking")
	fmt.Println("===")
	backTrackSolution(0, bd)
	fmt.Print("===")
	if uniqueBoardCount > 0 {
		fmt.Printf(" %d backtracking solutions\n", uniqueBoardCount)
		return
	}
	fmt.Println()
}

// backTrackSolution called recursively to find all valid boards.
// It can find the same board more than once, because it looks at
// each unsolved square.
func backTrackSolution(ply int, bd *Board) {
	if invalidBoardCount > 1 && invalidBoardCount%1000000 == 0 {
		fmt.Printf("%d invalid boards found so far\n", invalidBoardCount)
	}
	for rowNo := 0; rowNo < 9; rowNo++ {
		for colNo := 0; colNo < 9; colNo++ {
			if bd[rowNo][colNo].Solved {
				continue
			}
			if len(bd[rowNo][colNo].Possible) == 0 {
				// no possibilities: this position is invalid
				invalidBoardCount++
				return
			}
			// copy .Possible slice because it gets cut up recursively.
			possibleDigits := make([]int, len(bd[rowNo][colNo].Possible))
			copy(possibleDigits, bd[rowNo][colNo].Possible)
			for _, digit := range possibleDigits {
				// set digit as bd[][].Value
				bd[rowNo][colNo].Value = digit
				bd[rowNo][colNo].Solved = true

				// erase all possibilities this affects
				erasures := bd.erasePossibilities(rowNo, colNo, bd[rowNo][colNo].Block, digit)

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
					invalidBoardCount++
					continue
				}

				// check to see if this is a solution
				if valid, complete := bd.ValidAndComplete(); complete {
					invalidBoardCount++
					if valid {
						printUniqueBoards(bd)
						invalidBoardCount--
					}
				} else {
					// it's incomplete, open squares remain, recurse
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
			// give it a new .Possible array missing digitEliminate
			newpossibles := make([]int, l-1)
			j := 0
			for k := 0; k < i; k++ {
				newpossibles[j] = bd[row][col].Possible[k]
				j++
			}
			// skip bd[row][col].Possible[i] - it contains digitEliminate
			for k := i + 1; k < l; k++ {
				newpossibles[j] = bd[row][col].Possible[k]
				j++
			}
			tmp := bd[row][col].Possible
			bd[row][col].Possible = newpossibles
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
