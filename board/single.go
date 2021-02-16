package board

import "fmt"

func (bd *Board) FindSingles(announce bool) int {
	count := 0
	for rowNo := 0; rowNo < 9; rowNo++ {
		for colNo := 0; colNo < 9; colNo++ {
			if bd[rowNo][colNo].Solved {
				continue
			}
			if len(bd[rowNo][colNo].Possible) == 1 {
				count++
				bd.MarkSolved(rowNo, colNo, bd[rowNo][colNo].Possible[0])
				bd.EliminatePossibilities(rowNo, colNo, bd[rowNo][colNo].Block, bd[rowNo][colNo].Value)
				if announce {
					fmt.Printf("<%d,%d> solved with only digit %d\n", rowNo, colNo, bd[rowNo][colNo].Value)
				}
			}
		}
	}
	return count
}
