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

func (a A) equalWithMessage(actual, expected uint, message string) {
	if actual != expected {
		a.Fatalf(message, actual, expected)
	}
}

func (a A) equal(actual, expected uint) {
	a.equalWithMessage(actual, expected, "Expected %d, got %d")
}

func TestSemaphore_Capacity_OfNew(t *testing.T) {
	asser(t).equal(NewSemaphore().Capacity(), 1)
}

func TestSemaphore_Capacity_OfNewWithArg(t *testing.T) {
	asser(t).equal(NewSemaphoreWithCapacity(5).Capacity(), 5)
}

func TestSemaphore_QueueLength_OfNew(t *testing.T) {
	asser(t).equal(NewSemaphore().QueueLength(), 0)
}

func TestSemaphore_QueueLength_WithAcquiredPermit(t *testing.T) {
	var s = NewSemaphore()

	s.Acquire()

	asser(t).equal(s.QueueLength(), 1)
}

func TestSemaphore_QueueLength_WithAcquiredPermitReleased(t *testing.T) {
	var s = NewSemaphoreWithCapacity(3)

	s.Acquire()
	s.Acquire()
	s.Release()

	asser(t).equal(s.QueueLength(), 1)
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

	asser(t).equal(s.QueueLength(), 0)

	if !s.TryAcquire(500 * time.Millisecond) {
		t.Error("Could not acquire permit from Semaphore with spare")
	}

	asser(t).equal(s.QueueLength(), 1)
}

func TestSemaphore_Acquire_WithTimeout_AcquireTimedout(t *testing.T) {
	var s = NewSemaphore()

	asser(t).equal(s.QueueLength(), 0)

	s.Acquire()

	asser(t).equal(s.QueueLength(), 1)

	if s.TryAcquire(500 * time.Millisecond) {
		t.Error("Acquired permit from empty Semaphore")
	}

	asser(t).equal(s.QueueLength(), 1)
}
