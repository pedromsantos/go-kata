# Tic Tac Toe Kata

## Rules

- The game is played on a grid that's 3 squares by 3 squares
- Players alternate placing X's and O's in empty squares
- X always plays first
- Players cannot play on a played square
- A Player wins when it has three squares in a row:
  - Horizontally
  - Vertically
  - Diagonally
- If all nine squares are filled and neither player has won, the game is a draw

## Follow TDD Rules Strictly

1. ✅ Write production code only to pass a failing unit test
2. ✅ Write only enough of a unit test to make it fail
3. ✅ Write only enough production code to make the failing test pass

## First Run

Implement Tic Tac Toe as best as you can using TDD.

## Second Run: Object Calisthenics

Implement Tic Tac Toe strictly applying object calisthenics rules.

### Object Calisthenics Rules (in order of importance)

1. **Wrap all primitives and strings** - Wrap primitives in a type if it has behavior or is an important domain concept
2. **First class collections** - Wrap collections in a type if it has behavior
3. **One dot per line** - Law of Demeter: don't write `dog.Body.Tail.Wag()`, write `dog.ExpressHappiness()`
4. **No getters/setters** - Tell, don't ask
5. **No structs with more than two instance variables**
6. **Only one level of indentation per method**
7. **Don't use else**
8. **Don't abbreviate names**
9. **Keep all entities small** - 50 lines per file, 5 lines per method

## Great Habits

### Considerations When Writing a New Test

- Tests should test one thing only
- Create more specific tests to drive a more generic solution (triangulate)
- Give your tests meaningful names reflecting your business domain
- See the test fail for the right reason
- Organize tests in Arrange, Act, Assert blocks

### Considerations When Making a Failing Test Pass

- Write the simplest code to pass the test
- Use Transformation Priority Premise
- Consider using object calisthenics to drive design decisions

### Considerations After the Test Passes

- Use the Rule of Three to tackle duplication
- Refactor design constantly

## Resources

- [Object Calisthenics](https://williamdurand.fr/2013/06/03/object-calisthenics/)
- [TestDesiderata by Kent Beck](https://kentbeck.github.io/TestDesiderata)
