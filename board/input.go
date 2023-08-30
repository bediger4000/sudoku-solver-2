package board

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"unicode"
)

func (bd *Board) Print(out io.Writer) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if bd[i][j].Solved {
				fmt.Fprintf(out, "%1d ", bd[i][j].Value)
			} else {
				fmt.Fprintf(out, "_ ")
			}
		}
		fmt.Fprintf(out, "\n")
	}
}

func (bd *Board) Details(out io.Writer) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Fprintf(out, "<%d,%d>, block %d\n", bd[i][j].Row, bd[i][j].Col, bd[i][j].Block)
			fmt.Fprintf(out, "Solved %v\n", bd[i][j].Solved)
			fmt.Fprintf(out, "Value %v\n", bd[i][j].Value)
			fmt.Fprintf(out, "Possible %v\n", bd[i][j].Possible)
		}
		fmt.Fprintf(out, "\n")
	}
}

func (bd *Board) PrintAsInput(out io.Writer) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if bd[i][j].Solved {
				fmt.Fprintf(out, "%1d", bd[i][j].Value)
			} else {
				fmt.Fprintf(out, ".")
			}
		}
		fmt.Fprintf(out, "\n")
	}
}

func ReadBoard(in io.Reader) (*Board, error) {
	bd := New()
	r := bufio.NewReader(in)
	for row := 0; row < 9; {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				if row != 9 {
					return nil, fmt.Errorf("premature end-of-file at row %d\n", row)
				}
				break
			}
			return nil, err
		}
		if buf[0] == '#' {
			continue
		}
		buf = bytes.Trim(buf, " \t\n\r")

		col := 0
		for _, c := range buf {
			if c == ',' || c == ' ' {
				continue
			}
			n := int(c - '0')
			if c == '_' || c == '.' {
				n = 0
			}
			if n < 0 || n > 10 {
				// Will this ever happen?
				return nil, fmt.Errorf("numbers must be less than 10, greater than zero: %d (%c)\n", n, c)
			}
			if n != 0 {
				bd.MarkSolved(row, col, n)
			}
			col++
		}
		if col != 9 {
			return nil, fmt.Errorf("row %d had %d cols\n", row+1, col)
		}
		row++
	}

	// Based on the digits given in the input,
	// eliminate those digits in the row, column
	// and block they appear in.
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if bd[row][col].Solved {
				bd.EliminatePossibilities(row, col, bd[row][col].Block, bd[row][col].Value)
			}
		}
	}
	return bd, nil
}

func NewBoardFromString(str string) (*Board, error) {
	bd := New()

	for idx, r := range str {
		if !unicode.IsDigit(r) {
			return nil, fmt.Errorf("character %d: %c not a digit 0-9", idx+1, r)
		}

		n := r - '0'
		if n < 0 || n > 9 {
			return nil, fmt.Errorf("character %d: %c not a digit 0-9", idx+1, r)
		}

		if n == 0 {
			continue
		}

		x := idx / 9
		y := idx % 9

		bd.MarkSolved(x, y, int(n))
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if bd[row][col].Solved {
				bd.EliminatePossibilities(row, col, bd[row][col].Block, bd[row][col].Value)
			}
		}
	}

	return bd, nil
}
func FigureOutBoard() (*Board, error) {

	if flag.NArg() > 0 {
		// Something left on command line, might be SDM string, might be file name

		fin, err := os.Open(flag.Arg(0))
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				// It's probably an 81-character, 1-line, SDM format string
				return NewBoardFromString(flag.Arg(0))
			} else {
				// some real problem with the file
				return nil, err
			}
		} else {
			// It is an openable file
			defer fin.Close()
			return boardFromFile(fin)
		}
	}

	// Try to read a board from stdin
	return boardFromFile(os.Stdin)
}

func boardFromFile(fin *os.File) (*Board, error) {
	var boardBuffer []byte
	buf := make([]byte, 1024)

	var n int
	var err error
	for n, err = fin.Read(buf); n > 0 && err == nil; n, err = fin.Read(buf) {
		boardBuffer = append(boardBuffer, buf[:n]...)
	}

	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}

	lines := bytes.Split(boardBuffer, []byte{'\n'})
	if len(lines) > 8 {
		// almost certainly a 9-line by 9-character file
		return ReadBoard(bytes.NewBuffer(boardBuffer))
	}
	// go through boardBuffer, try to construct an 81-character SDM format string
	var sdmBuffer []byte
	for _, line := range lines {
		line = bytes.TrimSpace(line)
		if line[0] == '#' {
			continue
		}
		sdmBuffer = append(sdmBuffer, line...)
		if len(sdmBuffer) >= 81 {
			bd, err := NewBoardFromString(string(sdmBuffer))
			if err != nil {
				return nil, err
			}
			return bd, nil
		}
	}
	return nil, errors.New("didn't find a board")
}
