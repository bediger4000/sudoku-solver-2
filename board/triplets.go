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
	return totalEliminated
}

func (bd *Board) ColTriplets(announce bool) int {
	totalEliminated := 0
	return totalEliminated
}

func (bd *Board) BlockTriplets(announce bool) int {
	totalEliminated := 0
	return totalEliminated
}

func (bd *Board) HiddenTripletsEliminate(announce bool) int {
	candidatesEliminated := bd.RowHiddenTriplets(announce)
	candidatesEliminated += bd.ColHiddenTriplets(announce)
	candidatesEliminated += bd.BlockHiddenTriplets(announce)
	return candidatesEliminated
}

func (bd *Board) RowHiddenTriplets(announce bool) int {
	candidatesEliminated := 0

	for rowNo := 0; rowNo < 9; rowNo++ {
		bdRow := (*bd)[rowNo]
		cellsPossible := make(map[int][]*Cell)
		var count [10]int
		for c := 0; c < 9; c++ {
			if bdRow[c].Solved {
				continue
			}
			for _, v := range bdRow[c].Possible {
				count[v]++
				cellsPossible[v] = append(cellsPossible[v], &(bdRow[c]))
			}
		}
		var p2or3 []int
		for i := 1; i < 10; i++ {
			if count[i] == 2 || count[i] == 3 {
				p2or3 = append(p2or3, i)
			}
		}
		ln := len(p2or3)
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
			if len(c1) < 2 || len(c1) > 3 {
				continue
			}
			for j := i + 1; j < ln; j++ {
				c2 := cellsPossible[p2or3[j]]
				if len(c2) < 2 || len(c2) > 3 {
					continue
				}
				for k := j + 1; k < ln; k++ {
					c3 := cellsPossible[p2or3[k]]
					if len(c3) < 2 || len(c3) > 3 {
						continue
					}
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
						// p2or3[i], p2or3[j], p2or3[k] appear in the correct 3 cells.
						var tripletCells [3]*Cell
						idx := 0
						for _, cell := range all3cells {
							tripletCells[idx] = cell
							idx++
						}
						if announce {
							fmt.Printf("triplet <%d,%d,%d> all found in cells <%d,%d>, <%d,%d>, <%d,%d>\n",
								p2or3[i], p2or3[j], p2or3[k],
								tripletCells[0].Row, tripletCells[0].Col,
								tripletCells[1].Row, tripletCells[1].Col,
								tripletCells[2].Row, tripletCells[2].Col,
							)
						}
						// splice out all possible values except the
						// values in p2or3, from the cells in tripletCells
						for _, cell := range tripletCells {
							for _, v := range cell.Possible {
								if v == p2or3[i] ||
									v == p2or3[j] ||
									v == p2or3[k] {
									continue
								}
								spliced := bd.SpliceOut(cell.Row, cell.Col, v)
								if announce && spliced == 1 {
									fmt.Printf("\teliminated %d at <%d,%d>\n", v, cell.Row, cell.Col)
								}
								candidatesEliminated += spliced
							}
						}
					}
				}
			}
		}
	}

	return candidatesEliminated
}

func (bd *Board) ColHiddenTriplets(announce bool) int {
	candidatesEliminated := 0

	for colNo := 0; colNo < 9; colNo++ {
		cellsPossible := make(map[int][]*Cell)
		var count [10]int
		for r := 0; r < 9; r++ {
			if (*bd)[r][colNo].Solved {
				continue
			}
			for _, v := range (*bd)[r][colNo].Possible {
				count[v]++
				cellsPossible[v] = append(cellsPossible[v], &((*bd)[r][colNo]))
			}
		}
		var p2or3 []int
		for i := 1; i < 10; i++ {
			if count[i] == 2 || count[i] == 3 {
				p2or3 = append(p2or3, i)
			}
		}
		ln := len(p2or3)
		if ln < 3 {
			if announce {
				fmt.Printf("column %d does not have possible triplets\n", colNo)
			}
			continue // colNo loop
		}
		// At least 3 possible values with 2 or 3 appearances each.
		// find out if they appear in only 3 cells
		for i := 0; i < ln; i++ {
			c1 := cellsPossible[p2or3[i]]
			if len(c1) < 2 || len(c1) > 3 {
				continue
			}
			for j := i + 1; j < ln; j++ {
				c2 := cellsPossible[p2or3[j]]
				if len(c2) < 2 || len(c2) > 3 {
					continue
				}
				for k := j + 1; k < ln; k++ {
					c3 := cellsPossible[p2or3[k]]
					if len(c3) < 2 || len(c3) > 3 {
						continue
					}
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
						// p2or3[i], p2or3[j], p2or3[k] appear in the correct 3 cells.
						var tripletCells [3]*Cell
						idx := 0
						for _, cell := range all3cells {
							tripletCells[idx] = cell
							idx++
						}
						if announce {
							fmt.Printf("triplet <%d,%d,%d> all found in cells <%d,%d>, <%d,%d>, <%d,%d>\n",
								p2or3[i], p2or3[j], p2or3[k],
								tripletCells[0].Row, tripletCells[0].Col,
								tripletCells[1].Row, tripletCells[1].Col,
								tripletCells[2].Row, tripletCells[2].Col,
							)
						}
						// splice out all possible values except the
						// values in p2or3, from the cells in tripletCells
						for _, cell := range tripletCells {
							for _, v := range cell.Possible {
								if v == p2or3[i] ||
									v == p2or3[j] ||
									v == p2or3[k] {
									continue
								}
								spliced := bd.SpliceOut(cell.Row, cell.Col, v)
								if announce && spliced == 1 {
									fmt.Printf("\teliminated %d at <%d,%d>\n", v, cell.Row, cell.Col)
								}
								candidatesEliminated += spliced
							}
						}
					}
				}
			}
		}
	}

	return candidatesEliminated
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
