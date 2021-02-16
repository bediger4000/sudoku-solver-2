package board

import "fmt"

func (bd *Board) FindIsolates(announce bool) int {
	found := 0

	for rowNo := 0; rowNo < 9; rowNo++ {
		digitCount := make(map[int]int)

		for _, cell := range bd[rowNo] {
			if cell.Solved {
				continue
			}
			for _, possibleDigit := range cell.Possible {
				digitCount[possibleDigit]++
			}
		}

		for digit, count := range digitCount {
			if count == 1 {
				found++
				for colNo := 0; colNo < 9; colNo++ {
					if bd[rowNo][colNo].Solved {
						continue
					}
					for _, possibleDigit := range bd[rowNo][colNo].Possible {
						if possibleDigit != digit {
							continue
						}
						bd.MarkSolved(rowNo, colNo, digit)
						bd.EliminatePossibilities(rowNo, colNo, bd[rowNo][colNo].Block, bd[rowNo][colNo].Value)
						if announce {
							fmt.Printf("<%d,%d> solved with only possible digit %d in row %d\n", rowNo, colNo, bd[rowNo][colNo].Value, rowNo)
						}
					}
				}
			}
		}
	}

	for colNo := 0; colNo < 9; colNo++ {

		digitCount := make(map[int]int)

		for rowNo := 0; rowNo < 9; rowNo++ {
			if bd[rowNo][colNo].Solved {
				continue
			}
			for _, possibleDigit := range bd[rowNo][colNo].Possible {
				digitCount[possibleDigit]++
			}
		}

		for digit, count := range digitCount {
			if count != 1 {
				continue
			}
			found++
			for rowNo := 0; rowNo < 9; rowNo++ {
				if bd[rowNo][colNo].Solved {
					continue
				}
				for _, possibleDigit := range bd[rowNo][colNo].Possible {
					if possibleDigit != digit {
						continue
					}
					bd.MarkSolved(rowNo, colNo, digit)
					bd.EliminatePossibilities(rowNo, colNo, bd[rowNo][colNo].Block, bd[rowNo][colNo].Value)
					if announce {
						fmt.Printf("<%d,%d> solved with only possible digit %d in col %d\n", rowNo, colNo, bd[rowNo][colNo].Value, colNo)
					}
				}
			}
		}
	}

	for blockNo := 0; blockNo < 9; blockNo++ {
		digitCount := make(map[int]int)
		for colNo := 0; colNo < 9; colNo++ {
			for rowNo := 0; rowNo < 9; rowNo++ {
				if bd[rowNo][colNo].Solved {
					continue
				}
				if bd[rowNo][colNo].Block == blockNo {
					for _, possibleDigit := range bd[rowNo][colNo].Possible {
						digitCount[possibleDigit]++
					}
				}
			}
		}
		for digit, count := range digitCount {
			if count == 1 {
				found++
				for colNo := 0; colNo < 9; colNo++ {
					for rowNo := 0; rowNo < 9; rowNo++ {
						if bd[rowNo][colNo].Solved {
							continue
						}
						if bd[rowNo][colNo].Block == blockNo {
							for _, possibleDigit := range bd[rowNo][colNo].Possible {
								if possibleDigit != digit {
									continue
								}
								bd.MarkSolved(rowNo, colNo, digit)
								bd.EliminatePossibilities(rowNo, colNo, bd[rowNo][colNo].Block, bd[rowNo][colNo].Value)
								if announce {
									fmt.Printf("<%d,%d> solved with only possible digit %d in block %d\n", rowNo, colNo, bd[rowNo][colNo].Value, blockNo)
								}
							}
						}
					}
				}
			}
		}
	}

	return found
}

func (bd *Board) MarkSolved(rowNo, colNo, digit int) {
	bd[rowNo][colNo].Solved = true
	bd[rowNo][colNo].Value = digit
	bd[rowNo][colNo].Possible = []int{}
}
