package board

import "fmt"

func (bd *Board) NakedPairEliminate(announce bool) int {
	candidatesEliminated := bd.RowNakedPairs(announce)
	candidatesEliminated += bd.ColNakedPairs(announce)
	candidatesEliminated += bd.BlockNakedPairs(announce)
	return candidatesEliminated
}

// all the info needed to do matching pair elimination
type pairMarker struct {
	row int
	col int
	a   int
	b   int
	rep int
}

func (bd *Board) RowNakedPairs(announce bool) int {
	totalEliminated := 0

	for rowNo := 0; rowNo < 9; rowNo++ {

		// find cells in row with only 2 possibilities
		var cellsWith2 []*pairMarker
		for colNo := 0; colNo < 9; colNo++ {
			if len(bd[rowNo][colNo].Possible) == 2 {
				rep := 10*bd[rowNo][colNo].Possible[0] + bd[rowNo][colNo].Possible[1]
				cellsWith2 = append(cellsWith2, &pairMarker{row: rowNo, col: colNo, a: bd[rowNo][colNo].Possible[0], b: bd[rowNo][colNo].Possible[1], rep: rep})
			}
		}

		// find any sets of 2 cells with matching pairs
		// pairMarker.rep is 10*a + b, so identical reps match

		setsMatching := make(map[int][]*pairMarker)
		for _, cell := range cellsWith2 {
			setsMatching[cell.rep] = append(setsMatching[cell.rep], cell)
		}

		for _, cellArray := range setsMatching {
			if len(cellArray) == 2 {
				if announce {
					fmt.Printf("found naked pair (%d,%d) in row %d, cols %d and %d\n",
						cellArray[0].a, cellArray[0].b, rowNo, cellArray[0].col, cellArray[1].col)
				}

				// found 2 cells in row rowNo with matching pairs,
				// eliminate pairMarker.a and pairMarker.b from rest of row
				a, b := cellArray[0].a, cellArray[0].b
				for colNo := 0; colNo < 9; colNo++ {
					if (colNo == cellArray[0].col) && (rowNo == cellArray[0].row) {
						continue
					}
					if (colNo == cellArray[1].col) && (rowNo == cellArray[1].row) {
						continue
					}

					spliced := bd.SpliceOut(rowNo, colNo, a)

					if announce && spliced == 1 {
						fmt.Printf("\teliminated %d at <%d,%d>\n", a, rowNo, colNo)
					}

					totalEliminated += spliced

					spliced = bd.SpliceOut(rowNo, colNo, cellArray[0].b)

					if announce && spliced == 1 {
						fmt.Printf("\teliminated %d at <%d,%d>\n", b, rowNo, colNo)
					}

					totalEliminated += spliced

				}
			}
		}
	}
	return totalEliminated
}

func (bd *Board) ColNakedPairs(announce bool) int {
	totalEliminated := 0

	for colNo := 0; colNo < 9; colNo++ {

		// find cells in row with only 2 possibilities
		var cellsWith2 []*pairMarker
		for rowNo := 0; rowNo < 9; rowNo++ {
			if len(bd[rowNo][colNo].Possible) == 2 {
				rep := 10*bd[rowNo][colNo].Possible[0] + bd[rowNo][colNo].Possible[1]
				cellsWith2 = append(cellsWith2, &pairMarker{row: rowNo, col: colNo, a: bd[rowNo][colNo].Possible[0], b: bd[rowNo][colNo].Possible[1], rep: rep})
			}
		}

		// find any sets of 2 cells with matching pairs
		// pairMarker.rep is 10*a + b, so identical reps match

		setsMatching := make(map[int][]*pairMarker)
		for _, cell := range cellsWith2 {
			setsMatching[cell.rep] = append(setsMatching[cell.rep], cell)
		}

		for _, cellArray := range setsMatching {
			if len(cellArray) == 2 {
				if announce {
					fmt.Printf("found naked pair (%d,%d) in col %d, cols %d and %d\n",
						cellArray[0].a, cellArray[0].b, colNo, cellArray[0].col, cellArray[1].col)
				}

				// found 2 cells in column colNo with matching pairs,
				// eliminate pairMarker.a and pairMarker.b from rest of column

				for rowNo := 0; rowNo < 9; rowNo++ {
					if (colNo == cellArray[0].col) && (rowNo == cellArray[0].row) {
						continue
					}
					if (colNo == cellArray[1].col) && (rowNo == cellArray[1].row) {
						continue
					}

					a, b := cellArray[0].a, cellArray[0].b

					spliced := bd.SpliceOut(rowNo, colNo, a)

					if announce && spliced == 1 {
						fmt.Printf("\teliminated %d at <%d,%d>\n", a, rowNo, colNo)
					}

					totalEliminated += spliced

					spliced = bd.SpliceOut(rowNo, colNo, cellArray[0].b)
					if announce && spliced == 1 {
						fmt.Printf("\teliminated %d at <%d,%d>\n", b, rowNo, colNo)
					}

					totalEliminated += spliced
				}
			}
		}
	}
	return totalEliminated
}

func (bd *Board) BlockNakedPairs(announce bool) int {
	totalEliminated := 0

	for blockNo := 0; blockNo < 9; blockNo++ {
		// find cells in block blockNo with only 2 possibilities
		var cellsWith2 []*pairMarker
		for colNo := 0; colNo < 9; colNo++ {
			for rowNo := 0; rowNo < 9; rowNo++ {
				if bd[rowNo][colNo].Block != blockNo {
					continue
				}
				if len(bd[rowNo][colNo].Possible) == 2 {
					rep := 10*bd[rowNo][colNo].Possible[0] + bd[rowNo][colNo].Possible[1]
					cellsWith2 = append(cellsWith2, &pairMarker{row: rowNo, col: colNo, a: bd[rowNo][colNo].Possible[0], b: bd[rowNo][colNo].Possible[1], rep: rep})
				}
			}
		}

		// find any sets of 2 cells with matching pairs in this block
		// pairMarker.rep is 10*a + b, so identical reps match

		setsMatching := make(map[int][]*pairMarker)
		for _, cell := range cellsWith2 {
			setsMatching[cell.rep] = append(setsMatching[cell.rep], cell)
		}

		for _, cellArray := range setsMatching {
			if len(cellArray) == 2 {
				// found 2 cells in block blockNo with matching pairs,
				// eliminate pairMarker.a and pairMarker.b from rest of block
				if announce {
					fmt.Printf("found naked pair (%d,%d) in block %d, at <%d,%d> and <%d,%d>\n",
						cellArray[0].a, cellArray[0].b, blockNo,
						cellArray[0].row, cellArray[0].col, cellArray[1].row, cellArray[1].col)
				}

				for rowNo := 0; rowNo < 9; rowNo++ {
					for colNo := 0; colNo < 9; colNo++ {
						if bd[rowNo][colNo].Block != blockNo {
							continue
						}

						if (colNo == cellArray[0].col) && (rowNo == cellArray[0].row) {
							continue
						}
						if (colNo == cellArray[1].col) && (rowNo == cellArray[1].row) {
							continue
						}

						a, b := cellArray[0].a, cellArray[0].b

						spliced := bd.SpliceOut(rowNo, colNo, a)

						if announce && spliced > 0 {
							fmt.Printf("\teliminated %d at <%d,%d>\n", a, rowNo, colNo)
						}

						totalEliminated += spliced

						spliced = bd.SpliceOut(rowNo, colNo, b)

						if announce && spliced > 0 {
							fmt.Printf("\teliminated %d at <%d,%d>\n", b, rowNo, colNo)
						}

						totalEliminated += spliced
					}
				}

				// found 2 cells in block blockNo with matching pairs,
				// if they are in a row, or a column,
				// eliminate pairMarker.a and pairMarker.b from rest of that col or row
				if cellArray[0].row == cellArray[1].row {
					if announce {
						fmt.Printf("found naked pair (%d,%d) in block %d, and row %d\n",
							cellArray[0].a, cellArray[0].b, blockNo, cellArray[0].row)
					}
				}
				if cellArray[0].col == cellArray[1].col {
					if announce {
						fmt.Printf("found naked pair (%d,%d) in block %d, and col %d\n",
							cellArray[0].a, cellArray[0].b, blockNo, cellArray[0].col)
					}
				}
			}
		}
	}
	return totalEliminated
}

func (bd *Board) HiddenPairEliminate(announce bool) int {
	candidatesEliminated := bd.RowHiddenPairs(announce)
	candidatesEliminated += bd.ColHiddenPairs(announce)
	candidatesEliminated += bd.BlockHiddenPairs(announce)
	return candidatesEliminated
}

func (bd *Board) RowHiddenPairs(announce bool) int {
	eliminated := 0
	for rowNo := 0; rowNo < 9; rowNo++ {
		digitsCount := bd.digitsInRow(rowNo)
		var twoCount []int
		for digit, count := range digitsCount {
			if count == 2 {
				twoCount = append(twoCount, digit)
			}
		}
		if len(twoCount) < 2 {
			continue
		}

		// find cells that contain 2 of the digits in twoCount
		hiddenPairPairs := bd.findPairsInRow(rowNo, twoCount)

		for _, pair := range hiddenPairPairs {
			a, b := pair[0].a, pair[0].b
			if len(bd[pair[0].row][pair[0].col].Possible) > 2 ||
				len(bd[pair[1].row][pair[1].col].Possible) > 2 {
				if announce {
					fmt.Printf("Row %d Hidden pair [%d,%d]/[%d,%d] at <%d,%d> and <%d,%d>\n",
						rowNo,
						pair[0].a, pair[0].b, pair[1].a, pair[1].b,
						pair[0].row, pair[0].col, pair[1].row, pair[1].col,
					)
				}
				// eliminate all digits but a, b in both pairs
				for i := range pair {
					row, col := pair[i].row, pair[i].col
					if announce {
						fmt.Printf("\tRow %d, <%d,%d>, eliminating all but %d and %d from %v\n",
							rowNo, pair[i].row, pair[i].col,
							a, b,
							bd[row][col].Possible,
						)
					}
					oldCount := len(bd[row][col].Possible)
					bd[row][col].Possible = []int{a, b}
					eliminated += (oldCount - 2)
				}
			}
		}
	}

	return eliminated
}

func (bd *Board) findPairsInRow(rowNo int, pairDigits []int) [][2]*pairMarker {
	var r [][2]*pairMarker

	pairs := make(map[int][]*pairMarker)
	for _, a := range pairDigits {
		for _, b := range pairDigits {
			if b <= a {
				continue
			}
			// find pairs of cells in rowNo with a and b in Possible slice
			for colNo := 0; colNo < 9; colNo++ {
				for _, digit1 := range bd[rowNo][colNo].Possible {
					if digit1 == a {
						for _, digit2 := range bd[rowNo][colNo].Possible {
							if digit2 == b {
								rep := 10*a + b
								pairs[rep] = append(pairs[rep], &pairMarker{a: a, b: b, row: rowNo, col: colNo, rep: rep})
							}
						}
					}
				}
			}
		}
	}

	if len(pairs) > 0 {
		for _, pary := range pairs {
			if len(pary) == 2 {
				r = append(r, [2]*pairMarker{pary[0], pary[1]})
			}
		}
	}

	return r
}

func (bd *Board) findPairsInBlock(blockNo int, pairDigits []int) [][2]*pairMarker {
	var r [][2]*pairMarker

	pairs := make(map[int][]*pairMarker)
	for _, a := range pairDigits {
		for _, b := range pairDigits {
			if b <= a {
				continue
			}
			// find pairs of cells in blockNo with a and b in Possible slice
			for rowNo := 0; rowNo < 9; rowNo++ {
				for colNo := 0; colNo < 9; colNo++ {
					if bd[rowNo][colNo].Block != blockNo {
						continue
					}
					for _, digit1 := range bd[rowNo][colNo].Possible {
						if digit1 == a {
							for _, digit2 := range bd[rowNo][colNo].Possible {
								if digit2 == b {
									rep := 10*a + b
									pairs[rep] = append(pairs[rep], &pairMarker{a: a, b: b, row: rowNo, col: colNo, rep: rep})
								}
							}
						}
					}
				}
			}
		}
	}

	if len(pairs) > 0 {
		for _, pary := range pairs {
			if len(pary) == 2 {
				r = append(r, [2]*pairMarker{pary[0], pary[1]})
			}
		}
	}

	return r
}

func (bd *Board) findPairsInCol(colNo int, pairDigits []int) [][2]*pairMarker {
	var r [][2]*pairMarker

	pairs := make(map[int][]*pairMarker)
	for _, a := range pairDigits {
		for _, b := range pairDigits {
			if b <= a {
				continue
			}
			// find pairs of cells in rowNo with a and b in Possible slice
			for rowNo := 0; rowNo < 9; rowNo++ {
				for _, digit1 := range bd[rowNo][colNo].Possible {
					if digit1 == a {
						for _, digit2 := range bd[rowNo][colNo].Possible {
							if digit2 == b {
								rep := 10*a + b
								pairs[rep] = append(pairs[rep], &pairMarker{a: a, b: b, row: rowNo, col: colNo, rep: rep})
							}
						}
					}
				}
			}
		}
	}

	if len(pairs) > 0 {
		for _, pary := range pairs {
			if len(pary) == 2 {
				r = append(r, [2]*pairMarker{pary[0], pary[1]})
			}
		}
	}

	return r
}

func (bd *Board) digitsInRow(rowNo int) []int {
	digitCount := make([]int, 10) // counts of possible digits in this row
	for colNo := 0; colNo < 9; colNo++ {
		if bd[rowNo][colNo].Solved {
			continue
		}
		for _, possibleDigit := range bd[rowNo][colNo].Possible {
			digitCount[possibleDigit]++
		}
	}
	return digitCount
}

func (bd *Board) digitsInCol(colNo int) []int {
	digitCount := make([]int, 10) // counts of possible digits in this row
	for rowNo := 0; rowNo < 9; rowNo++ {
		if bd[rowNo][colNo].Solved {
			continue
		}
		for _, possibleDigit := range bd[rowNo][colNo].Possible {
			digitCount[possibleDigit]++
		}
	}
	return digitCount
}

func (bd *Board) digitsInBlock(blockNo int) []int {
	digitCount := make([]int, 10) // counts of possible digits in this block
	for rowNo := 0; rowNo < 9; rowNo++ {
		for colNo := 0; colNo < 9; colNo++ {
			if bd[rowNo][colNo].Block != blockNo {
				continue
			}
			if bd[rowNo][colNo].Solved {
				continue
			}
			for _, possibleDigit := range bd[rowNo][colNo].Possible {
				digitCount[possibleDigit]++
			}
		}
	}
	return digitCount
}

func (bd *Board) ColHiddenPairs(announce bool) int {
	eliminated := 0

	for colNo := 0; colNo < 9; colNo++ {
		digitsCount := bd.digitsInCol(colNo)
		var twoCount []int
		for digit, count := range digitsCount {
			if count == 2 {
				twoCount = append(twoCount, digit)
			}
		}
		if len(twoCount) < 2 {
			continue
		}
		// find cells that contain 2 of the digits in twoCount
		hiddenPairPairs := bd.findPairsInCol(colNo, twoCount)

		for _, pair := range hiddenPairPairs {
			if len(bd[pair[0].row][pair[0].col].Possible) > 2 ||
				len(bd[pair[1].row][pair[1].col].Possible) > 2 {
				if announce {
					fmt.Printf("Column %d Hidden pair [%d,%d]/[%d,%d] at <%d,%d> and <%d,%d>\n",
						colNo,
						pair[0].a, pair[0].b, pair[1].a, pair[1].b,
						pair[0].row, pair[0].col, pair[1].row, pair[1].col,
					)
				}
				a, b := pair[0].a, pair[0].b
				// eliminate all digits but a, b in both pairs
				for i := range pair {
					row, col := pair[i].row, pair[i].col
					if announce {
						fmt.Printf("\tColumn %d, <%d,%d>, eliminating all but %d and %d from %v\n",
							colNo, pair[i].row, pair[i].col,
							a, b,
							bd[row][col].Possible,
						)
					}
					oldCount := len(bd[row][col].Possible)
					bd[row][col].Possible = []int{a, b}
					eliminated += (oldCount - 2)
				}
			}
		}
	}

	return eliminated
}

func (bd *Board) BlockHiddenPairs(announce bool) int {
	eliminated := 0
	for blockNo := 0; blockNo < 9; blockNo++ {
		digitsCount := bd.digitsInBlock(blockNo)
		var twoCount []int
		for digit, count := range digitsCount {
			if count == 2 {
				twoCount = append(twoCount, digit)
			}
		}
		if len(twoCount) < 2 {
			continue
		}

		// find cells that contain 2 of the digits in twoCount
		hiddenPairPairs := bd.findPairsInBlock(blockNo, twoCount)

		for _, pair := range hiddenPairPairs {
			a, b := pair[0].a, pair[0].b
			if len(bd[pair[0].row][pair[0].col].Possible) > 2 ||
				len(bd[pair[1].row][pair[1].col].Possible) > 2 {
				if announce {
					fmt.Printf("Block %d Hidden pair [%d,%d]/[%d,%d] at <%d,%d> and <%d,%d>\n",
						blockNo,
						pair[0].a, pair[0].b, pair[1].a, pair[1].b,
						pair[0].row, pair[0].col, pair[1].row, pair[1].col,
					)
				}
				// eliminate all digits but a, b in both pairs
				for i := range pair {
					row, col := pair[i].row, pair[i].col
					if announce {
						fmt.Printf("\tBlock %d, <%d,%d>, eliminating all but %d and %d from %v\n",
							blockNo, pair[i].row, pair[i].col,
							a, b,
							bd[row][col].Possible,
						)
					}
					oldCount := len(bd[row][col].Possible)
					bd[row][col].Possible = []int{a, b}
					eliminated += (oldCount - 2)
				}
			}
		}
	}

	return eliminated
}
