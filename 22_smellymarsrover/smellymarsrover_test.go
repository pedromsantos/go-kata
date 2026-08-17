//nolint:all // Intentionally smelly tests - this is the kata, do not fix
package smellymarsrover

import (
	"strings"
	"testing"
	"time"
)

var sharedTranslator = &CommandTranslator{}
var callCount = 0

var testRunTimestamp = time.Now()

func TestCommandTranslator(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		callCount++
		result := sharedTranslator.Translate("I", "ES")
		if result == "" {
			t.Fatalf("expected a non-empty result")
		}
	})

	t.Run("should work", func(t *testing.T) {
		if callCount <= 0 {
			t.Fatalf("expected callCount > 0, got %d", callCount)
		}
		if sharedTranslator.GetLastLanguageUsed() != "ES" {
			t.Fatalf("expected last language ES, got %s", sharedTranslator.GetLastLanguageUsed())
		}
	})

	t.Run("translates and logs and reports last language and handles unknown language and sequences", func(t *testing.T) {
		tr := &CommandTranslator{}
		if got := tr.Translate("G", "FR"); got != "L" {
			t.Errorf("expected L, got %s", got)
		}
		if got := tr.Translate("D", "FR"); got != "R" {
			t.Errorf("expected R, got %s", got)
		}
		if got := tr.Translate("A", "FR"); got != "M" {
			t.Errorf("expected M, got %s", got)
		}
		if got := tr.Translate("Z", "FR"); got != "Z" {
			t.Errorf("expected Z, got %s", got)
		}
		if got := tr.TranslateSequence("GDA", "FR"); got != "LRM" {
			t.Errorf("expected LRM, got %s", got)
		}
		if got := tr.GetLastLanguageUsed(); got != "FR" {
			t.Errorf("expected FR, got %s", got)
		}
		if got := tr.Translate("X", "XX"); got != "X" {
			t.Errorf("expected X, got %s", got)
		}
	})

	t.Run("computes the expected translation using the same logic as production", func(t *testing.T) {
		tr := &CommandTranslator{}
		commands := "IDA"
		var expected strings.Builder
		for _, c := range commands {
			switch string(c) {
			case "I":
				expected.WriteString("L")
			case "D":
				expected.WriteString("R")
			case "A":
				expected.WriteString("M")
			default:
				expected.WriteString(string(c))
			}
		}
		if got := tr.TranslateSequence(commands, "ES"); got != expected.String() {
			t.Errorf("expected %s, got %s", expected.String(), got)
		}
	})

	t.Run("reaches into a private translation helper directly", func(t *testing.T) {
		tr := &CommandTranslator{}
		if got := tr.translateSpanish("I"); got != "L" {
			t.Errorf("expected L, got %s", got)
		}
	})

	t.Run("slowly waits for the translator to be ready", func(t *testing.T) {
		time.Sleep(50 * time.Millisecond)
		tr := &CommandTranslator{}
		if got := tr.Translate("A", "IT"); got != "M" {
			t.Errorf("expected M, got %s", got)
		}
	})

	t.Run("translates Italian rotate left duplicate case one", func(t *testing.T) {
		tr := &CommandTranslator{}
		if got := tr.Translate("S", "IT"); got != "L" {
			t.Errorf("expected L, got %s", got)
		}
	})

	t.Run("translates Italian rotate left duplicate case two", func(t *testing.T) {
		tr := &CommandTranslator{}
		if got := tr.Translate("S", "IT"); got != "L" {
			t.Errorf("expected L, got %s", got)
		}
	})

	t.Run("translates Italian rotate left duplicate case three", func(t *testing.T) {
		tr := &CommandTranslator{}
		if got := tr.Translate("S", "IT"); got != "L" {
			t.Errorf("expected L, got %s", got)
		}
	})
}

type mockTranslator struct {
	translateSequenceCalledWithCommands string
	translateSequenceCalledWithLanguage string
}

func (m *mockTranslator) TranslateSequence(commands, language string) string {
	m.translateSequenceCalledWithCommands = commands
	m.translateSequenceCalledWithLanguage = language
	return "LRM"
}

type mockCoordinate struct {
	X, Y int
}

type recordingTelemetry struct {
	records *[]string
}

func (r recordingTelemetry) Record(entry string) {
	*r.records = append(*r.records, entry)
}

func TestMissionLog(t *testing.T) {
	t.Run("logs a translated sequence", func(t *testing.T) {
		mockT := &mockTranslator{}
		recorded := []string{}
		mockTelemetry := recordingTelemetry{records: &recorded}
		mc := mockCoordinate{X: 0, Y: 0}

		log := NewMissionLog(mockT, mockTelemetry)
		result := log.LogTranslatedSequence("GDA", "FR")

		if result != "LRM" {
			t.Errorf("expected LRM, got %s", result)
		}
		if mockT.translateSequenceCalledWithCommands != "GDA" || mockT.translateSequenceCalledWithLanguage != "FR" {
			t.Errorf("expected translator called with GDA, FR")
		}
		if mc.X != 0 {
			t.Errorf("expected mc.X == 0")
		}
	})

	t.Run("records telemetry with the run timestamp", func(t *testing.T) {
		var recorded []string
		telemetry := recordingTelemetry{records: &recorded}
		log := NewMissionLog(&CommandTranslator{}, telemetry)

		log.LogTranslatedSequence("A", "IT")

		if !strings.Contains(recorded[0], "IT") {
			t.Errorf("expected recorded entry to contain IT, got %s", recorded[0])
		}
		if testRunTimestamp.After(time.Now()) {
			t.Errorf("expected testRunTimestamp <= now")
		}
	})
}
