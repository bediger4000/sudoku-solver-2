package main

import (
	"fmt"
	"io"
	"os"
	"sudoku2/board"
)

func main() {
	bd := board.ReadBoard(os.Stdin)
	identity(bd, os.Stdout)
	reflectvertical(bd, os.Stdout)
	reflecthorizontal(bd, os.Stdout)
	reflecthorizontal(bd, os.Stdout)
	halfturn(bd, os.Stdout)
	transpose(bd, os.Stdout)
	quarterturnclockwiise(bd, os.Stdout)
	quarterturncounterclockwiise(bd, os.Stdout)
	diagnoalflip(bd, os.Stdout)
}

// Reflect around vertical axis
func reflectvertical(bd *board.Board, out io.Writer) {
	fmt.Fprintf(out, "=== Reflect around vertical axis\n")
	for r := 0; r < 9; r++ {
		for c := 8; c >= 0; c-- {
			if bd[r][c].Solved {
				fmt.Fprintf(out, "%d ", bd[r][c].Value)
			} else {
				fmt.Fprintf(out, "_ ")
			}
		}
		fmt.Fprintf(out, "\n")
	}
}

func reflecthorizontal(bd *board.Board, out io.Writer) {
	fmt.Fprintf(out, "=== Reflect around horizontal axis\n")
	for r := 8; r >= 0; r-- {
		for c := 0; c < 9; c++ {
			if bd[r][c].Solved {
				fmt.Fprintf(out, "%d ", bd[r][c].Value)
			} else {
				fmt.Fprintf(out, "_ ")
			}
		}
		fmt.Fprintf(out, "\n")
	}
}

func halfturn(bd *board.Board, out io.Writer) {
	fmt.Fprintf(out, "=== Half turn in plane\n")
	for r := 8; r >= 0; r-- {
		for c := 8; c >= 0; c-- {
			if bd[r][c].Solved {
				fmt.Fprintf(out, "%d ", bd[r][c].Value)
			} else {
				fmt.Fprintf(out, "_ ")
			}
		}
		fmt.Fprintf(out, "\n")
	}
}

func identity(bd *board.Board, out io.Writer) {
	fmt.Fprintf(out, "=== Identity\n")
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if bd[r][c].Solved {
				fmt.Fprintf(out, "%d ", bd[r][c].Value)
			} else {
				fmt.Fprintf(out, "_ ")
			}
		}
		fmt.Fprintf(out, "\n")
	}
}

func transpose(bd *board.Board, out io.Writer) {
	fmt.Fprintf(out, "=== Transpose\n")
	for c := 0; c < 9; c++ {
		for r := 0; r < 9; r++ {
			if bd[r][c].Solved {
				fmt.Fprintf(out, "%d ", bd[r][c].Value)
			} else {
				fmt.Fprintf(out, "_ ")
			}
		}
		fmt.Fprintf(out, "\n")
	}
}

func quarterturnclockwiise(bd *board.Board, out io.Writer) {
	fmt.Fprintf(out, "=== Quarter turn clockwise\n")
	for c := 0; c < 9; c++ {
		for r := 8; r >= 0; r-- {
			if bd[r][c].Solved {
				fmt.Fprintf(out, "%d ", bd[r][c].Value)
			} else {
				fmt.Fprintf(out, "_ ")
			}
		}
		fmt.Fprintf(out, "\n")
	}
}

func quarterturncounterclockwiise(bd *board.Board, out io.Writer) {
	fmt.Fprintf(out, "=== Quarter turn counterclockwise\n")
	for c := 8; c >= 0; c-- {
		for r := 0; r < 9; r++ {
			if bd[r][c].Solved {
				fmt.Fprintf(out, "%d ", bd[r][c].Value)
			} else {
				fmt.Fprintf(out, "_ ")
			}
		}
		fmt.Fprintf(out, "\n")
	}
}

func diagnoalflip(bd *board.Board, out io.Writer) {
	fmt.Fprintf(out, "=== flip diagonally\n")
	for c := 8; c >= 0; c-- {
		for r := 8; r >= 0; r-- {
			if bd[r][c].Solved {
				fmt.Fprintf(out, "%d ", bd[r][c].Value)
			} else {
				fmt.Fprintf(out, "_ ")
			}
		}
		fmt.Fprintf(out, "\n")
	}
}
