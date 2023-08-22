package board

import (
	"fmt"
)

func (bd *Board) Triplets(announce bool) int {
	candidatesEliminated := bd.RowTriplets(announce)
	candidatesEliminated += bd.ColTriplets(announce)
	candidatesEliminated += bd.BlockTriplets(announce)
	return candidatesEliminated
}

func (bd *Board) RowTriplets(announce bool) int {
	totalEliminated := 0

	for rowNo := 0; rowNo < 9; rowNo++ {

		// find cells in row with only 2 possibilities
		var cellsWith2 []*pairMarker
		for colNo := 0; colNo < 9; colNo++ {
			if len(bd[rowNo][colNo].Possible) == 2 {
				// fmt.Printf("Appending <%d,%d>, a: %d, b: %d\n",
				//	rowNo, colNo, bd[rowNo][colNo].Possible[0], bd[rowNo][colNo].Possible[1])
				rep := 10*bd[rowNo][colNo].Possible[0] + bd[rowNo][colNo].Possible[1]
				cellsWith2 = append(cellsWith2, &pairMarker{row: rowNo, col: colNo, a: bd[rowNo][colNo].Possible[0], b: bd[rowNo][colNo].Possible[1], rep: rep})
			}
		}

		// there have to be at least 3 cells with only a pair of possibilities
		if len(cellsWith2) < 3 {
			continue
		}

		// find 3 cells where .a, .b have values p, q, r, (p,q), (q,r), (p,r)
		ln := len(cellsWith2)
		for i := range cellsWith2 {
			for j := i + 1; j < ln; j++ {
				// fmt.Printf("i %d: a: %d, b: %d\n", i, cellsWith2[i].a, cellsWith2[i].b)
				// fmt.Printf("j %d: a: %d, b: %d\n", j, cellsWith2[j].a, cellsWith2[j].b)
				var p, q, r int
				oneValueMatches := false
				if cellsWith2[i].a == cellsWith2[j].a && cellsWith2[i].b != cellsWith2[j].b {
					p = cellsWith2[i].a // common value
					q = cellsWith2[i].b
					r = cellsWith2[j].b
					// fmt.Printf("match 1, p: %d, q: %d, r: %d\n", p, q, r)
					oneValueMatches = true
				}
				if cellsWith2[i].a == cellsWith2[j].b && cellsWith2[i].b != cellsWith2[j].a {
					p = cellsWith2[i].a // common value
					q = cellsWith2[i].b
					r = cellsWith2[j].a
					// fmt.Printf("match 2, p: %d, q: %d, r: %d\n", p, q, r)
					oneValueMatches = true
				}
				if cellsWith2[i].b == cellsWith2[j].b && cellsWith2[i].a != cellsWith2[j].a {
					p = cellsWith2[i].b // common value
					q = cellsWith2[i].a
					r = cellsWith2[j].a
					// fmt.Printf("match 3, p: %d, q: %d, r: %d\n", p, q, r)
					oneValueMatches = true
				}
				if cellsWith2[i].b == cellsWith2[j].a && cellsWith2[i].a != cellsWith2[j].b {
					p = cellsWith2[i].b // common value
					q = cellsWith2[i].a
					r = cellsWith2[j].b
					// fmt.Printf("match 3, p: %d, q: %d, r: %d\n", p, q, r)
					oneValueMatches = true
				}
				if !oneValueMatches {
					continue
				}
				// find a cell where <a,b> == <q,r> or <a,b> == <r,q>
				for k := j + 1; k < ln; k++ {
					// fmt.Printf("k %d: a: %d, b: %d\n", k, cellsWith2[k].a, cellsWith2[k].b)
					if cellsWith2[k].a == q && cellsWith2[k].b == r ||
						cellsWith2[k].a == r && cellsWith2[k].b == q {
						if announce {
							fmt.Printf("found triplet (%d,%d,%d) in row %d, cols %d, %d, %d\n",
								p, q, r, rowNo,
								cellsWith2[i].col, cellsWith2[j].col, cellsWith2[k].col,
							)
						}
						// Eliminate any p,q,r possibiliites in other cells in this row
						for c := 0; c < 9; c++ {
							if cellsWith2[i].col == c || cellsWith2[j].col == c || cellsWith2[k].col == c {
								// one of the cells that have the 3 triplet values
								continue
							}
							n := 0
							if pE := bd.SpliceOut(rowNo, c, p); announce && pE == 1 {
								fmt.Printf("\teliminated %d at <%d,%d>\n", p, rowNo, c)
								totalEliminated++
								n++
							}
							if qE := bd.SpliceOut(rowNo, c, q); announce && qE == 1 {
								fmt.Printf("\teliminated %d at <%d,%d>\n", q, rowNo, c)
								totalEliminated++
								n++
							}
							if rE := bd.SpliceOut(rowNo, c, r); announce && rE == 1 {
								fmt.Printf("\teliminated %d at <%d,%d>\n", q, rowNo, c)
								totalEliminated++
								n++
							}
							// if more than 1 possibility gets eliminated, something's wrong.
							if n > 1 {
								fmt.Printf("PROBLEM! Eliminated %d possibilities with triplet (%d,%d,%d) in row %d, cols %d, %d, %d\n",
									n, p, q, r, rowNo,
									cellsWith2[i].col, cellsWith2[j].col, cellsWith2[k].col,
								)
							}
						}
					}
				}
			}
		}
	}

	return totalEliminated
}

func (bd *Board) ColTriplets(announce bool) int {
	totalEliminated := 0

	for colNo := 0; colNo < 9; colNo++ {

		// find cells in row with only 2 possibilities
		var cellsWith2 []*pairMarker
		for rowNo := 0; rowNo < 9; rowNo++ {
			if len(bd[rowNo][colNo].Possible) == 2 {
				// fmt.Printf("Appending <%d,%d>, a: %d, b: %d\n",
				//	rowNo, colNo, bd[rowNo][colNo].Possible[0], bd[rowNo][colNo].Possible[1])
				rep := 10*bd[rowNo][colNo].Possible[0] + bd[rowNo][colNo].Possible[1]
				cellsWith2 = append(cellsWith2, &pairMarker{row: rowNo, col: colNo, a: bd[rowNo][colNo].Possible[0], b: bd[rowNo][colNo].Possible[1], rep: rep})
			}
		}

		// there have to be at least 3 cells with only a pair of possibilities
		if len(cellsWith2) < 3 {
			continue
		}

		// find 3 cells where .a, .b have values p, q, r, (p,q), (q,r), (p,r)
		ln := len(cellsWith2)
		for i := range cellsWith2 {
			for j := i + 1; j < ln; j++ {
				// fmt.Printf("i %d: a: %d, b: %d\n", i, cellsWith2[i].a, cellsWith2[i].b)
				// fmt.Printf("j %d: a: %d, b: %d\n", j, cellsWith2[j].a, cellsWith2[j].b)
				var p, q, r int
				oneValueMatches := false
				if cellsWith2[i].a == cellsWith2[j].a && cellsWith2[i].b != cellsWith2[j].b {
					p = cellsWith2[i].a // common value
					q = cellsWith2[i].b
					r = cellsWith2[j].b
					// fmt.Printf("match 1, p: %d, q: %d, r: %d\n", p, q, r)
					oneValueMatches = true
				}
				if cellsWith2[i].a == cellsWith2[j].b && cellsWith2[i].b != cellsWith2[j].a {
					p = cellsWith2[i].a // common value
					q = cellsWith2[i].b
					r = cellsWith2[j].a
					// fmt.Printf("match 2, p: %d, q: %d, r: %d\n", p, q, r)
					oneValueMatches = true
				}
				if cellsWith2[i].b == cellsWith2[j].b && cellsWith2[i].a != cellsWith2[j].a {
					p = cellsWith2[i].b // common value
					q = cellsWith2[i].a
					r = cellsWith2[j].a
					// fmt.Printf("match 3, p: %d, q: %d, r: %d\n", p, q, r)
					oneValueMatches = true
				}
				if cellsWith2[i].b == cellsWith2[j].a && cellsWith2[i].a != cellsWith2[j].b {
					p = cellsWith2[i].b // common value
					q = cellsWith2[i].a
					r = cellsWith2[j].b
					// fmt.Printf("match 3, p: %d, q: %d, r: %d\n", p, q, r)
					oneValueMatches = true
				}
				if !oneValueMatches {
					continue
				}
				// find a cell where <a,b> == <q,r> or <a,b> == <r,q>
				for k := j + 1; k < ln; k++ {
					// fmt.Printf("k %d: a: %d, b: %d\n", k, cellsWith2[k].a, cellsWith2[k].b)
					if cellsWith2[k].a == q && cellsWith2[k].b == r ||
						cellsWith2[k].a == r && cellsWith2[k].b == q {
						if announce {
							fmt.Printf("found triplet (%d,%d,%d) in col %d, cols %d, %d, %d\n",
								p, q, r, colNo,
								cellsWith2[i].col, cellsWith2[j].col, cellsWith2[k].col,
							)
						}
						// Eliminate any p,q,r possibiliites in other cells in this row
						for row := 0; row < 9; row++ {
							if cellsWith2[i].row == row || cellsWith2[j].row == row || cellsWith2[k].row == row {
								// one of the cells that have the 3 triplet values
								continue
							}
							n := 0
							if pE := bd.SpliceOut(row, colNo, p); announce && pE == 1 {
								fmt.Printf("\teliminated %d at <%d,%d>\n", p, row, colNo)
								totalEliminated++
								n++
							}
							if qE := bd.SpliceOut(row, colNo, q); announce && qE == 1 {
								fmt.Printf("\teliminated %d at <%d,%d>\n", q, row, colNo)
								totalEliminated++
								n++
							}
							if rE := bd.SpliceOut(row, colNo, r); announce && rE == 1 {
								fmt.Printf("\teliminated %d at <%d,%d>\n", q, row, colNo)
								totalEliminated++
								n++
							}
							// if more than 1 possibility gets eliminated, something's wrong.
							if n > 1 {
								fmt.Printf("PROBLEM! Eliminated %d possibilities with triplet (%d,%d,%d) in colNo %d, rows %d, %d, %d\n",
									n, p, q, r, colNo,
									cellsWith2[i].row, cellsWith2[j].row, cellsWith2[k].row,
								)
							}
						}
					}
				}
			}
		}
	}

	return totalEliminated
}

func (bd *Board) BlockTriplets(announce bool) int {
	totalEliminated := 0

	return totalEliminated
}

func (bd *Board) HiddenTriplets(announce bool) int {
	candidatesEliminated := bd.RowHiddenTriplets(announce)
	candidatesEliminated += bd.ColHiddenTriplets(announce)
	candidatesEliminated += bd.BlockHiddenTriplets(announce)
	return candidatesEliminated
}

func (bd *Board) RowHiddenTriplets(announce bool) int {
	fmt.Println("Enter RowHiddenTriplets")
	defer fmt.Println("Exit RowHiddenTriplets")
	candidatesEliminated := 0

	for rowNo := 0; rowNo < 9; rowNo++ {
		fmt.Printf("Examining row %d\n", rowNo)
		bdRow := (*bd)[rowNo]
		cellsPossible := make(map[int][]*Cell)
		var count [10]int
		for c := 0; c < 9; c++ {
			for _, v := range bdRow[c].Possible {
				count[v]++
				cellsPossible[v] = append(cellsPossible[v], &(bdRow[c]))
			}
		}
		fmt.Printf("Values count %v\n", count)
		var p2or3 []int
		for i := 1; i < 10; i++ {
			if count[i] == 2 || count[i] == 3 {
				p2or3 = append(p2or3, i)
			}
		}
		ln := len(p2or3)
		fmt.Printf("\t%d values appear 2 or 3 times: %v\n", ln, p2or3)
		if ln < 3 {
			if announce {
				fmt.Printf("row %d does not have possible triplets\n", rowNo)
			}
			continue // rowNo loop
		}
		// At least 3 possible values with 2 or 3 appearances each.
		// find out if they appear in only 3 cells
		for i := 0; i < ln; i++ {
			c1 := cellsPossible[p2or3[i]]
			fmt.Printf("\tvalue 1 %d appears in %d cells\n", p2or3[i], len(c1))
			if len(c1) < 2 || len(c1) > 3 {
				continue
			}
			for j := i + 1; j < ln; j++ {
				c2 := cellsPossible[p2or3[j]]
				fmt.Printf("\tvalue 2 %d appears in %d cells\n", p2or3[j], len(c2))
				if len(c2) < 2 || len(c2) > 3 {
					continue
				}
				for k := j + 1; k < ln; k++ {
					c3 := cellsPossible[p2or3[k]]
					fmt.Printf("\tvalue 3 %d appears in %d cells\n", p2or3[k], len(c3))
					if len(c3) < 2 || len(c3) > 3 {
						continue
					}
					fmt.Printf("\t values %d,%d,%d all found in 3 cells\n",
						p2or3[i], p2or3[j], p2or3[k],
					)
					// p2or3[i], p2or3[j], p2or3[k] appear in 3 cells.
					// are they 3 and only 3 different cells?
					all3cells := make(map[int]*Cell)
					for _, cls := range [][]*Cell{c1, c2, c3} {
						for _, cell := range cls {
							hash := 10*cell.Row + cell.Col
							all3cells[hash] = cell
						}
					}
					if len(all3cells) == 3 {
						var tripletCells [3]*Cell
						idx := 0
						for _, cell := range all3cells {
							tripletCells[idx] = cell
							idx++
						}
						if announce {
							fmt.Printf("triple <%d,%d,%d> all found in cells <%d,%d>, <%d,%d>, <%d,%d>\n",
								p2or3[i], p2or3[j], p2or3[k],
								tripletCells[0].Row, tripletCells[0].Col,
								tripletCells[1].Row, tripletCells[1].Col,
								tripletCells[2].Row, tripletCells[2].Col,
							)
						}

					}
				}
			}
		}
	}

	return candidatesEliminated
}

func (bd *Board) ColHiddenTriplets(announce bool) int {
	return 0
}

func (bd *Board) BlockHiddenTriplets(announce bool) int {
	return 0
}

func countRowPossible(bd *Board, rowNo int) map[int]int {
	possiblesCount := make(map[int]int)
	row := (*bd)[rowNo]
	for c := 0; c < 9; c++ {
		if row[c].Solved {
			continue
		}
		for j := range row[c].Possible {
			possiblesCount[row[c].Possible[j]]++
		}
	}
	return possiblesCount
}

func countColPossible(bd *Board, colNo int) map[int]int {
	possiblesCount := make(map[int]int)
	for r := 0; r < 9; r++ {
		if (*bd)[r][colNo].Solved {
			continue
		}
		for j := range (*bd)[r][colNo].Possible {
			possiblesCount[(*bd)[r][colNo].Possible[j]]++
		}
	}
	return possiblesCount
}

func countBlockPossible(bd *Board, blockNo int) map[int]int {
	possiblesCount := make(map[int]int)
	return possiblesCount
}
