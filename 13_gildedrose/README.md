# Gilded Rose Kata

## Source

- <https://github.com/NotMyself/GildedRose>
- <https://github.com/emilybache/GildedRose-Refactoring-Kata>

## The Problem

Hi and welcome to team Gilded Rose. As you know, we are a small inn with a prime location in a prominent city ran by a friendly innkeeper named Allison. We also buy and sell only the finest goods. Unfortunately, our goods are constantly degrading in quality as they approach their sell by date. We have a system in place that updates our inventory for us. It was developed by a no-nonsense type named Leeroy, who has moved on to new adventures.

## Tasks Overview

### 1. Characterization Tests

Goal: document the system as it is. Grow tests along a path of increasing generality: start with unit tests, expand them with parameterized cases, and finish with property-based checks. Keep code coverage on to spot missing scenarios throughout.

#### 1.1 Unit Tests

- Describe current behavior with focused assertions (actual behavior over intended behavior)
- Test-driven loop:
  1. Write a failing assertion
  2. Run to observe actual behavior
  3. Update expectation to match reality
  4. Repeat and watch coverage to find gaps

#### 1.2 Table-Driven Tests

- Generalize unit tests by enumerating meaningful input sets
- Use coverage reports to see which combinations still need attention

```go
testCases := []struct {
    name            string
    inputName       string
    inputSellIn     int
    inputQuality    int
    expectedQuality int
}{
    {"normal item", "foo", 10, 20, 19},
    {"aged brie increases", "Aged Brie", 10, 20, 21},
}

for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
        items := []*Item{{Name: tc.inputName, SellIn: tc.inputSellIn, Quality: tc.inputQuality}}
        UpdateQuality(items)
        assert.Equal(t, tc.expectedQuality, items[0].Quality)
    })
}
```

#### 1.3 Property-Based Tests

- Capture broader invariants with random or generated data
- Uses [rapid](https://github.com/flyingmutant/rapid) for property-based testing

```go
import "pgregory.net/rapid"

t.Run("quality is never negative", func(t *testing.T) {
    rapid.Check(t, func(t *rapid.T) {
        name := rapid.String().Draw(t, "name")
        sellIn := rapid.IntRange(-10, 50).Draw(t, "sellIn")
        quality := rapid.IntRange(0, 100).Draw(t, "quality")

        items := []*Item{{Name: name, SellIn: sellIn, Quality: quality}}
        UpdateQuality(items)

        assert.GreaterOrEqual(t, items[0].Quality, 0)
    })
})
```

### 2. Golden Master Tests (Approval/Snapshot Testing)

- Capture the full current output and lock it in `.approved.txt` files
- Coverage is critical: ensure snapshots exercise all code paths
- Uses [go-approval-tests](https://github.com/approvals/go-approval-tests)

```go
import approvals "github.com/approvals/go-approval-tests"

func TestApproval(t *testing.T) {
    var result strings.Builder
    names := []string{"foo", "Aged Brie", "Backstage passes to a TAFKAL80ETC concert"}

    for _, name := range names {
        for sellIn := -1; sellIn <= 11; sellIn++ {
            for quality := 0; quality <= 50; quality++ {
                items := []*Item{{Name: name, SellIn: sellIn, Quality: quality}}
                UpdateQuality(items)
                result.WriteString(fmt.Sprintf("%s,%d,%d => %d,%d\n",
                    name, sellIn, quality, items[0].SellIn, items[0].Quality))
            }
        }
    }
    approvals.VerifyString(t, result.String())
}
```

### Characterization vs Golden Master

| Aspect         | Characterization tests                         | Golden Master tests                                 |
| -------------- | ---------------------------------------------- | --------------------------------------------------- |
| Purpose        | Document current behavior                      | Freeze full outputs to detect regressions wholesale |
| Granularity    | Fine-grained assertions on specific behaviors  | Coarse snapshots of overall results                 |
| Evolution      | Easy to adjust per case as understanding grows | Stable baseline; snapshots updated deliberately     |
| Coverage role  | Highlights missing cases as you generalize     | Ensures snapshots actually exercise critical paths  |
| Failure signal | Points to a specific scenario                  | Signals any change; inspect diff to diagnose        |

### 3. Mutation Testing

#### Code Coverage vs Test Effectiveness

- **Code coverage** answers "which lines executed?" — a line that runs is not necessarily tested
- **Test effectiveness** answers "do assertions actually catch bugs?" — mutation testing measures this
- High coverage with weak assertions gives false confidence; mutants expose those gaps

#### Running Mutation Tests

```bash
go install github.com/zimmski/go-mutesting/cmd/go-mutesting@latest
go-mutesting ./11_gildedrose/...
```

### 4. Refactoring Code

Refactor using the "Lift-Up Conditional" technique:

1. Extract code to refactor into new function
2. Identify and copy condition to lift-up
3. Make the condition positive
4. Duplicate the body in both branches
5. Use coverage to delete dead code
6. Repeat

### 5. Refactoring Design

Refactor using "Replace Conditional with Polymorphism":

1. Replace Type Code with Interfaces
2. Use the factory pattern and the strategy pattern

## Resources

- [Refactoring Guru: Replace Conditional with Polymorphism](https://refactoring.guru/replace-conditional-with-polymorphism)
- [TestDesiderata by Kent Beck](https://kentbeck.github.io/TestDesiderata)
