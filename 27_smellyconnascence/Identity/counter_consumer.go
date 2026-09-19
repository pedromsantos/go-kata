package identity

// CounterConsumer relies on the package-level GlobalCounter singleton.
type CounterConsumer struct{}

// RecordVisit increments the shared GlobalCounter.
func (CounterConsumer) RecordVisit() int {
	return GlobalCounter.Increment()
}

// TotalVisits returns the shared GlobalCounter's current value.
func (CounterConsumer) TotalVisits() int {
	return GlobalCounter.Current()
}
