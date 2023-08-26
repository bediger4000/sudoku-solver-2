package board

import (
	"fmt"
)

func (bd *Board) NakedTripletsEliminate(announce bool) int {
	candidatesEliminated := bd.RowNakedTriplets(announce)
	candidatesEliminated += bd.ColNakedTriplets(announce)
	candidatesEliminated += bd.BlockNakedTriplets(announce)
	return candidatesEliminated
}

func (bd *Board) RowNakedTriplets(announce bool) int {
	totalEliminated := 0
	return totalEliminated
}

func (bd *Board) ColNakedTriplets(announce bool) int {
	totalEliminated := 0
	return totalEliminated
}

func (bd *Board) BlockNakedTriplets(announce bool) int {
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
		// find out if any 3 values in p2or3 appear in only 3 cells
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
					// all 3 values in p2or3 appear in exactly 3 cells, c1, c2, c3
					candidatesEliminated += elimNonTriples(bd, c1, c2, c3, p2or3[i], p2or3[j], p2or3[k], announce)
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
		candidatesEliminated += findTriplets(bd, count, cellsPossible, colNo, "column", announce)
	}

	return candidatesEliminated
}

// findTriplets looks through count and cellsPossible, finds triplets,
// eliminates non-triplet possible values from the cells holding the triplets
func findTriplets(bd *Board, count [10]int, cellsPossible map[int][]*Cell, number int, house string, announce bool) int {
	candidatesEliminated := 0
	// p2or3 holds all values in this "house" (row, col or block) that
	// appear either 2 or 3 times.
	var p2or3 []int
	for i := 1; i < 10; i++ {
		if count[i] == 2 || count[i] == 3 {
			p2or3 = append(p2or3, i)
		}
	}
	ln := len(p2or3)

	if ln < 3 {
		if announce {
			fmt.Printf("%s %d does not have possible triplets\n", house, number)
		}
		return 0
	}

	// At least 3 possible values with 2 or 3 appearances each.
	// find out if they appear in only 3 cells
	for i := 0; i < ln; i++ {
		c1 := cellsPossible[p2or3[i]]
		if len(c1) < 2 || len(c1) > 3 {
			// value p2or3[i] appears less than 2 or more than
			// 3 times in this house
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
				// all 3 values in p2or3 appear in exactly 3 cells, c1, c2, c3
				candidatesEliminated += elimNonTriples(bd, c1, c2, c3, p2or3[i], p2or3[j], p2or3[k], announce)
			}
		}
	}
	return candidatesEliminated
}

// elimNonTriples eliminates all values *other* that the triplet
// p,q,r from the 3 cells c1, c2, c3 that contain them.
func elimNonTriples(bd *Board, c1, c2, c3 []*Cell, p, q, r int, announce bool) int {
	candidatesEliminated := 0

	// p, q, r appear as possible values in 3 cells.
	// Are those cells exactly 3 different cells?
	allCells := make(map[int]*Cell)
	var tripletCells []*Cell
	for _, cls := range [][]*Cell{c1, c2, c3} {
		for _, cell := range cls {
			hash := 10*cell.Row + cell.Col
			allCells[hash] = cell
			tripletCells = append(tripletCells, cell)
		}
	}
	if len(allCells) != 3 {
		return 0
	}

	// p, q, r all appear in the correct 3 cells.

	if announce {
		fmt.Printf("triplet (%d,%d,%d) all found in cells <%d,%d>, <%d,%d>, <%d,%d>\n",
			p, q, r,
			tripletCells[0].Row, tripletCells[0].Col,
			tripletCells[1].Row, tripletCells[1].Col,
			tripletCells[2].Row, tripletCells[2].Col,
		)
	}

	// splice out all possible values except the
	// values in p2or3, from the cells in tripletCells
	for _, cell := range tripletCells {
		for idx := 0; idx < len(cell.Possible); {
			v := cell.Possible[idx]
			if v == p ||
				v == q ||
				v == r {
				idx++
				continue
			}
			if spliced := bd.SpliceOut(cell.Row, cell.Col, v); spliced > 0 {
				if announce {
					fmt.Printf("\teliminated %d at <%d,%d>\n", v, cell.Row, cell.Col)
				}
				candidatesEliminated += spliced
				continue
			}
			idx++
		}
	}
	return candidatesEliminated
}

func (bd *Board) BlockHiddenTriplets(announce bool) int {
	candidatesEliminated := 0

	for blockNo := 0; blockNo < 9; blockNo++ {
		cellsPossible := make(map[int][]*Cell)
		var count [10]int
		for r := 0; r < 9; r++ {
			for c := 0; c < 9; c++ {
				if (*bd)[r][c].Block != blockNo || (*bd)[r][c].Solved {
					continue
				}
				for _, v := range (*bd)[r][c].Possible {
					count[v]++
					cellsPossible[v] = append(cellsPossible[v], &((*bd)[r][c]))
				}
			}
		}
		candidatesEliminated += findTriplets(bd, count, cellsPossible, blockNo, "block", announce)
	}

	return candidatesEliminated
}
