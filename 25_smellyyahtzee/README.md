# Smelly Yahtzee

This kata is a **fixture, not a from-scratch TDD exercise** — the same kind of
legacy-code/test-smell exercise as `22_smellymarsrover`, but scoped down to
dice rolling: `Die`, `DiceCup` (with an injectable random source), and a
`TurnLog` that writes to a `TelemetryPort`. `smellyyahtzee.go` and
`smellyyahtzee_test.go` contain no explanatory comments; finding the smells is
part of the exercise.

## Goals

- Make dice rolling deterministic and testable.
- Characterize rolling, selecting dice, and rerolling selected dice.
- Improve the test suite so it protects observable game behavior.
- Keep tests fast, isolated, readable, and mutation-resistant.
- Refactor in small behavior-preserving steps.

## Run

```sh
go test -v ./25_smellyyahtzee/...
```
