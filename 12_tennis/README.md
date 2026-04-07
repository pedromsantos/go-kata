# Tennis Kata

## Source

<https://github.com/emilybache/Tennis-Refactoring-Kata>

## Problem Statement

You are taking over a tennis scoring system project from a colleague who:

- Has spent 8.5 hours of a 10-hour billable project
- Claims the work is complete with passing tests
- Is currently unavailable

Your manager has requested you to:

1. Spend the remaining billable time (~1 hour) improving the code
2. Prepare feedback on the current design
3. Be ready to discuss the value of refactoring beyond billable hours

## Refactoring Task

Your goal is to improve code readability while maintaining functionality.

### Prerequisites

- Verify all tests pass: `go test ./...`
- Check coverage: `go test -cover ./...`

### Refactoring Guidelines

#### Best Practices

1. **Maintain Working Code**
   - Run tests after each change
   - Verify coverage remains high
   - Commit frequently
   - Use `git checkout` if tests fail

2. **Improve Readability**
   - Remove Clutter: Standardize formatting, use clear names
   - Clean Documentation: Remove redundant comments
   - Make Knowledge Explicit: Replace magic numbers with constants
   - Optimize Structure: Refine variable scope

3. **Reduce Complexity**
   - Simplify Code: Break down long methods
   - Eliminate Duplication: Extract common patterns
   - Minimize Nesting: Exit methods early when possible

## Scoring Rules Reference

| Points | Call    |
| ------ | ------- |
| 0      | Love    |
| 1      | Fifteen |
| 2      | Thirty  |
| 3      | Forty   |

- When tied: "X-All" (e.g., "Fifteen-All")
- When tied at 3+: "Deuce"
- When one player is ahead by 1 at 3+: "Advantage <player>"
- When one player is ahead by 2+ at 3+: "Win for <player>"

## Resources

- [Abracadabra VSCode Extension](https://github.com/nicoespeon/abracadabra)
- [TestDesiderata by Kent Beck](https://kentbeck.github.io/TestDesiderata)
