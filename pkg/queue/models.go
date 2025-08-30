package queue

import "errors"

type QueueIface[T any] interface {
	Enqueue(elem T) error
	Dequeue() (T, error)
	Peek() (T, error)
	Len() int
	IsEmpty() bool
	PopWhile(pred func(T) bool) []T
}

var (
	ErrEmpty = errors.New("queue is empty")
)
