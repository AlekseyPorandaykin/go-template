package cache

import (
	"sync"
	"time"
)

type TwoQueue[F comparable, S comparable, T interface{}] struct {
	data map[F]map[S]itemStorage[T]
	mu   sync.Mutex
}

func NewTwoQueueCache[F comparable, S comparable, T interface{}]() *TwoQueue[F, S, T] {
	return &TwoQueue[F, S, T]{
		data: make(map[F]map[S]itemStorage[T]),
	}
}

func (c *TwoQueue[F, S, T]) Set(firstKey F, secondKey S, val T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.data[firstKey] == nil {
		c.data[firstKey] = make(map[S]itemStorage[T])
	}
	c.data[firstKey][secondKey] = itemStorage[T]{data: val}
}
func (c *TwoQueue[F, S, T]) SetWithTTL(firstKey F, secondKey S, val T, ttl time.Duration) {
	now := time.Now().In(time.UTC)
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.data[firstKey] == nil {
		c.data[firstKey] = make(map[S]itemStorage[T])
	}
	c.data[firstKey][secondKey] = itemStorage[T]{data: val, ttl: now.Add(ttl)}
}

func (c *TwoQueue[F, S, T]) Get(firstKey F, secondKey S) *T {
	now := time.Now().In(time.UTC)
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.data[firstKey] == nil {
		return nil
	}
	data, has := c.data[firstKey][secondKey]
	if !has {
		return nil
	}
	if !data.ttl.IsZero() && data.ttl.Before(now) {
		delete(c.data[firstKey], secondKey)
		return nil
	}
	return &data.data
}

func (c *TwoQueue[F, S, T]) GetAndDelete(firstKey F, secondKey S) *T {
	val := c.Get(firstKey, secondKey)
	delete(c.data[firstKey], secondKey)
	return val
}

func (c *TwoQueue[F, S, T]) ValuesByFirstKey(firstKey F) []T {
	now := time.Now().In(time.UTC)
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.data[firstKey] == nil || len(c.data[firstKey]) == 0 {
		return nil
	}
	result := make([]T, 0, len(c.data[firstKey]))
	for _, item := range c.data[firstKey] {
		if !item.ttl.IsZero() && item.ttl.Before(now) {
			continue
		}
		result = append(result, item.data)
	}
	return result
}

func (c *TwoQueue[F, S, T]) ValuesBySecondKey(secondKey S) []T {
	now := time.Now().In(time.UTC)
	c.mu.Lock()
	defer c.mu.Unlock()
	result := make([]T, 0, 1_000)
	for _, values := range c.data {
		if val, has := values[secondKey]; has {
			if !val.ttl.IsZero() && val.ttl.Before(now) {
				continue
			}
			result = append(result, val.data)
		}
	}
	return result
}

func (c *TwoQueue[F, S, T]) Values() []T {
	now := time.Now().In(time.UTC)
	c.mu.Lock()
	defer c.mu.Unlock()
	result := make([]T, 0, 1_000)
	for _, values := range c.data {
		for _, value := range values {
			if !value.ttl.IsZero() && value.ttl.Before(now) {
				continue
			}
			result = append(result, value.data)
		}
	}
	return result
}

func (c *TwoQueue[F, S, T]) ValuesByCondition(condition func(item T) bool) []T {
	now := time.Now().In(time.UTC)
	c.mu.Lock()
	defer c.mu.Unlock()
	result := make([]T, 0, 1_000)
	for _, values := range c.data {
		for _, value := range values {
			if !value.ttl.IsZero() && value.ttl.Before(now) {
				continue
			}
			if condition(value.data) {
				result = append(result, value.data)
			}
		}
	}
	return result
}
func (c *TwoQueue[F, S, T]) ListByCondition(firstKey F, secondKey S, condition func(item T) bool) []T {
	now := time.Now().In(time.UTC)
	c.mu.Lock()
	defer c.mu.Unlock()
	result := make([]T, 0, 1_000)
	for _, values := range c.data {
		for _, value := range values {
			if !value.ttl.IsZero() && value.ttl.Before(now) {
				continue
			}
			if condition(value.data) {
				result = append(result, value.data)
			}
		}
	}
	return result
}

func (c *TwoQueue[F, S, T]) DeleteByCondition(condition func(item T) bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for firstKey, values := range c.data {
		for secondKey, value := range values {
			if condition(value.data) {
				delete(c.data[firstKey], secondKey)
			}
		}
	}
}
