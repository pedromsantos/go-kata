# Character Copier Kata

## Source

<https://www.planetgeek.ch/2010/03/31/mocking-kata-copier-net/>

## Problem Description

The Character Copier is a simple struct that reads characters from a source and copies them to a destination one character at a time.

### Behavior

- When the `Copy` method is called, it reads characters from the source and copies them to the destination
- The copying continues until a newline character (`\n`) is encountered
- The newline character should not be copied to the destination

### Implementation Requirements

- Implement the character copier using Test Doubles for the source and destination
- Explore different testing approaches:
  - Using Spies (manually written mocks)
  - Using a mocking framework like [gomock](https://github.com/golang/mock)

### Example Input/Output

| Input (GetChar) | Output (SetChar) | Notes                 |
| --------------- | ---------------- | --------------------- |
| 'a'             | 'a'              | Copy single character |
| 'b'             | 'b'              | Copy single character |
| '\n'            | N/A              | Stop on newline       |

### Interfaces

```go
type Source interface {
    GetChar() rune
}

type Destination interface {
    SetChar(c rune)
}
```

## Follow TDD Rules Strictly

1. ✅ Write production code only to pass a failing unit test
2. ✅ Write only enough of a unit test to make it fail
3. ✅ Write only enough production code to make the failing test pass

## Resources

- [gomock](https://github.com/golang/mock)
- [TestDesiderata by Kent Beck](https://kentbeck.github.io/TestDesiderata)
