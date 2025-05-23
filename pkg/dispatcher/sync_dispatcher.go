package dispatcher

import "time"

type SyncProducer[T interface{}] struct {
	listeners      []Listener[T]
	preHandler     func(e Event[T])
	postHandler    func(e Event[T], duration time.Duration)
	preDispatcher  func(e Event[T])
	postDispatcher func(e Event[T])
	closeCh        chan struct{}
	events         chan Event[T]
}

func NewSyncProducer[T interface{}]() Dispatcher[T] {
	return &SyncProducer[T]{
		listeners: make([]Listener[T], 0),
		closeCh:   make(chan struct{}, 1),
		events:    make(chan Event[T], 100),
	}
}

func (c *SyncProducer[T]) SetPreHandler(handler func(e Event[T])) {
	c.preHandler = handler
}

func (c *SyncProducer[T]) SetPostHandler(handler func(e Event[T], duration time.Duration)) {
	c.postHandler = handler
}

func (c *SyncProducer[T]) SetPreDispatcher(handler func(e Event[T])) {
	c.preDispatcher = handler
}
func (c *SyncProducer[T]) SetPostDispatcher(handler func(e Event[T])) {
	c.postDispatcher = handler
}

func (c *SyncProducer[T]) Dispatch(e Event[T]) {
	if c.preDispatcher != nil {
		c.preDispatcher(e)
	}
	for _, consumer := range c.listeners {
		now := time.Now()
		if c.preHandler != nil {
			c.preHandler(e)
		}
		consumer.Handle(e)
		if c.postHandler != nil {
			c.postHandler(e, time.Now().Sub(now))
		}
	}
	if c.postDispatcher != nil {
		c.postDispatcher(e)
	}
}

func (c *SyncProducer[T]) Subscribe(consumer Listener[T]) {
	c.listeners = append(c.listeners, consumer)
}

func (c *SyncProducer[T]) Listen() {

}

func (c *SyncProducer[T]) Close() {
	c.closeCh <- struct{}{}
}
