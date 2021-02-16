package board

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
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

func New() Board {
	var bd Board
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			bd[i][j].Row = i
			bd[i][j].Col = j
			bd[i][j].Possible = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			bd[i][j].Block = blockNumber[i][j]
		}
	}
	return bd
}

func (bd Board) Print(out io.Writer) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if bd[i][j].Solved {
				fmt.Fprintf(out, "%1d ", bd[i][j].Value)
			} else {
				fmt.Fprintf(out, "_ ")
			}
		}
		fmt.Fprintf(out, "\n")
	}
}

func (bd Board) PrintAsInput(out io.Writer) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if bd[i][j].Solved {
				fmt.Fprintf(out, "%1d", bd[i][j].Value)
			} else {
				fmt.Fprintf(out, ".")
			}
		}
		fmt.Fprintf(out, "\n")
	}
}

func ReadBoard(in io.Reader) Board {
	bd := New()
	r := bufio.NewReader(in)
	for row := 0; row < 9; {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				if row != 9 {
					log.Fatalf("Premature end-of-file at row %d\n", row)
				}
				break
			}
			log.Fatal(err)
		}
		if buf[0] == '#' {
			continue
		}
		buf = bytes.Trim(buf, " \t\n\r")

		col := 0
		for _, c := range buf {
			if c == ',' || c == ' ' {
				continue
			}
			n := int(c - '0')
			if c == '_' || c == '.' {
				n = 0
			}
			if n < 0 || n > 10 {
				// Will this ever happen?
				log.Fatalf("Numbers must be less than 10, greater than zero: %d (%c)\n", n, c)
			}
			if n != 0 {
				bd.MarkSolved(row, col, n)
			}
			col++
		}
		if col != 9 {
			log.Fatalf("Row %d had %d cols\n", row+1, col)
		}
		row++
	}

	// Based on the digits given in the input,
	// eliminate those digits in the row, column
	// and block they appear in.
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if bd[row][col].Solved {
				bd.EliminatePossibilities(row, col, bd[row][col].Block, bd[row][col].Value)
			}
		}
	}
	return bd
}
