package semaphore

import (
	"testing"
	"time"
)

type A struct {
	*testing.T
}

func asser(t *testing.T) *A {
	return &A{t}
}

func (a A) equalWithMessage(expected uint, actual uint, message string) {
	if actual != expected {
		a.Fatalf(message, expected, actual)
	}
}

func (a A) equal(expected uint, actual uint) {
	a.equalWithMessage(expected, actual, "Expected %d, got %d")
}

func TestSemaphore_Capacity_OfNew(t *testing.T) {
	asser(t).equal(1, NewSemaphore().Capacity())
}

func TestSemaphore_Capacity_OfNewWithArg(t *testing.T) {
	asser(t).equal(5, NewSemaphoreWithCapacity(5).Capacity())
}

func TestSemaphore_QueueLength_OfNew(t *testing.T) {
	asser(t).equal(0, NewSemaphore().QueueLength())
}

func TestSemaphore_QueueLength_WithAcquiredPermit(t *testing.T) {
	var s = NewSemaphore()

	s.Acquire()

	asser(t).equal(1, s.QueueLength())
}

func TestSemaphore_QueueLength_WithAcquiredPermitReleased(t *testing.T) {
	var s = NewSemaphoreWithCapacity(3)

	s.Acquire()
	s.Acquire()
	s.Release()

	asser(t).equal(1, s.QueueLength())
}

func TestSemaphore_Acquire_BeyondCapacity(t *testing.T) {
	var s = NewSemaphore()

	s.Acquire()

	go func() {
		s.Acquire()
		t.Error("Should not be able to acquire second permit")
	}()

	time.Sleep(100 * time.Millisecond)
}

func TestSemaphore_Acquire_WithTimeout_AcquirePermit(t *testing.T) {
	var s = NewSemaphore()

	asser(t).equal(0, s.QueueLength())

	if !s.TryAcquire(500 * time.Millisecond) {
		t.Error("Could not acquire permit from Semaphore with spare")
	}

	asser(t).equal(1, s.QueueLength())
}

func TestSemaphore_Acquire_WithTimeout_AcquireTimedout(t *testing.T) {
	var s = NewSemaphore()

	asser(t).equal(0, s.QueueLength())

	s.Acquire()

	asser(t).equal(1, s.QueueLength())

	if s.TryAcquire(500 * time.Millisecond) {
		t.Error("Acquired permit from empty Semaphore")
	}

	asser(t).equal(1, s.QueueLength())
}
