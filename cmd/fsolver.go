package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sudoku2/board"
)

func main() {

	var announceSolutions bool
	var nakedPairElimination bool
	var hiddenPairElimination bool
	var hiddenTripletsElimination bool
	var nakedTripletsElimination bool
	var pointingElimination bool
	var xwingElimination bool
	var verifyAndCompare bool

	flag.BoolVar(&announceSolutions, "a", false, "announce solution digits")
	flag.BoolVar(&nakedPairElimination, "N", false, "perform naked pair elimination")
	flag.BoolVar(&hiddenPairElimination, "H", false, "perform hidden pair elimination")
	flag.BoolVar(&hiddenTripletsElimination, "T", false, "perform hidden triplets elimination")
	flag.BoolVar(&nakedTripletsElimination, "t", false, "perform naked triplets elimination")
	flag.BoolVar(&pointingElimination, "P", false, "perform block pointing elimination")
	flag.BoolVar(&xwingElimination, "X", false, "perform Xwing elimination")
	flag.BoolVar(&verifyAndCompare, "v", false, "verify and compare solutions")

	flag.Parse()

	fin := os.Stdin
	if flag.NArg() > 0 {
		var err error
		if fin, err = os.Open(flag.Arg(0)); err != nil {
			log.Fatal(err)
		}
		defer fin.Close()
	}

	scanner := bufio.NewScanner(fin)
	lineCounter := 0

	stumpedCount, solvedCount, mismatchCount := 0, 0, 0

	for scanner.Scan() {
		lineCounter++
		line := scanner.Text()

		bd, err := board.NewBoardFromString(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line %d: %v\n", err)
			continue
		}

		if !bd.Valid() {
			fmt.Fprintf(os.Stderr, "line %d invalid board\n", lineCounter)
			continue
		}

		var bdCopy *board.Board
		if verifyAndCompare {
			bdCopy = bd.Copy()
		}

		// read in a board, try to solve it
		for iteration := 0; true; {
			iteration++

			m := bd.FindSingles(announceSolutions)
			if announceSolutions {
				fmt.Printf("found %d singles\n", m)
			}

			n := bd.FindIsolates(announceSolutions)
			if announceSolutions {
				fmt.Printf("found %d isolates\n", n)
			}
			m += n

			n = bd.NakedPairEliminate(announceSolutions)
			if announceSolutions {
				fmt.Printf("eliminated %d candidates via naked pair\n", n)
			}
			m += n

			n = bd.HiddenPairEliminate(announceSolutions)
			if announceSolutions {
				fmt.Printf("eliminated %d candidates via hidden pair\n", n)
			}
			m += n

			n = bd.HiddenTripletsEliminate(announceSolutions)
			if announceSolutions {
				fmt.Printf("eliminated %d candidates via hidden triplets\n", n)
			}
			m += n

			n = bd.NakedTripletsEliminate(announceSolutions)
			if announceSolutions {
				fmt.Printf("eliminated %d candidates via naked triplets\n", n)
			}
			m += n

			n = bd.PointingElimination(announceSolutions)
			if announceSolutions {
				fmt.Printf("eliminated %d candidates via pointing\n", n)
			}
			m += n

			n = bd.XwingEliminate(announceSolutions)
			if announceSolutions {
				fmt.Printf("eliminated %d candidates via XWing\n", n)
			}
			m += n

			valid, complete := bd.ValidAndComplete()

			if !valid {
				fmt.Fprintf(os.Stderr, "Invalid board, line %d, iteration %d:\n", lineCounter, iteration)
				bd.Print(os.Stderr)
				break
			}

			if complete {
				solvedCount++
				fmt.Printf("Line %d, iteration %d solved:\n", lineCounter, iteration)
				if announceSolutions {
					bd.Print(os.Stdout)
				}
				if verifyAndCompare {
					solvedBd := board.BackTrackSolved(bdCopy)
					if !board.CompareSolutions(solvedBd, bd, true) {
						mismatchCount++
						fmt.Printf("Backtracking soluiont:\n")
						solvedBd.Print(os.Stdout)
						fmt.Printf("Standard soluiont:\n")
						bd.Print(os.Stdout)
					}
				}
				break
			}

			if m == 0 {
				stumpedCount++
				fmt.Printf("Line %d, iteration %d stumped:\n", lineCounter, iteration)
				bd.Print(os.Stdout)
				break
			}
		}
	}

	fmt.Printf("%d input puzzles, %d solved (%d didn't verify), %d stumped\n",
		lineCounter, solvedCount, mismatchCount, stumpedCount,
	)

	if err := scanner.Err(); err != nil {
		log.Fatalf("problem line %d: %v", lineCounter, err)
	}

}
