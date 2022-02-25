package stadium

import (
	"sync"
)

type Counter struct {
	mu   sync.Mutex
	rank int
}

func (c *Counter) GetRank() int {
	c.mu.Lock()
	c.rank += 1
	var rank = c.rank
	c.mu.Unlock()
	return rank
}
