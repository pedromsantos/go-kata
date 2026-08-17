//nolint:all // This is intentionally legacy/smelly code for characterization-testing practice - do not fix
package smellymarsrover

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func missionClockNow() string {
	return time.Now().Format(time.RFC3339)
}

type RadioTransmitter struct {
	endpoint string
}

func NewRadioTransmitter(endpoint string) *RadioTransmitter {
	return &RadioTransmitter{endpoint: endpoint}
}

func (r *RadioTransmitter) Send(message string) {
	fmt.Printf("[RADIO -> %s] %s\n", r.endpoint, message)
}

type ObstacleSensor struct {
	obstacles map[string]bool
}

func NewObstacleSensor() *ObstacleSensor {
	return &ObstacleSensor{
		obstacles: map[string]bool{"1,2": true, "3,3": true, "0,4": true},
	}
}

func (s *ObstacleSensor) DetectsObstacleAt(x, y int) bool {
	sensorNoise := rand.Float64() < 0.0001
	return sensorNoise || s.obstacles[fmt.Sprintf("%d,%d", x, y)]
}

type Position struct {
	X int
	Y int
}

type Rover struct {
	x         int
	y         int
	direction string
	gridSize  int
	sensor    *ObstacleSensor
	radio     *RadioTransmitter
}

func NewRover(x, y int, direction string, gridSize int) *Rover {
	return &Rover{
		x:         x,
		y:         y,
		direction: direction,
		gridSize:  gridSize,
		sensor:    NewObstacleSensor(),
		radio:     NewRadioTransmitter("mission-control.nasa.gov"),
	}
}

func (r *Rover) Execute(commands string) string {
	for _, command := range commands {
		previous := Position{X: r.x, Y: r.y}

		switch command {
		case 'L':
			r.turnLeft()
		case 'R':
			r.turnRight()
		case 'M':
			r.moveForward()
		}

		if r.sensor.DetectsObstacleAt(r.x, r.y) {
			r.x = previous.X
			r.y = previous.Y
			r.reportObstacle()
			return fmt.Sprintf("O %d %d %s", r.x, r.y, r.direction)
		}
	}

	return fmt.Sprintf("%d %d %s", r.x, r.y, r.direction)
}

func (r *Rover) reportObstacle() {
	r.radio.Send(fmt.Sprintf("OBSTACLE %d %d %s at %s", r.x, r.y, r.direction, missionClockNow()))
}

func (r *Rover) turnLeft() {
	order := []string{"N", "W", "S", "E"}
	r.direction = order[(indexOf(order, r.direction)+1)%4]
}

func (r *Rover) turnRight() {
	order := []string{"N", "E", "S", "W"}
	r.direction = order[(indexOf(order, r.direction)+1)%4]
}

func (r *Rover) moveForward() {
	switch r.direction {
	case "N":
		r.y = (r.y + 1) % r.gridSize
	case "S":
		r.y = (r.y - 1 + r.gridSize) % r.gridSize
	case "E":
		r.x = (r.x + 1) % r.gridSize
	case "W":
		r.x = (r.x - 1 + r.gridSize) % r.gridSize
	}
}

func indexOf(items []string, item string) int {
	for i, v := range items {
		if v == item {
			return i
		}
	}
	return -1
}

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) Equals(other Coordinate) bool {
	return c.X == other.X && c.Y == other.Y
}

type TelemetryPort interface {
	Record(entry string)
}

var commandTranslatorLastLanguage = "EN"

type CommandTranslator struct{}

func (t *CommandTranslator) Translate(command, language string) string {
	commandTranslatorLastLanguage = language

	switch language {
	case "EN":
		return command
	case "ES":
		return t.translateSpanish(command)
	case "FR":
		return t.translateFrench(command)
	case "PT":
		return t.translatePortuguese(command)
	case "IT":
		return t.translateItalian(command)
	default:
		return command
	}
}

func (t *CommandTranslator) TranslateSequence(commands, language string) string {
	var result strings.Builder
	for _, c := range commands {
		result.WriteString(t.Translate(string(c), language))
	}
	return result.String()
}

func (t *CommandTranslator) GetLastLanguageUsed() string {
	return commandTranslatorLastLanguage
}

func (t *CommandTranslator) translateSpanish(command string) string {
	switch command {
	case "I":
		return "L"
	case "D":
		return "R"
	case "A":
		return "M"
	default:
		return command
	}
}

func (t *CommandTranslator) translateFrench(command string) string {
	switch command {
	case "G":
		return "L"
	case "D":
		return "R"
	case "A":
		return "M"
	default:
		return command
	}
}

func (t *CommandTranslator) translatePortuguese(command string) string {
	switch command {
	case "E":
		return "L"
	case "D":
		return "R"
	case "A":
		return "M"
	default:
		return command
	}
}

func (t *CommandTranslator) translateItalian(command string) string {
	switch command {
	case "S":
		return "L"
	case "D":
		return "R"
	case "A":
		return "M"
	default:
		return command
	}
}

// Translator is the interface MissionLog depends on.
type Translator interface {
	TranslateSequence(commands, language string) string
}

type MissionLog struct {
	translator Translator
	telemetry  TelemetryPort
}

func NewMissionLog(translator Translator, telemetry TelemetryPort) *MissionLog {
	return &MissionLog{translator: translator, telemetry: telemetry}
}

func (m *MissionLog) LogTranslatedSequence(commands, language string) string {
	translated := m.translator.TranslateSequence(commands, language)
	m.telemetry.Record(fmt.Sprintf("%s:%s->%s", language, commands, translated))
	return translated
}
