package semaphore

type permit struct{}
type Semaphore chan permit

var aPermit permit

func NewSemaphore() Semaphore {
	return NewSemaphoreWith(1)
}

func NewSemaphoreWith(capacity uint) Semaphore {
	return Semaphore(make(chan permit, capacity))
}

func (s Semaphore) Capacity() uint {
	return uint(cap(s))
}

func (s Semaphore) QueueLength() uint {
  return uint(len(s))
}

func (s Semaphore) Acquire() {
  s <- aPermit
}

func (s Semaphore) Release() {
  <- s
}
