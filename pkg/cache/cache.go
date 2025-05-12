package cache

import "time"

type itemStorage[T interface{}] struct {
	data T
	ttl  time.Time
}
