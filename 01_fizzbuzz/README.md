# FizzBuzz Kata

Write a function that takes a number and returns it as a string.

- For multiples of three return "fizz" instead of the stringified number.
- For multiples of five return "buzz" instead of the stringified number.
- For numbers which are multiples of both three and five return "fizzbuzz" instead of the stringified number.

## Examples

| **Input** | **Output** |
| --------- | ---------- |
| **1**     | "1"        |
| **2**     | "2"        |
| **3**     | "fizz"     |
| **4**     | "4"        |
| **5**     | "buzz"     |
| **6**     | "fizz"     |
| **7**     | "7"        |
| **8**     | "8"        |
| **9**     | "fizz"     |
| **10**    | "buzz"     |
| **11**    | "11"       |
| **12**    | "fizz"     |
| **13**    | "13"       |
| **14**    | "14"       |
| **15**    | "fizzbuzz" |

## Follow TDD Rules Strictly

1. ✅ Write production code only to pass a failing unit test
2. ✅ Write only enough of a unit test to make it fail
3. ✅ Write only enough production code to make the failing test pass

## Great Habits

### Considerations When Writing a New Test

- Tests should **test one thing only.**
- Create more specific tests to drive a more generic solution (**triangulate**) by adding new tests that force your code to pivot.
- Give your tests meaningful names (behavior/goal-oriented) expressing your business domain.
- Always see the test fail for the right reason.
- Ensure you have meaningful feedback from the failing test.
- Keep your tests and production code separate.
- Organize your unit tests to reflect your production code (similar project structure).

### Considerations When Making a Failing Test Pass

- Write the simplest code to pass the test.
- Fake it or use obvious implementation.
- It is okay to write any code that makes you get to the refactor phase more quickly.

### Considerations After the Test Passes

- Use the Rule of Three to tackle duplication.

## Resources

- [TestDesiderata by Kent Beck](https://kentbeck.github.io/TestDesiderata)
