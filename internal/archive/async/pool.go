package async

import "sync"

// Pool manages a Pool of concurrent jobs on your behalf
type Pool[T any] struct {
	mu    sync.Mutex
	c     chan T
	count int
}

// NewPool constructs and returns a new Pool
func NewPool[T any]() *Pool[T] {
	c := make(chan T, 1)
	return &Pool[T]{
		c:     c,
		count: 0,
	}
}

// Add appends a new item of work to your existing Pool
//
// Add is non-blocking
func (p *Pool[T]) Add(f func() T) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.count++
	go func() { p.c <- f() }()
}

// Range ranges over all the channels using an iterator function.
// The iterator can conditionally return early by returning t, true
//
// Range is a blocking call
func (p *Pool[T]) Range(iter func(i int, t T) bool) (T, bool) {
	for i := 0; i < p.count; i++ {
		t := <-p.c
		stop := iter(i, t)
		if stop {
			return t, true
		}
	}

	var t T

	return t, false
}

// Accumulate blocks on all channels to accumulates return values into a slice.
func (p *Pool[T]) Accumulate() []T {
	var out []T
	for i := 0; i < p.count; i++ {
		out = append(out, <-p.c)
	}
	return out
}
