//nolint:all // This is intentionally legacy/smelly code for characterization-testing practice - do not fix
package smellyyahtzee

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
)

type Die struct {
	Value int
}

func NewDie(value int) Die {
	return Die{Value: value}
}

func (d Die) Equals(other Die) bool {
	return d.Value == other.Value
}

type TelemetryPort interface {
	Record(entry string)
}

type DiceCup struct {
	dice            []Die
	selectedIndexes []int
	randomSource    func() float64
}

func NewDiceCup(randomSource func() float64) *DiceCup {
	if randomSource == nil {
		randomSource = rand.Float64
	}
	return &DiceCup{randomSource: randomSource}
}

func (c *DiceCup) Roll() []Die {
	c.dice = make([]Die, 5)
	for i := range c.dice {
		c.dice[i] = c.rollDie()
	}
	c.selectedIndexes = nil
	return c.dice
}

func (c *DiceCup) SelectForReroll(indexes []int) {
	c.selectedIndexes = append([]int{}, indexes...)
}

func (c *DiceCup) RerollSelected() []Die {
	for _, index := range c.selectedIndexes {
		c.dice[index] = c.rollDie()
	}
	c.selectedIndexes = nil
	return c.dice
}

func (c *DiceCup) CurrentDice() []Die {
	return c.dice
}

func (c *DiceCup) rollDie() Die {
	return Die{Value: int(math.Floor(c.randomSource()*6)) + 1}
}

// Cup is the narrow interface TurnLog depends on, satisfied implicitly by
// *DiceCup.
type Cup interface {
	RerollSelected() []Die
}

type TurnLog struct {
	diceCup   Cup
	telemetry TelemetryPort
}

func NewTurnLog(diceCup Cup, telemetry TelemetryPort) *TurnLog {
	return &TurnLog{diceCup: diceCup, telemetry: telemetry}
}

func (t *TurnLog) RerollSelectedDice() []Die {
	dice := t.diceCup.RerollSelected()
	values := make([]string, len(dice))
	for i, die := range dice {
		values[i] = strconv.Itoa(die.Value)
	}
	t.telemetry.Record("rerolled:" + strings.Join(values, ","))
	return dice
}
