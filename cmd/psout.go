package main

import (
	"flag"
	"log"
	"os"
	"sudoku2/board"
)

func main() {
	var printPossiblePS bool
	flag.BoolVar(&printPossiblePS, "C", false, "print digit possibilities in PostScript output")
	postScriptFileName := flag.String("p", "", "PostScript output in this file")
	flag.Parse()

	fin := os.Stdin
	if flag.NArg() > 0 {
		var err error
		fin, err = os.Open(flag.Arg(0))
		defer fin.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	bd := board.ReadBoard(fin)

	if *postScriptFileName != "" {
		fout, err := os.Create(*postScriptFileName)
		if err != nil {
			log.Fatal(err)
		}
		(&bd).EmitPostScript(fout, *postScriptFileName, printPossiblePS)
	}
}
