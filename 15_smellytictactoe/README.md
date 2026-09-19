# Code Smells Kata - TicTacToe Refactoring

## Overview

This kata contains a deliberately "smelly" implementation of TicTacToe that needs refactoring. Your goal is to identify and fix various code smells while maintaining functionality.

The implementation is split across five files (`game.go`, `board.go`,
`row_winner_checker.go`, `column_winner_checker.go`,
`diagonal_winner_checker.go`) specifically so the cross-file smells below
(Shotgun Surgery, Duplicated Code, Divergent Change) are genuinely
cross-file, not just repeated logic within one file — this also makes
the kata a verification fixture for
[jev-review](https://github.com/pedromsantos/jev-review)'s module-level
rules. `RowWinnerChecker`, `ColumnWinnerChecker`, and
`DiagonalWinnerChecker` each independently re-implement the exact same
"are these three tiles taken and equal" matching logic (rows, columns,
and both diagonals are all checked); `Game` is the file that would need
editing for several unrelated reasons (move-validation rules, wiring in
a new line-checking strategy).

## Code Smells to Look For

The implementation contains the following code smells:

1. **Primitive Obsession**
   - Using primitive types (string for symbol) where custom types would be more appropriate

2. **Feature Envy**
   - `Game` accesses `Board`'s internals excessively

3. **Data Class**
   - `Tile` is just a data holder with no behavior

4. **Message Chain**
   - Long chains like `g.board.TileAt(x, y).Symbol` violate Law of Demeter

5. **Long Method/Function**
   - `Winner()` is too long and does multiple things

6. **Comments**
   - Excessive comments that could be replaced with clearer code

7. **Long Parameter List**
   - Functions with too many parameters

8. **Shotgun Surgery**
   - Changes requiring multiple small edits across many files — e.g. a change to
     the matching rule in `RowWinnerChecker`/`ColumnWinnerChecker`/`DiagonalWinnerChecker`
     needs editing all three files to stay consistent

9. **Duplicated Code**
   - The "are these three tiles taken and equal" matching logic is repeated,
     independently, across `RowWinnerChecker`, `ColumnWinnerChecker`, and
     `DiagonalWinnerChecker`

10. **Large Struct**
    - Structs with too many responsibilities

11. **Divergent Change**
    - `Game` changes for multiple unrelated reasons — move-validation rule
      changes, wiring in a new line-checking strategy, and orchestration
      changes

12. **Data Clump**
    - `x` and `y` coordinates always appear together

13. **Dead Code**
    - Unused code that should be removed

## Tasks

1. Review the code and identify all code smells
2. Add comments marking each code smell you find
3. Refactor the code using small, incremental steps
4. Ensure all tests remain passing after each refactoring

## Suggested Refactorings

### 1. Data Clumps → Position Value Object

```go
type Position struct {
    X, Y int
}
```

### 2. Primitive Obsession → Player Type

```go
type Player string

const (
    PlayerX Player = "X"
    PlayerO Player = "O"
    Empty   Player = " "
)
```

### 3. Long Method → Extract Functions

```go
func (g *Game) checkRow(row int) string
func (g *Game) checkWinningLine(positions []Position) string
```

### 4. Feature Envy → Move Logic to Board

Move winning condition logic to `Board` struct

### 5. Message Chains → Encapsulate Operations

```go
func (b *Board) SymbolAt(x, y int) string
func (b *Board) IsEmpty(x, y int) bool
```

## Tips

- Make one change at a time
- Run tests after each change: `go test ./...`
- Use Git to commit after each successful refactor
- Revert if tests fail

## Resources

- [Code Smells Video Tutorial](https://www.youtube.com/watch?v=MM6_tyvBRXE)
- [Refactoring Guru - Code Smells](https://refactoring.guru/refactoring/smells)
- [Comprehensive Code Smells Guide](https://luzkan.github.io/smells/)
- [TestDesiderata by Kent Beck](https://kentbeck.github.io/TestDesiderata)
