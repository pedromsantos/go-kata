# Leap Year Kata

Write a function that returns true or false depending on whether its input integer is a leap year or not.

## Rules

A leap year is defined as one that:

- is divisible by 4
- but is not otherwise divisible by 100
- unless it is also divisible by 400

## Examples

| **Input** | **Output** | **Reason**           |
| --------- | ---------- | -------------------- |
| **1**     | false      | not divisible by 4   |
| **4**     | true       | divisible by 4       |
| **100**   | false      | divisible by 100     |
| **400**   | true       | divisible by 400     |
| **2001**  | false      | typical common year  |
| **1996**  | true       | typical leap year    |
| **1900**  | false      | atypical common year |
| **2000**  | true       | atypical leap year   |

## Follow TDD Rules Strictly

1. ✅ Write production code only to pass a failing unit test
2. ✅ Write only enough of a unit test to make it fail
3. ✅ Write only enough production code to make the failing test pass

## Resources

- [TestDesiderata by Kent Beck](https://kentbeck.github.io/TestDesiderata)
