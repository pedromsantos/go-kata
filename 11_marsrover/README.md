# Mars Rover Kata

## Problem Description

A NASA robot rover has landed on Mars. The rover must navigate a rectangular plateau divided into a grid, capturing terrain images with its on-board cameras to send back to Earth.

## Input Specification

1. **Plateau Size**: Upper-right coordinates (e.g., `5 5`)
   - Lower-left coordinates are always `0 0`

2. **Initial Position**: Starting position and direction (e.g., `1 2 N`)
   - Direction can be N(orth), E(ast), S(outh), or W(est)
   - Moving North from (x, y) leads to (x, y+1)

3. **Command Sequence**: Series of movement instructions (e.g., `LMLMLMLMM`)
   - `L`: Rotate 90° left
   - `R`: Rotate 90° right
   - `M`: Move forward one grid point

## Test Cases

### Test Case 1

Input:

```
5 5
1 2 N
LMLMLMLMM
```

Output: `1 3 N`

### Test Case 2

Input:

```
5 5
3 3 E
MMRMMRMRRM
```

Output: `5 1 E`

## Implementation Guidelines

### First Implementation

1. Use Test-Driven Development (TDD)
2. Start with the provided acceptance tests

### Extra: Obstacles and Wrap Around

- Implement wrapping at edges
- Implement obstacle detection before each move
- If a command encounters an obstacle, rover moves up to the last possible point and reports the obstacle

Example output with obstacle: `O 1 3 N`

### Refactoring Phase

Consider using design patterns:

- Command Pattern
- State Pattern
- Strategy Pattern

## Follow TDD Rules Strictly

1. ✅ Write production code only to pass a failing unit test
2. ✅ Write only enough of a unit test to make it fail
3. ✅ Write only enough production code to make the failing test pass

## Resources

- [TestDesiderata by Kent Beck](https://kentbeck.github.io/TestDesiderata)
