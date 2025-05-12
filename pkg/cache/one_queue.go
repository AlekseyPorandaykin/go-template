package cache

import (
	"sync"
)

type OneQueue[K comparable, T interface{}] struct {
	data map[K]T
	mu   sync.Mutex
}

func NewOneQueue[K comparable, T interface{}]() *OneQueue[K, T] {
	return &OneQueue[K, T]{data: make(map[K]T)}
}

func (c *OneQueue[K, T]) Set(key K, val T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = val
}

func (c *OneQueue[K, T]) Get(key K) *T {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, has := c.data[key]
	if has {
		return &val
	}
	return nil
}

func (c *OneQueue[K, T]) Values() []*T {
	c.mu.Lock()
	defer c.mu.Unlock()
	data := make([]*T, 0, len(c.data))
	for i := range c.data {
		val := c.data[i]
		data = append(data, &val)
	}

	return data
}
