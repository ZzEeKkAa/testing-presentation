package stable

type CounterGood struct {
	c int
}

func NewCounterGood() *CounterGood {
	return &CounterGood{
		c: -1,
	}
}

func (c *CounterGood) Inc() int {
	c.c++
	c.c++
	return c.c
}

func (c *CounterGood) Val() int {
	return c.c
}
