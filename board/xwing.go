package board

import "fmt"

func (bd *Board) XwingEliminate(announce bool) int {
	var eliminated int
	eliminated += bd.xwingRowEliminate(announce)
	eliminated += bd.xwingColEliminate(announce)
	return eliminated
}

// xwingRowEliminate finds 2 rows with some possible value in the same columns,
// eliminates that possible value from the columns.
func (bd *Board) xwingRowEliminate(announce bool) int {

	var colsForRow [9]map[int][]int

	for idx := range colsForRow {
		colsForRow[idx] = make(map[int][]int)
	}

	for rowNo := 0; rowNo < 9; rowNo++ {
		for colNo := 0; colNo < 9; colNo++ {
			if bd[rowNo][colNo].Solved {
				continue
			}
			for _, poss := range bd[rowNo][colNo].Possible {
				colsForRow[rowNo][poss] = append(colsForRow[rowNo][poss], colNo)
			}
		}
	}

	eliminatedCount := 0

	for rowNo := 0; rowNo < 9; rowNo++ {
		colsFor := colsForRow[rowNo]
		for poss := 1; poss <= 9; poss++ {
			if len(colsFor[poss]) == 2 {
				for otherRow := 0; otherRow < 9; otherRow++ {
					if rowNo == otherRow {
						continue
					}
					otherCols := colsForRow[otherRow]
					if len(otherCols[poss]) == 2 {
						if (colsFor[poss][0] == otherCols[poss][0] && colsFor[poss][1] == otherCols[poss][1]) ||
							(colsFor[poss][0] == otherCols[poss][1] && colsFor[poss][1] == otherCols[poss][0]) {
							// Eliminate poss in colsFor[poss][0] and colsFor[poss][1]
							if announce {
								fmt.Printf(
									"Xwing for %d, row <%d,%d>/<%d,%d>  <%d,%d>/<%d,%d>\n",
									poss,
									rowNo, colsFor[poss][0],
									rowNo, colsFor[poss][1],
									otherRow, otherCols[poss][0],
									otherRow, otherCols[poss][1],
								)
							}
							colA, colB := colsFor[poss][0], colsFor[poss][1]
							for newRow := 0; newRow < 9; newRow++ {
								if newRow == rowNo || newRow == otherRow {
									continue
								}
								m := bd.SpliceOut(newRow, colA, poss)
								if m == 1 && announce {
									fmt.Printf("Xwing eliminate %d at <%d,%d>\n", poss, newRow, colA)
								}
								eliminatedCount += m
								m = bd.SpliceOut(newRow, colB, poss)
								if m == 1 && announce {
									fmt.Printf("Xwing eliminate %d at <%d,%d>\n", poss, newRow, colB)
								}
								eliminatedCount += m
							}
						}
					}
				}
			}
		}
	}

	return eliminatedCount
}

// xwingColEliminate finds 2 cols with some possible value in the same rows,
// eliminates that possible value from the rows.
func (bd *Board) xwingColEliminate(announce bool) int {
	var rowsForCol [9]map[int][]int

	for idx := range rowsForCol {
		rowsForCol[idx] = make(map[int][]int)
	}

	for colNo := 0; colNo < 9; colNo++ {
		for rowNo := 0; rowNo < 9; rowNo++ {
			if bd[rowNo][colNo].Solved {
				continue
			}
			for _, poss := range bd[rowNo][colNo].Possible {
				rowsForCol[colNo][poss] = append(rowsForCol[colNo][poss], rowNo)
			}
		}
	}

	eliminatedCount := 0

	for colNo := 0; colNo < 9; colNo++ {
		rowsFor := rowsForCol[colNo]
		for poss := 1; poss <= 9; poss++ {
			if len(rowsFor[poss]) == 2 {
				for otherCol := 0; otherCol < 9; otherCol++ {
					if colNo == otherCol {
						continue
					}
					otherRows := rowsForCol[otherCol]
					if len(otherRows[poss]) == 2 {
						if (rowsFor[poss][0] == otherRows[poss][0] && rowsFor[poss][1] == otherRows[poss][1]) ||
							(rowsFor[poss][0] == otherRows[poss][1] && rowsFor[poss][1] == otherRows[poss][0]) {
							if announce {
								fmt.Printf(
									"Xwing for %d, col <%d,%d>/<%d,%d>  <%d,%d>/<%d,%d>\n",
									poss,
									rowsFor[poss][0], colNo,
									rowsFor[poss][1], colNo,
									otherRows[poss][0], otherCol,
									otherRows[poss][1], otherCol,
								)
							}
							rowA, rowB := rowsFor[poss][0], rowsFor[poss][1]
							for newCol := 0; newCol < 9; newCol++ {
								if newCol == colNo || newCol == otherCol {
									continue
								}
								m := bd.SpliceOut(rowA, newCol, poss)
								if m == 1 && announce {
									fmt.Printf("Xwing eliminate %d at <%d,%d>\n", poss, rowA, newCol)
								}
								eliminatedCount += m
								m = bd.SpliceOut(rowB, newCol, poss)
								if m == 1 && announce {
									fmt.Printf("Xwing eliminate %d at <%d,%d>\n", poss, rowB, newCol)
								}
								eliminatedCount += m
							}
						}
					}
				}
			}
		}
	}

	return eliminatedCount
}
