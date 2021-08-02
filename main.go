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
	var solveByBackTracking bool
	var pointingElimination bool
	var finalOutput bool
	flag.BoolVar(&finalOutput, "f", false, "print final board in PostScript")
	flag.BoolVar(&validateOnly, "v", false, "validate input board only")
	flag.BoolVar(&printPossible, "c", false, "on incomplete solution, print digit possibilities")
	flag.BoolVar(&printPossiblePS, "C", false, "print digit possibilities in PostScript output")
	flag.BoolVar(&printIntermediate, "i", false, "print intermediate solved boards")
	flag.BoolVar(&printAsInput, "t", false, "print final board in input format")
	flag.BoolVar(&announceSolutions, "a", false, "announce solution digits")
	flag.BoolVar(&nakedPairElimination, "N", false, "perform naked pair elimination")
	flag.BoolVar(&hiddenPairElimination, "H", false, "perform hidden pair elimination")
	flag.BoolVar(&pointingElimination, "P", false, "perform block pointing elimination")
	flag.BoolVar(&solveByBackTracking, "B", false, "solve by backtracking, if necessary")
	postScriptFileName := flag.String("p", "", "PostScript output in this file")
	flag.Parse()

	fin := os.Stdin
	if flag.NArg() > 0 {
		var err error
		fin, err = os.Open(flag.Arg(0))
		defer fin.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
	bd := board.ReadBoard(fin)

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
		(&bd).EmitPostScript(fout, *postScriptFileName, printPossiblePS)
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
	for n > 0 && count < 81 {
		count++
		n = 0

		m = (&bd).FindSingles(announceSolutions)
		fmt.Printf("Found %d single digits\n", m)
		n += m

		m = (&bd).FindIsolates(announceSolutions)
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
			m = (&bd).NakedPairEliminate(announceSolutions)
			fmt.Printf("Eliminated %d candidates via naked pairs\n", m)
			n += m
		}

		if hiddenPairElimination {
			m = (&bd).HiddenPairEliminate(announceSolutions)
			fmt.Printf("Eliminated %d candidates via hidden pairs\n", m)
			n += m
		}

		if n == 0 && pointingElimination {
			fmt.Println("--- before pointing ---")
			bd.PrintAsInput(os.Stdout)
			fmt.Println("--- before pointing ---")
			m = (&bd).PointingElimination(announceSolutions)
			fmt.Printf("Eliminated %d candidates via pointing\n", m)
			n += m
			fmt.Println("--- after pointing ---")
			bd.PrintAsInput(os.Stdout)
			fmt.Println("--- after pointing ---")
		}

		if !(&bd).Valid() {
			break
		}
	}

	if solveByBackTracking && !bd.Finished() {
		fmt.Println("Solving via backtracking")
		board.BackTrackSolution(&bd)
	}

	if finalOutput && *postScriptFileName != "" {
		fout, err := os.Create(*postScriptFileName)
		if err != nil {
			log.Fatal(err)
		}
		(&bd).EmitPostScript(fout, *postScriptFileName, printPossiblePS)
		return
	}

	if printAsInput {
		bd.PrintAsInput(os.Stdout)
		return
	}
	bd.Print(os.Stdout)
}
