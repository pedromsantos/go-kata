# Code Smells Kata - TicTacToe Refactoring

## Overview

This kata contains a deliberately "smelly" implementation of TicTacToe that needs refactoring. Your goal is to identify and fix various code smells while maintaining functionality.

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

8. **Duplicated Code**
   - Row-checking logic is repeated three times

9. **Large Struct**
   - Structs with too many responsibilities

10. **Data Clump**
    - `x` and `y` coordinates always appear together

11. **Dead Code**
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
