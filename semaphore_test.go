package semaphore

import (
  "testing"
)

func assertEqual(t *testing.T, expected uint, actual uint, message string) {
	if actual != expected {
		t.Errorf("Expected capacity %d got %d", expected, actual)
	}
}

func TestSemaphore_Capacity_OfNew(t *testing.T) {
  assertEqual(t, 1, NewSemaphore().Capacity(), "Expected capacity %d got %d")
}

func TestSemaphore_Capacity_OfNewWithArg(t *testing.T) {
	assertEqual(t, 5, NewSemaphoreWith(5).Capacity(), "Expected capacity of %d, got %d")
}

func TestSemaphore_QueueLength_OfNew(t *testing.T) {
  assertEqual(t, 0, NewSemaphore().QueueLength(), "Expected queue of %d got %d")
}

func TestSemaphore_QueueLength_WithAcquiredPermit(t *testing.T) {
  var s = NewSemaphore()

  s.Acquire()

  assertEqual(t, 1, s.QueueLength(),"Expected queue of %d got %d") 
}

func TestSemaphore_QueueLength_WithAcquiredPermitReleased(t *testing.T) {
  var s = NewSemaphoreWith(3)

  s.Acquire()
  s.Acquire()
  s.Release()

  assertEqual(t, 1, s.QueueLength(), "Expected queue of %d got %d")
}

func TestSemaphore_Acquire_WithTimeout_AcquirePermit(t *testing.T) {
  var s = NewSemaphore()

  assertEqual(t, 0, s.QueueLength(),"Expected queue of %d got %d") 

  if !s.TryAcquire() {
    t.Error("Could not acquire permit from Semaphore with spare")
  }

  assertEqual(t, 1, s.QueueLength(),"Expected queue of %d got %d") 
}

func TestSemaphore_Acquire_WithTimeout_AcquireTimedout(t *testing.T) {
  var s = NewSemaphore()

  assertEqual(t, 0, s.QueueLength(),"Expected queue of %d got %d") 

  s.Acquire()

  assertEqual(t, 1, s.QueueLength(),"Expected queue of %d got %d") 

  if s.TryAcquire() {
    t.Error("Acquired permit from empty Semaphore")
  }

  assertEqual(t, 1, s.QueueLength(),"Expected queue of %d got %d") 
}
