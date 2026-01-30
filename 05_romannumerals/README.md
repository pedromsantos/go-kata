# Roman Numerals Kata

## First Run

Implement a Roman numeral converter. The code must be able to take numbers up to 3999 and convert to their Roman equivalent.

## Examples

| **Arabic** | **Roman** |
| ---------- | --------- |
| 1          | I         |
| 2          | II        |
| 3          | III       |
| 4          | IV        |
| 5          | V         |
| 6          | VI        |
| 7          | VII       |
| 8          | VIII      |
| 9          | IX        |
| 10         | X         |
| 40         | XL        |
| 50         | L         |
| 90         | XC        |
| 100        | C         |
| 400        | CD        |
| 500        | D         |
| 900        | CM        |
| 1000       | M         |

## Follow TDD Rules Strictly

1. ✅ Write production code only to pass a failing unit test
2. ✅ Write only enough of a unit test to make it fail
3. ✅ Write only enough production code to make the failing test pass

## Second Run: Transformation Priority Premise

Use the Transformation Priority Premise (TPP) to evolve your code.

### TPP Table

| #   | Transformation              | Description               |
| --- | --------------------------- | ------------------------- |
| 1   | {} → nil                    | Return nil/zero value     |
| 2   | nil → constant              | Return a constant         |
| 3   | constant → constant+        | Concatenate constants     |
| 4   | constant → scalar           | Use argument directly     |
| 5   | statement → statements      | Add more statements       |
| 6   | unconditional → conditional | Add if/else               |
| 7   | scalar → array              | Use array/slice           |
| 8   | array → container           | Use map                   |
| 9   | statement → tail recursion  | Add recursive call at end |
| 10  | if → loop                   | Replace if with for       |
| 11  | statement → recursion       | Full recursion            |
| 12  | expression → function       | Extract to function       |
| 13  | variable → mutation         | Mutate variable           |

Prefer simpler transformations (top) over complex ones (bottom).

## Resources

- [TestDesiderata by Kent Beck](https://kentbeck.github.io/TestDesiderata)
- [Transformation Priority Premise](https://blog.cleancoder.com/uncle-bob/2013/05/27/TheTransformationPriorityPremise.html)
