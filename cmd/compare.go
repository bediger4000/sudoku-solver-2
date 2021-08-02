package main

import (
	"fmt"
	"log"
	"os"
	"sudoku2/board"
)

func main() {

	bd1 := readboard(os.Args[1])
	bd2 := readboard(os.Args[2])

	if completes(bd1, bd2) {
		fmt.Printf("%s solves %s\n", os.Args[2], os.Args[1])
	} else {
		fmt.Printf("%s does not solve %s\n", os.Args[2], os.Args[1])
	}
	return
}

func completes(start, solution board.Board) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if start[i][j].Solved {
				if start[i][j].Value != solution[i][j].Value {
					return false
				}
			}
		}
	}
	return true
}

func readboard(fname string) board.Board {
	f1, err := os.Open(fname)
	defer f1.Close()
	if err != nil {
		log.Fatal(err)
	}
	return board.ReadBoard(f1)
}
