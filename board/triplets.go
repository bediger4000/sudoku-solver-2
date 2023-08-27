package board

import (
	"fmt"
	"sort"
)

func (bd *Board) NakedTripletsEliminate(announce bool) int {
	candidatesEliminated := bd.RowNakedTriplets(announce)
	candidatesEliminated += bd.ColNakedTriplets(announce)
	candidatesEliminated += bd.BlockNakedTriplets(announce)
	return candidatesEliminated
}

func (bd *Board) RowNakedTriplets(announce bool) int {
	totalEliminated := 0
	for rowNo := 0; rowNo < 9; rowNo++ {
		var house [9]*Cell
		idx := 0
		for c := 0; c < 9; c++ {
			house[idx] = &((*bd)[rowNo][c])
			idx++
		}
		totalEliminated += findNakedTriplets(bd, "row", rowNo, house, announce)
	}
	return totalEliminated
}

func findNakedTriplets(bd *Board, phrase string, houseNo int, house [9]*Cell, announce bool) int {
	candidatesEliminated := 0

	for i := 0; i < 9; i++ {
		c1 := house[i]
		if c1.Solved || len(c1.Possible) < 2 || len(c1.Possible) > 3 {
			continue
		}
		for j := i + 1; j < 9; j++ {
			c2 := house[j]
			if c2.Solved || len(c2.Possible) < 2 || len(c2.Possible) > 3 {
				continue
			}
			for k := j + 1; k < 9; k++ {
				c3 := house[k]
				if c3.Solved || len(c3.Possible) < 2 || len(c3.Possible) > 3 {
					continue
				}

				// unsolved cells c1, c2, c3 all have 2 or 3 possible values
				// Are those possible values taken together only 3 values?
				appearances := make(map[int]int)
				for _, v := range c1.Possible {
					appearances[v]++
				}
				for _, v := range c2.Possible {
					appearances[v]++
				}
				for _, v := range c3.Possible {
					appearances[v]++
				}

				if len(appearances) != 3 {
					break // for k loop
				}

				// Only 3 possible values are keys in appearances
				notEnough := false
				for _, count := range appearances {
					if count < 2 || count > 3 {
						notEnough = true
					}
				}
				if notEnough {
					break // for k loop
				}

				var tripletValues [3]int
				idx := 0
				for v, _ := range appearances {
					tripletValues[idx] = v
					idx++
				}

				sort.Ints(tripletValues[:])

				p, q, r := tripletValues[0], tripletValues[1], tripletValues[2]

				if announce {
					fmt.Printf("%s %d, found naked triplet (%d,%d,%d) in <%d,%d>, <%d,%d>, <%d,%d>\n",
						phrase, houseNo,
						p, q, r,
						c1.Row, c1.Col,
						c2.Row, c2.Col,
						c3.Row, c3.Col,
					)
				}

				for _, cell := range house {
					if cell.Solved || cell == c1 || cell == c2 || cell == c3 {
						continue
					}
					for _, v := range cell.Possible {
						if v != p && v != q && v != r {
							continue
						}
						if x := bd.SpliceOut(cell.Row, cell.Col, v); x > 0 {
							if announce {
								fmt.Printf("\telminated %d from <%d,%d>\n", v, cell.Row, cell.Col)
							}
							candidatesEliminated++
						}
					}
				}
			}
		}
	}

	return candidatesEliminated
}

func (bd *Board) ColNakedTriplets(announce bool) int {
	totalEliminated := 0
	for colNo := 0; colNo < 9; colNo++ {
		var house [9]*Cell
		idx := 0
		for r := 0; r < 9; r++ {
			house[idx] = &((*bd)[r][colNo])
			idx++
		}
		totalEliminated += findNakedTriplets(bd, "column", colNo, house, announce)
	}
	return totalEliminated
}

func (bd *Board) BlockNakedTriplets(announce bool) int {
	totalEliminated := 0
	for blockNo := 0; blockNo < 9; blockNo++ {
		var house [9]*Cell
		idx := 0
		for r := 0; r < 9; r++ {
			for c := 0; c < 9; c++ {
				if (*bd)[r][c].Block != blockNo {
					continue
				}
				house[idx] = &((*bd)[r][c])
				idx++
			}
		}
		totalEliminated += findNakedTriplets(bd, "block", blockNo, house, announce)
	}
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
		candidatesEliminated += findTriplets(bd, count, cellsPossible, rowNo, "row", announce)
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
	for _, cls := range [][]*Cell{c1, c2, c3} {
		for _, cell := range cls {
			hash := 10*cell.Row + cell.Col
			allCells[hash] = cell
		}
	}
	if len(allCells) != 3 {
		return 0
	}

	// p, q, r all appear in the correct 3 cells.

	var tripletCells []*Cell
	for _, cell := range allCells {
		tripletCells = append(tripletCells, cell)
	}
	orderCells(tripletCells)

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

type CellSlice []*Cell

func (cs CellSlice) Len() int { return len(cs) }
func (cs CellSlice) Less(i, j int) bool {
	if cs[i].Row < cs[j].Row {
		return true
	}
	if cs[i].Col < cs[j].Col {
		return true
	}
	return false
}
func (cs CellSlice) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

func orderCells(tripletCells []*Cell) {
	sort.Sort(CellSlice(tripletCells))
}
