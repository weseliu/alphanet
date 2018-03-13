package net

const DefaultQueueSize = 100

type EventQueue interface {
	StartLoop()
	StopLoop(result int)

	Wait() int
	Post(callback func())
}

type eventQueue struct {
	queue chan func()
	exitSignal chan int
	capturePanic bool
}

func (Self *eventQueue) StartLoop() {
	go func() {
		for callback := range Self.queue {
			Self.protectedCall(callback)
		}
	}()
}

func (Self *eventQueue) StopLoop(result int) {
	Self.exitSignal <- result
}

func (Self *eventQueue) Wait() int {
	return <-Self.exitSignal
}

func (Self *eventQueue) Post(callback func()) {
	if callback == nil {
		return
	}

	Self.queue <- callback
}

func (Self *eventQueue) protectedCall(callback func()) {
	if callback == nil {
		return
	}

	callback()
}

func NewEventQueue() EventQueue {
	return NewEventQueueByLen(DefaultQueueSize)
}

func NewEventQueueByLen(l int) EventQueue {
	self := &eventQueue{
		queue:      make(chan func(), l),
		exitSignal: make(chan int),
	}
	return self
}
