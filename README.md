# Second Mediocre Sudoku Solver

My [other sudoku solver](https://github.com/bediger4000/mediocre-sudoku-solver)
ended up too hard to understand after an interval of neglect.
Here's my second attempt.
It gets exactly the same answers as the first solver.

## Building

    $ cd $GOPATH/src
    $ git clone https://github.com/bediger4000/sudoku-solver-2.git ./sudoku2
    $ cd sudoku2
    $ go build sudoku2

## Running

```sh
$ ./sudoku2 -H -N -a input.file
```

You have to turn on naked pair elimination (-N) and hidden pair elimination (-H).
The "-a" flag announces eliminations and cell solutions.
Almost all puzzles rated "medium" or harder need both extra eliminations.
