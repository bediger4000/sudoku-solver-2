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

### Backtracking Solution

This is problem 14.2 in the [Daily Coding Problem book](),
in chapter 14, "Backtracking".

The book gives this example Sudoku problem:

```
2 5 _ _ 3 _ 9 _ 1
_ 1 _ _ _ 4 _ _ _

_ _ 5 2 _ _ _ _ _
_ _ _ _ 9 8 1 _ _
_ 4 _ _ _ 3 _ _ _
_ _ _ 3 6 _ _ 7 2
_ 7 _ _ _ _ _ _ 3
9 _ 3 _ _ _ 6 _ 4
```

My current program solves this without using Naked Pair or Hidden Pair,
much less needing a trial-and-error backtracking algorithm.
The example code given in *Daily Coding Problem* doesn't even erase
possible values based on another such value in the same row, column or block.

I added backtracking (-B flag), which will get attempted
after the usual methods all find 0 eliminations and/or solutions.

This puzzle can't be solved with basic eliminations, Naked Pair,
and Hidden Pair:

```
_ _ 3 9 4 1 _ _ _ 
_ 1 _ _ _ 6 _ 9 _ 
_ 9 6 _ _ 5 1 _ 3 
6 2 8 4 _ 7 3 _ 9 
1 7 _ _ _ _ _ _ 4 
3 _ _ _ _ _ 7 2 _ 
9 _ 7 8 _ _ 5 _ 2 
_ 6 _ 5 _ _ _ 3 7 
_ _ _ _ 7 2 9 _ _ 
```

After doing eliminations the empty squares have these possiblities:

* <0,0> can hold [2 5 7 8]
* <0,1> can hold [5 8]
* <0,6> can hold [2 6 8]
* <0,7> can hold [5 6 7 8]
* <0,8> can hold [5 6 8]
* <1,0> can hold [2 4 5 7 8]
* <1,2> can hold [2 4 5]
* <1,3> can hold [2 3 7]
* <1,4> can hold [2 3 8]
* <1,6> can hold [2 4 8]
* <1,8> can hold [5 8]
* <2,0> can hold [2 4 7 8]
* <2,3> can hold [2 7]
* <2,4> can hold [2 8]
* <2,7> can hold [4 7 8]
* <3,4> can hold [1 5]
* <3,7> can hold [1 5]
* <4,2> can hold [5 9]
* <4,3> can hold [2 3 6]
* <4,4> can hold [2 3 5 6 8 9]
* <4,5> can hold [3 8 9]
* <4,6> can hold [6 8]
* <4,7> can hold [5 6 8]
* <5,1> can hold [4 5]
* <5,2> can hold [4 5 9]
* <5,3> can hold [1 6]
* <5,4> can hold [1 5 6 8 9]
* <5,5> can hold [8 9]
* <5,8> can hold [1 5 6 8]
* <6,1> can hold [3 4]
* <6,4> can hold [1 6]
* <6,5> can hold [3 4]
* <6,7> can hold [1 6]
* <7,0> can hold [2 4 8]
* <7,2> can hold [1 2 4]
* <7,4> can hold [1 9]
* <7,5> can hold [4 9]
* <7,6> can hold [4 8]
* <8,0> can hold [4 5 8]
* <8,1> can hold [3 4 5 8]
* <8,2> can hold [1 4 5]
* <8,3> can hold [1 3 6]
* <8,7> can hold [1 4 6 8]
* <8,8> can hold [1 6 8]

There's 44 unsolved cells.
It might take 2352735051982045184 attempts
I don't see backtracking as a very good way to solve Sudoku puzzles,
unless there's only a very few open squares.
That situation probably means that the puzzle in question doesn't have
a unique solution either.

I suspect that the interviewers who posed the "Solve Sudoku efficiently"
problem wanted some kind of constraint solver.
