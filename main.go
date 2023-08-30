package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sudoku2/board"
)

func main() {
	var validateOnly bool
	var printPossible bool
	var printPossiblePS bool
	var printIntermediate bool
	var printAsInput bool
	var announceSolutions bool
	var nakedPairElimination bool
	var hiddenPairElimination bool
	var hiddenTripletsElimination bool
	var nakedTripletsElimination bool
	var solveByBackTracking bool
	var backTrackingOnly bool
	var pointingElimination bool
	var xwingElimination bool
	var finalOutput bool
	flag.BoolVar(&finalOutput, "f", false, "print final board in PostScript")
	flag.BoolVar(&validateOnly, "v", false, "validate input board only")
	flag.BoolVar(&printPossible, "c", false, "on incomplete solution, print digit possibilities")
	flag.BoolVar(&printPossiblePS, "C", false, "print digit possibilities in PostScript output")
	flag.BoolVar(&printIntermediate, "i", false, "print intermediate solved boards")
	flag.BoolVar(&printAsInput, "q", false, "print final board in input format")
	flag.BoolVar(&announceSolutions, "a", false, "announce solution digits")
	flag.BoolVar(&nakedPairElimination, "N", false, "perform naked pair elimination")
	flag.BoolVar(&hiddenPairElimination, "H", false, "perform hidden pair elimination")
	flag.BoolVar(&hiddenTripletsElimination, "T", false, "perform hidden triplets elimination")
	flag.BoolVar(&nakedTripletsElimination, "t", false, "perform naked triplets elimination")
	flag.BoolVar(&pointingElimination, "P", false, "perform block pointing elimination")
	flag.BoolVar(&xwingElimination, "X", false, "perform Xwing elimination")
	flag.BoolVar(&solveByBackTracking, "B", false, "solve by backtracking, if necessary")
	flag.BoolVar(&backTrackingOnly, "b", false, "solve by backtracking only, no other eliminations")
	postScriptFileName := flag.String("p", "", "PostScript output in this file")
	phrase := flag.String("y", "", "phrase to print in PostScript")
	flag.Parse()

	bd, err := board.FigureOutBoard()
	if err != nil {
		log.Fatal(err)
	}

	if validateOnly {
		if !bd.Valid() {
			return
		}
		fmt.Printf("Valid input board\n")
		return
	}

	if !finalOutput && *postScriptFileName != "" {
		fout, err := os.Create(*postScriptFileName)
		if err != nil {
			log.Fatal(err)
		}
		if *phrase == "" {
			*phrase = *postScriptFileName
		}
		bd.EmitPostScript(fout, *phrase, printPossiblePS)
		return
	}

	bd.Print(os.Stdout)

	if printPossible {
		for row := 0; row < 9; row++ {
			for col := 0; col < 9; col++ {
				if bd[row][col].Solved {
					fmt.Printf("<%d,%d> %d %v\n", row, col, bd[row][col].Block, bd[row][col].Value)
				} else {
					fmt.Printf("<%d,%d> %d %v\n", row, col, bd[row][col].Block, bd[row][col].Possible)
				}
			}
		}
	}

	if !bd.Valid() {
		return
	}

	n := 1
	m := 0
	count := 0
	for !backTrackingOnly && n > 0 && count < 81 {
		count++
		fmt.Printf("-- Iteration %d\n", count)
		n = 0

		m = bd.FindSingles(announceSolutions)
		fmt.Printf("Found %d single digits\n", m)
		n += m

		m = bd.FindIsolates(announceSolutions)
		fmt.Printf("Found %d isolated digits\n", m)
		n += m

		if printIntermediate {
			fmt.Println("===intermediate===")
			if printAsInput {
				bd.PrintAsInput(os.Stdout)
			} else {
				bd.Print(os.Stdout)
			}
			fmt.Println("==================")
		}

		if nakedPairElimination {
			m = bd.NakedPairEliminate(announceSolutions)
			fmt.Printf("Eliminated %d candidates via naked pairs\n", m)
			n += m
		}

		if hiddenPairElimination {
			m = bd.HiddenPairEliminate(announceSolutions)
			fmt.Printf("Eliminated %d candidates via hidden pairs\n", m)
			n += m
		}

		if hiddenTripletsElimination {
			m = bd.HiddenTripletsEliminate(announceSolutions)
			fmt.Printf("Eliminated %d candidates via hidden triples\n", m)
			n += m
		}

		if nakedTripletsElimination {
			m = bd.NakedTripletsEliminate(announceSolutions)
			fmt.Printf("Eliminated %d candidates via naked triples\n", m)
			n += m
		}

		if n == 0 && pointingElimination {
			m = bd.PointingElimination(announceSolutions)
			fmt.Printf("Eliminated %d candidates via pointing\n", m)
			n += m
		}

		if n == 0 && xwingElimination {
			m = bd.XwingEliminate(announceSolutions)
			fmt.Printf("Eliminated %d candidates via xwing\n", m)
			n += m
		}

		if !bd.Valid() {
			break
		}
	}

	if (solveByBackTracking || backTrackingOnly) && !bd.Finished() {
		fmt.Println("Solving via backtracking")
		board.BackTrackSolution(bd)
	}

	if finalOutput && *postScriptFileName != "" {
		fout, err := os.Create(*postScriptFileName)
		if err != nil {
			log.Fatal(err)
		}
		bd.EmitPostScript(fout, *postScriptFileName, printPossiblePS)
		return
	}

	if printAsInput {
		bd.PrintAsInput(os.Stdout)
		return
	}
	bd.Print(os.Stdout)
}
