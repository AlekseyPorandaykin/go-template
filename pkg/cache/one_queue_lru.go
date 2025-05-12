package cache

import (
	lru "github.com/hashicorp/golang-lru/v2"
)

type OneQueueLru[F comparable, S comparable, T interface{}] struct {
	data    map[F]*lru.Cache[S, T]
	maxKeys int
}

func NewOneQueueLru[F comparable, S comparable, T interface{}](maxKeys int) *OneQueueLru[F, S, T] {
	return &OneQueueLru[F, S, T]{
		data:    make(map[F]*lru.Cache[S, T]),
		maxKeys: maxKeys,
	}
}

func (c *OneQueueLru[F, S, T]) Set(firstKey F, secondKey S, val T) {
	if _, has := c.data[firstKey]; !has {
		cache, _ := lru.New[S, T](c.maxKeys)
		c.data[firstKey] = cache
	}
	c.data[firstKey].Add(secondKey, val)
}

func (c *OneQueueLru[F, S, T]) Get(firstKey F, secondKey S) *T {
	if _, has := c.data[firstKey]; !has {
		return nil
	}
	val, has := c.data[firstKey].Get(secondKey)
	if !has {
		return nil
	}
	return &val
}

func (c *OneQueueLru[F, S, T]) ValuesByFirstKey(firstKey F) []T {
	if _, has := c.data[firstKey]; !has {
		return nil
	}
	return c.data[firstKey].Values()
}
