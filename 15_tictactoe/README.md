# Tic Tac Toe Kata (Output Version)

## Rules

- X always plays first
- Players alternate placing X's and O's on the board
- Players cannot play on a played position
- A Player wins when it has three in a row:
  - Horizontally
  - Vertically
  - Diagonally
- If all nine squares are filled and neither player has won, the game is a draw

## Special Requirements

In this version of TicTacToe, nothing is returned but a call to an Output interface is made to print game events.

## Output Interface

```go
type Output interface {
    PlayerXPlayed(position int)
    PlayerOPlayed(position int)
    PlayerXWon()
    PlayerOWon()
    Draw()
    InvalidMove(position int)
}
```

## Implementation Approach

1. Use Test Doubles (Spies) to verify Output calls
2. Focus on behavior verification rather than state verification
3. Apply TDD strictly

## Follow TDD Rules Strictly

1. ✅ Write production code only to pass a failing unit test
2. ✅ Write only enough of a unit test to make it fail
3. ✅ Write only enough production code to make the failing test pass

## Resources

- [TestDesiderata by Kent Beck](https://kentbeck.github.io/TestDesiderata)
