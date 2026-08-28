package counter

import (
	"sync/atomic"
)

type Counter struct {
	count atomic.Int64
}

func (c *Counter) Inc() {
	c.count.Add(1)
}
