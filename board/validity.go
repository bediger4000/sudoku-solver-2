package board

import "fmt"

func (bd *Board) Valid() bool {
	for rowNo := 0; rowNo < 9; rowNo++ {
		digitCounts := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		incomplete := false
		invalid := false
		sum := 0
		for colNo := 0; colNo < 9; colNo++ {
			if bd[rowNo][colNo].Solved {
				digitCounts[bd[rowNo][colNo].Value]++
				sum += bd[rowNo][colNo].Value
			} else {
				incomplete = true
			}
		}
		for digit, count := range digitCounts {
			if digit != 0 && count > 1 {
				fmt.Printf("Row %d has %d %d digits\n", rowNo, count, digit)
				invalid = true
			}
		}
		if !incomplete && sum != 45 {
			fmt.Printf("Something wrong with row %d\n", rowNo)
			invalid = true
		}
		if invalid {
			fmt.Printf("Row %d invalid\n", rowNo)
		}
	}

	for colNo := 0; colNo < 9; colNo++ {
		digitCounts := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		incomplete := false
		invalid := false
		sum := 0
		for rowNo := 0; rowNo < 9; rowNo++ {
			if bd[rowNo][colNo].Solved {
				digitCounts[bd[rowNo][colNo].Value]++
				sum += bd[rowNo][colNo].Value
			} else {
				incomplete = true
			}
		}
		for digit, count := range digitCounts {
			if digit != 0 && count > 1 {
				fmt.Printf("Col %d has %d %d digits\n", colNo, count, digit)
				invalid = true
			}
		}
		if !incomplete && sum != 45 {
			fmt.Printf("Something wrong with col %d\n", colNo)
			invalid = true
		}
		if invalid {
			fmt.Printf("Column %d invalid\n", colNo)
		}
	}

	for blockNo := 0; blockNo < 9; blockNo++ {
		digitCounts := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		incomplete := false
		invalid := false
		sum := 0
		for colNo := 0; colNo < 9; colNo++ {
			for rowNo := 0; rowNo < 9; rowNo++ {
				if bd[rowNo][colNo].Block != blockNo {
					continue
				}
				if bd[rowNo][colNo].Solved {
					digitCounts[bd[rowNo][colNo].Value]++
					sum += bd[rowNo][colNo].Value
				} else {
					incomplete = true
				}
			}
		}
		if !incomplete && sum != 45 {
			fmt.Printf("Something wrong with block %d\n", blockNo)
			invalid = true
		}
		for digit, count := range digitCounts {
			if digit != 0 && count > 1 {
				fmt.Printf("Block %d has %d %d digits\n", blockNo, count, digit)
				invalid = true
			}
		}
		if invalid {
			fmt.Printf("Block %d invalid\n", blockNo)
		}
	}
	return true
}

// Finished returns true if the board is filled in,
// false otherwise. Makes no judgement of the validity
// of the filled-in numbers.
func (bd *Board) Finished() bool {
	for colNo := 0; colNo < 9; colNo++ {
		for rowNo := 0; rowNo < 9; rowNo++ {
			if !bd[rowNo][colNo].Solved {
				return false
			}
		}
	}
	return true
}
