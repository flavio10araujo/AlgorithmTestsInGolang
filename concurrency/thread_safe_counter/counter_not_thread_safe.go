package main

// CounterNotThreadSafe is not thread-safe
type CounterNotThreadSafe struct {
	count int64
}

func NewCounterNotThreadSafe() *CounterNotThreadSafe {
	return &CounterNotThreadSafe{
		count: 0,
	}
}

func (c *CounterNotThreadSafe) Inc() {
	c.count++
}

func (c *CounterNotThreadSafe) Get() int64 {
	return c.count
}
