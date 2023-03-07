package sync

import "sync"

type Counter struct{
	mutex sync.Mutex
	count int
}

func CreateCounter() *Counter{
	return &Counter{}
}

func (c *Counter) Value() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.count
}

func (c *Counter) Increment() {
	c.count++
}