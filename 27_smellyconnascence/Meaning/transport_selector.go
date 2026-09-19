package meaning

import "fmt"

// TransportSelector's "1"/"2"/"3"/"4" only mean bike/car/train/bus by an
// unstated convention shared between the caller and this switch -- nothing
// in the code documents or enforces the mapping. Connascence of Meaning.
type TransportSelector struct {
	selected []string
}

// SetTransport records the transport named by code.
func (s *TransportSelector) SetTransport(transport string) error {
	switch transport {
	case "1":
		s.selected = append(s.selected, "bike")
	case "2":
		s.selected = append(s.selected, "car")
	case "3":
		s.selected = append(s.selected, "train")
	case "4":
		s.selected = append(s.selected, "bus")
	default:
		return fmt.Errorf("unknown transport code: %s", transport)
	}
	return nil
}
