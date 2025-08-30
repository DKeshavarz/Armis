package queue

import (
	"errors"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	q := New[int]()
	if !q.IsEmpty() {
		t.Errorf("expected new queue to be empty, but IsEmpty() = false")
	}
	if q.Len() != 0 {
		t.Errorf("expected Len() = 0, got %d", q.Len())
	}
}

func TestEnqueueDequeue(t *testing.T) {
	q := New[string]()

	// Enqueue elements
	err := q.Enqueue("a")
	if err != nil {
		t.Errorf("unexpected error on Enqueue: %v", err)
	}
	if q.Len() != 1 {
		t.Errorf("expected Len() = 1, got %d", q.Len())
	}

	err = q.Enqueue("b")
	if err != nil {
		t.Errorf("unexpected error on Enqueue: %v", err)
	}
	if q.Len() != 2 {
		t.Errorf("expected Len() = 2, got %d", q.Len())
	}

	// Dequeue
	val, err := q.Dequeue()
	if err != nil {
		t.Errorf("unexpected error on Dequeue: %v", err)
	}
	if val != "a" {
		t.Errorf("expected Dequeue 'a', got %q", val)
	}
	if q.Len() != 1 {
		t.Errorf("expected Len() = 1 after Dequeue, got %d", q.Len())
	}

	val, err = q.Dequeue()
	if err != nil {
		t.Errorf("unexpected error on Dequeue: %v", err)
	}
	if val != "b" {
		t.Errorf("expected Dequeue 'b', got %q", val)
	}
	if q.IsEmpty() != true {
		t.Errorf("expected IsEmpty() = true after Dequeue, got false")
	}

	// Dequeue on empty
	_, err = q.Dequeue()
	if !errors.Is(err, ErrEmpty) {
		t.Errorf("expected ErrEmpty on empty Dequeue, got %v", err)
	}
}

func TestPeek(t *testing.T) {
	q := New[int]()

	// Peek on empty
	_, err := q.Peek()
	if !errors.Is(err, ErrEmpty) {
		t.Errorf("expected ErrEmpty on empty Peek, got %v", err)
	}

	// Enqueue and Peek
	q.Enqueue(42)
	val, err := q.Peek()
	if err != nil {
		t.Errorf("unexpected error on Peek: %v", err)
	}
	if val != 42 {
		t.Errorf("expected Peek 42, got %d", val)
	}
	if q.Len() != 1 {
		t.Errorf("expected Len() unchanged after Peek, got %d", q.Len())
	}
}

func TestPopWhile(t *testing.T) {
	q := New[int]()

	// PopWhile on empty
	result := q.PopWhile(func(n int) bool { return n < 5 })
	if result != nil {
		t.Errorf("expected nil on empty PopWhile, got %v", result)
	}

	// Enqueue elements
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)

	// PopWhile with predicate matching prefix
	result = q.PopWhile(func(n int) bool { return n < 4 })
	expected := []int{1, 2, 3}
	if len(result) != len(expected) {
		t.Errorf("expected PopWhile result len %d, got %d", len(expected), len(result))
	}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("expected PopWhile result[%d] = %d, got %d", i, expected[i], result[i])
		}
	}
	if q.Len() != 3 {
		t.Errorf("expected Len() = 3 after PopWhile, got %d", q.Len())
	}

	// PopWhile with no matches
	result = q.PopWhile(func(n int) bool { return n < 4 })
	if len(result) != 0 {
		t.Errorf("expected empty result for no matches, got %v", result)
	}
	if q.Len() != 3 {
		t.Errorf("expected Len() unchanged, got %d", q.Len())
	}

	// PopWhile matching all remaining
	result = q.PopWhile(func(n int) bool { return true })
	expected = []int{4, 5, 6}
	if len(result) != len(expected) {
		t.Errorf("expected PopWhile result len %d, got %d", len(expected), len(result))
	}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("expected PopWhile result[%d] = %d, got %d", i, expected[i], result[i])
		}
	}
	if !q.IsEmpty() {
		t.Errorf("expected empty queue after popping all")
	}
}

func TestConcurrentOperations(t *testing.T) {
	q := New[int]()
	var wg sync.WaitGroup
	numGoroutines := 100
	numOpsPerGoroutine := 100

	// Concurrent Enqueue
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			for j := 0; j < numOpsPerGoroutine; j++ {
				q.Enqueue(start + j)
			}
		}(i * numOpsPerGoroutine)
	}
	wg.Wait()

	expectedLen := numGoroutines * numOpsPerGoroutine
	if q.Len() != expectedLen {
		t.Errorf("expected Len() = %d after concurrent Enqueue, got %d", expectedLen, q.Len())
	}

	// Concurrent Dequeue
	dequeued := make(chan int, expectedLen)
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				val, err := q.Dequeue()
				if errors.Is(err, ErrEmpty) {
					return
				}
				if err != nil {
					t.Errorf("unexpected error in concurrent Dequeue: %v", err)
					return
				}
				dequeued <- val
			}
		}()
	}
	wg.Wait()
	close(dequeued)

	count := 0
	for range dequeued {
		count++
	}
	if count != expectedLen {
		t.Errorf("expected %d elements dequeued concurrently, got %d", expectedLen, count)
	}
	if !q.IsEmpty() {
		t.Errorf("expected empty queue after concurrent Dequeue")
	}
}