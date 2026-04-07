# Stack Kata

## Requirements

Implement a Stack with the following methods:

- `Push(value int)` - Adds a value to the top of the stack
- `Pop() (int, error)` - Removes and returns the top value from the stack
- The stack should return an error when attempting to pop from an empty stack

## Test Cases

| Input (Push)  | Expected Output (Pop) |
| ------------- | --------------------- |
| Empty Stack   | error                 |
| Push(1)       | 1                     |
| Push(2)       | 2                     |
| Push(1, 2)    | 2, then 1             |
| Push(1, 2, 3) | 3, then 2, then 1     |

## TDD Rules

1. ✅ Write production code only to pass a failing unit test
2. ✅ Write only enough of a unit test to make it fail
3. ✅ Write only enough production code to make the failing test pass

## Additional Guidelines

- Follow the "Red-Green-Refactor" cycle
- Keep tests and production code as simple as possible
- Commit after each successful test

## Resources

- [TestDesiderata by Kent Beck](https://kentbeck.github.io/TestDesiderata)
