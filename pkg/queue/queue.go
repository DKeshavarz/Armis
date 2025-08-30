package queue

import "sync"

type queue[T any] struct {
	mu     sync.Mutex
	q      []T
}

func New[T any]() QueueIface[T] {
	impl := &queue[T]{
		q: make([]T, 0),
	}
	return impl
}

func (s *queue[T]) Enqueue(elem T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.q = append(s.q, elem)
	return nil
}

func (s *queue[T]) Peek() (T, error) {
	var zero T
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.q) == 0 {
		return zero, ErrEmpty
	}
	return s.q[0], nil
}

func (s *queue[T]) Dequeue() (T, error) {
	var zero T
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.q) == 0 {
		return zero, ErrEmpty
	}
	head := s.q[0]
	// zero slot for GC
	s.q[0] = zero
	s.q = s.q[1:]
	return head, nil
}
func (s *queue[T]) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.q)
}

func (s *queue[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *queue[T]) PopWhile(pred func(elem T) bool) []T {
	nowZero := make([]T, 0)
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.q) == 0 {
		return nil
	}
	i := 0
	for i < len(s.q) {
		if pred(s.q[i]) {
			nowZero = append(nowZero, s.q[i])
			i++
			continue
		}
		break
	}
	if i > 0 {
		// zero out for GC safety
		var zero T
		for j := 0; j < i; j++ {
			s.q[j] = zero
		}
		s.q = s.q[i:]
	}
	return nowZero
}
