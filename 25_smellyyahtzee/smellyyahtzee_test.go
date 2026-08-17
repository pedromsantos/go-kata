//nolint:all // Intentionally smelly tests - this is the kata, do not fix
package smellyyahtzee

import (
	"reflect"
	"testing"
	"time"
)

var sharedCup = NewDiceCup(func() float64 { return 0 })
var rollCount = 0

var testRun = time.Now()

func TestYahtzeeDiceSelection(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		rollCount++
		dice := sharedCup.Roll()
		if dice == nil {
			t.Fatalf("expected dice to be defined")
		}
	})

	t.Run("should work", func(t *testing.T) {
		if rollCount <= 0 {
			t.Fatalf("expected rollCount > 0, got %d", rollCount)
		}
		if len(sharedCup.CurrentDice()) != 5 {
			t.Fatalf("expected 5 dice, got %d", len(sharedCup.CurrentDice()))
		}
	})

	t.Run("rolls dice and selects dice and rerolls dice and clears selection", func(t *testing.T) {
		cup := NewDiceCup(func() float64 { return 0.5 })
		rolled := cup.Roll()
		cup.SelectForReroll([]int{0, 2})
		rerolled := cup.RerollSelected()

		if len(rolled) != 5 {
			t.Errorf("expected 5 dice, got %d", len(rolled))
		}
		if rolled[0].Value != 4 {
			t.Errorf("expected 4, got %d", rolled[0].Value)
		}
		if rolled[1].Value != 4 {
			t.Errorf("expected 4, got %d", rolled[1].Value)
		}
		if rerolled[2].Value != 4 {
			t.Errorf("expected 4, got %d", rerolled[2].Value)
		}
		if len(cup.CurrentDice()) != 5 {
			t.Errorf("expected 5 dice, got %d", len(cup.CurrentDice()))
		}
		if reflect.ValueOf(rerolled).Pointer() != reflect.ValueOf(cup.CurrentDice()).Pointer() {
			t.Errorf("expected rerolled to be the same slice as cup.CurrentDice()")
		}
	})

	t.Run("does things", func(t *testing.T) {
		dice := NewDiceCup(func() float64 { return 0 }).Roll()

		if len(dice) != 5 {
			t.Errorf("expected 5 dice, got %d", len(dice))
		}
		if dice[0].Value != 1 || dice[1].Value != 1 || dice[2].Value != 1 || dice[3].Value != 1 {
			t.Errorf("expected dice 0-3 to be 1")
		}
	})

	t.Run("computes expected dice with the same branching as the cup", func(t *testing.T) {
		values := []float64{0.01, 0.2, 0.4, 0.7, 0.99}
		next := 0
		cup := NewDiceCup(func() float64 {
			v := values[next]
			next++
			return v
		})

		var expected []int
		for _, value := range []float64{0.01, 0.2, 0.4, 0.7, 0.99} {
			switch {
			case value < 1.0/6:
				expected = append(expected, 1)
			case value < 2.0/6:
				expected = append(expected, 2)
			case value < 3.0/6:
				expected = append(expected, 3)
			case value < 5.0/6:
				expected = append(expected, 5)
			default:
				expected = append(expected, 6)
			}
		}

		rolled := cup.Roll()
		got := make([]int, len(rolled))
		for i, die := range rolled {
			got[i] = die.Value
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("reaches into the private die roller", func(t *testing.T) {
		cup := NewDiceCup(func() float64 { return 0 })
		die := cup.rollDie()
		if die.Value != 1 {
			t.Errorf("expected 1, got %d", die.Value)
		}
	})

	t.Run("slowly waits before rolling", func(t *testing.T) {
		time.Sleep(20 * time.Millisecond)
		if got := NewDiceCup(func() float64 { return 0 }).Roll()[0].Value; got != 1 {
			t.Errorf("expected 1, got %d", got)
		}
	})

	t.Run("rerolls the first die duplicate case one", func(t *testing.T) {
		cup := NewDiceCup(func() float64 { return 0 })
		cup.Roll()
		cup.SelectForReroll([]int{0})
		if got := cup.RerollSelected()[0].Value; got != 1 {
			t.Errorf("expected 1, got %d", got)
		}
	})

	t.Run("rerolls the first die duplicate case two", func(t *testing.T) {
		cup := NewDiceCup(func() float64 { return 0 })
		cup.Roll()
		cup.SelectForReroll([]int{0})
		if got := cup.RerollSelected()[0].Value; got != 1 {
			t.Errorf("expected 1, got %d", got)
		}
	})

	t.Run("rerolls the first die duplicate case three", func(t *testing.T) {
		cup := NewDiceCup(func() float64 { return 0 })
		cup.Roll()
		cup.SelectForReroll([]int{0})
		if got := cup.RerollSelected()[0].Value; got != 1 {
			t.Errorf("expected 1, got %d", got)
		}
	})
}

type mockCup struct {
	rerollSelectedCalls int
}

func (m *mockCup) RerollSelected() []Die {
	m.rerollSelectedCalls++
	return []Die{{Value: 1}, {Value: 2}}
}

type recordingTelemetry struct {
	records *[]string
}

func (r recordingTelemetry) Record(entry string) {
	*r.records = append(*r.records, entry)
}

func TestTurnLog(t *testing.T) {
	t.Run("logs rerolled dice", func(t *testing.T) {
		cup := &mockCup{}
		var recorded []string
		telemetry := recordingTelemetry{records: &recorded}
		mockDie := Die{Value: 6}

		log := NewTurnLog(cup, telemetry)
		dice := log.RerollSelectedDice()

		values := make([]int, len(dice))
		for i, die := range dice {
			values[i] = die.Value
		}
		if !reflect.DeepEqual(values, []int{1, 2}) {
			t.Errorf("expected [1 2], got %v", values)
		}
		if len(recorded) != 1 || recorded[0] != "rerolled:1,2" {
			t.Errorf("expected telemetry to record rerolled:1,2, got %v", recorded)
		}
		if cup.rerollSelectedCalls != 1 {
			t.Errorf("expected RerollSelected called once, got %d", cup.rerollSelectedCalls)
		}
		if mockDie.Value != 6 {
			t.Errorf("expected mockDie.Value == 6")
		}
	})

	t.Run("records a timestamp that is always in the past", func(t *testing.T) {
		if testRun.After(time.Now()) {
			t.Errorf("expected testRun <= now")
		}
	})
}
