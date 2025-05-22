package dispatcher

import "time"

type Event[T interface{}] struct {
	Name string
	Body T
}

type Listener[T interface{}] interface {
	Handle(e Event[T])
}

type Producer[T interface{}] interface {
	SetPreDispatcher(handler func(e Event[T]))
	SetPostDispatcher(handler func(e Event[T]))
	Dispatch(e Event[T])
}

type Consumer[T interface{}] interface {
	SetPreHandler(handler func(e Event[T]))
	SetPostHandler(handler func(e Event[T], duration time.Duration))
	Subscribe(consumer Listener[T])
	Listen()
	Close()
}

type Dispatcher[T interface{}] interface {
	Producer[T]
	Consumer[T]
}
