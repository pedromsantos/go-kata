package identity

// counter is the package-level singleton every consumer shares.
type counter struct {
	value int
}

// Increment increases the counter's value by one and returns it.
func (c *counter) Increment() int {
	c.value++
	return c.value
}

// Current returns the counter's current value.
func (c *counter) Current() int {
	return c.value
}

// GlobalCounter is the shared package-level instance. Correctness of every
// consumer depends on them all sharing this exact single instance -- there
// is no way to have two independent counters, and the dependency is
// invisible from any one consumer's own code. Connascence of Identity.
var GlobalCounter = &counter{}
