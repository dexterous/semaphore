package semaphore

import "testing"

func TestSemaphore_Capacity_OfNew(t *testing.T) {
  var expected uint = 1
  var actual = NewSemaphore().Capacity()
	if actual != expected {
		t.Errorf("Expected capacity %d got %d", expected, actual)
	}
}

func TestSemaphore_Capacity_OfNewWithArg(t *testing.T) {
	var expected uint = 5
	var actual = NewSemaphoreWith(expected).Capacity()

	if actual != expected {
		t.Errorf("Expected capacity of %d, got %d", expected, actual)
	}
}

func TestSemaphore_QueueLength_OfNew(t *testing.T) {
  var expected uint = 0
  var actual = NewSemaphore().QueueLength()
  if actual != expected {
    t.Errorf("Expected queue of %d got %d", expected, actual)
  }
}

func TestSemaphore_QueueLength_WithAcquiredPermit(t *testing.T) {
  var expected uint = 1
  var s = NewSemaphore()

  s.Acquire()

  var actual = s.QueueLength()

  if actual != expected {
    t.Errorf("Expected queue of %d got %d", expected, actual)
  }
}

func TestSemaphore_QueueLength_WithAcquiredPermitReleased(t *testing.T) {
  var expected uint = 1
  var s = NewSemaphoreWith(3)

  s.Acquire()
  s.Acquire()
  s.Release()

  var actual = s.QueueLength()

  if actual != expected {
    t.Errorf("Expected queue of %d got %d", expected, actual)
  }
}
