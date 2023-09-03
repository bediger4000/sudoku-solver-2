package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sudoku2/board"
)

func main() {

	announceSolutions := false

	fin := os.Stdin
	if len(os.Args) > 1 {
		var err error
		if fin, err = os.Open(os.Args[1]); err != nil {
			log.Fatal(err)
		}
		defer fin.Close()
	}

	scanner := bufio.NewScanner(fin)
	lineCounter := 0

	for scanner.Scan() {
		lineCounter++
		line := scanner.Text()

		bd, err := board.NewBoardFromString(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line %d: %v\n", err)
			continue
		}
		for row := 0; row < 9; row++ {
			for col := 0; col < 9; col++ {
				if bd[row][col].Solved {
					bd.EliminatePossibilities(row, col, bd[row][col].Block, bd[row][col].Value)
				}
			}
		}
		// bd.Print(os.Stdout)
		// bd.Details(os.Stdout)
		if !bd.Valid() {
			fmt.Fprintf(os.Stderr, "line %d invalid board\n", lineCounter)
			continue
		}

		// read in a board, try to solve it
		for iteration := 0; true; {
			iteration++
			//	fmt.Printf("-- Iteration %d\n", iteration)

			m := bd.FindSingles(announceSolutions)
			//	fmt.Printf("found %d singles\n", m)

			n := bd.FindIsolates(announceSolutions)
			//	fmt.Printf("found %d isolates\n", n)
			m += n

			n = bd.NakedPairEliminate(announceSolutions)
			//	fmt.Printf("eliminated %d candidates via naked pair\n", n)
			m += n

			n = bd.HiddenPairEliminate(announceSolutions)
			//	fmt.Printf("eliminated %d candidates via hidden pair\n", n)
			m += n

			n = bd.HiddenTripletsEliminate(announceSolutions)
			//	fmt.Printf("eliminated %d candidates via hidden triplets\n", n)
			m += n

			n = bd.NakedTripletsEliminate(announceSolutions)
			//	fmt.Printf("eliminated %d candidates via naked triplets\n", n)
			m += n

			n = bd.PointingElimination(announceSolutions)
			//	fmt.Printf("eliminated %d candidates via pointing\n", n)
			m += n

			n = bd.XwingEliminate(announceSolutions)
			//	fmt.Printf("eliminated %d candidates via XWing\n", n)
			m += n

			valid, complete := bd.ValidAndComplete()

			if !valid {
				fmt.Fprintf(os.Stderr, "Invalid board, line %d, iteration %d:\n", lineCounter, iteration)
				bd.Print(os.Stderr)
				break
			}

			if complete {
				fmt.Printf("Line %d, iteration %d solved:\n", lineCounter, iteration)
				//	bd.Print(os.Stdout)
				break
			}

			if m == 0 {
				fmt.Printf("Line %d, iteration %d stumped:\n", lineCounter, iteration)
				bd.Print(os.Stdout)
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("problem line %d: %v", lineCounter, err)
	}

}
