# Yahtzee Kata

## Rules

The game of Yahtzee is a simple dice game. Each round, each player rolls five six-sided dice. The player may choose to reroll some or all of the dice up to three times (including the original roll).

The player then places the roll at a category, such as ones, twos, sixes, pair, two pairs, etc. If the roll is compatible with the category, the player gets a score according to the rules. If not compatible, the player gets zero.

## Categories & Scoring Rules

### Number Categories (Ones through Sixes)

- Score: Sum of all dice matching the category number
- Examples:
  - Roll [1,1,2,4,4] on "Fours" scores 8 points
  - Roll [2,2,2,5,6] on "Twos" scores 6 points

### Combinations

| Category       | Rule                             | Example          | Score |
| -------------- | -------------------------------- | ---------------- | ----- |
| Pair           | Sum of two highest matching dice | [3,3,3,4,4] → 4s | 8     |
| Two Pairs      | Sum of all dice in both pairs    | [1,1,3,3,2]      | 8     |
| Three of Kind  | Sum of the three matching dice   | [3,3,3,4,5]      | 9     |
| Four of Kind   | Sum of the four matching dice    | [2,2,2,2,5]      | 8     |
| Small Straight | Exactly [1,2,3,4,5]              | [1,2,3,4,5]      | 15    |
| Large Straight | Exactly [2,3,4,5,6]              | [2,3,4,5,6]      | 20    |
| Full House     | Three of one + two of another    | [2,2,2,3,3]      | 25    |
| Yahtzee        | All five dice same               | [4,4,4,4,4]      | 50    |

## Follow TDD Rules Strictly

1. ✅ Write production code only to pass a failing unit test
2. ✅ Write only enough of a unit test to make it fail
3. ✅ Write only enough production code to make the failing test pass

## Object Calisthenics

Consider applying these rules:

- Wrap all primitives and strings
- First class collections
- One dot per line (Law of Demeter)
- No getters/setters
- Only one level of indentation per method
- Don't use else
- Don't abbreviate names

## Resources

- [Object Calisthenics](https://williamdurand.fr/2013/06/03/object-calisthenics/)
- [TestDesiderata by Kent Beck](https://kentbeck.github.io/TestDesiderata)
