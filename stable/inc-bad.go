package stable

type CounterBad struct {
	c int
}

func NewCounterBad() *CounterBad {
	return &CounterBad{}
}

func (c *CounterBad) Inc() int {
	c.c++
	c.c++
	return c.c
}

func (c *CounterBad) Val() int {
	return c.c
}
