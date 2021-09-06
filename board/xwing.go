package board

import "fmt"

func (bd *Board) XwingEliminate(announce bool) int {

	for rowNo := 0; rowNo < 9; rowNo++ {
		// find cells in row with only 2 possibilities
		cellsWith := make(map[int][]*pairMarker)
		for colNo := 0; colNo < 9; colNo++ {
			if len(bd[rowNo][colNo].Possible) == 2 {
				rep := 10*bd[rowNo][colNo].Possible[0] + bd[rowNo][colNo].Possible[1]
				cellsWith[rep] = append(cellsWith[rep],
					&pairMarker{
						row: rowNo,
						col: colNo,
						a:   bd[rowNo][colNo].Possible[0],
						b:   bd[rowNo][colNo].Possible[1],
						rep: rep,
					},
				)
			}
		}

		for _, cells := range cellsWith {
			if len(cells) == 2 {
				// find out if any row has a possible value of cells[0].a or cells[0].b
				// in columns cells[0].col and cells[1].col
				a, b := cells[0].a, cells[0].b
				colA, colB := cells[0].col, cells[1].col
				for xrow := 0; xrow < 9; xrow++ {
					if xrow == rowNo {
						continue
					}
					if bd[xrow][colA].Solved || bd[xrow][colB].Solved {
						continue
					}
					var possible [10]int
					for _, poss := range bd[xrow][colA].Possible {
						possible[poss]++
					}
					for _, poss := range bd[xrow][colB].Possible {
						possible[poss]++
					}

					for i := 1; i <= 9; i++ {
						if possible[i] == 2 {
							if i == a {
								fmt.Printf("can xwing %d in row %d, cols %d, %d\n",
									a, xrow, colA, colB,
								)
							} else if i == b {
								fmt.Printf("can xwing %d in row %d, cols %d, %d\n",
									b, xrow, colA, colB,
								)
							}
							fmt.Printf("vs [%d,%d] in <%d,%d> <%d,%d>\n",
								a, b,
								cells[0].row, cells[0].col,
								cells[1].row, cells[1].col,
							)
						}
					}
				}
			}
		}
	}

	return 0
}
