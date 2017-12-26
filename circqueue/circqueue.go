package circqueue

import (
	"github.com/pkg/errors"
)

// Queue is a circular queue of fixed length
// s is the oldest element in the queue or -1 if there are no elements
// t is the newest element in the queue or -1 if there are no elements
// a is a slice of fixed length that will hold the elements of the queue
// but those elements can only be read considering s and t
// Internal implementation: we will always have t greater or equal to s,
// considering they are on a circle with the pozitive sens
// from west ( 0 ) to est ( capacity ).
// s can never pass by t ( when they are equal on pop, they will be resetted )
// t is always at least s, and t cannot become equal to s on push, as that will mean full queue
// Can we make it variable size, meaning that when we make a push on full queue, to increase the queue ?
// Can we also decrease the size, if we find we have too many uoccupied places ?
type Queue struct {
	s, t int
	a    []int
}

func New(cap int) *Queue {
	return &Queue{
		s: -1,
		t: -1,
		a: make([]int, cap),
	}
}

func (q *Queue) Pop() (int, error) {
	if q.s == -1 {
		return 0, errors.New("no more elements")
	}
	el := q.a[q.s]
	if q.s == q.t {
		// no more elements, reset
		q.s, q.t = -1, -1
	} else {
		// we set the oldest element to be the next one,
		// which exists because there is at least one more element
		q.s = (q.s + 1) % cap(q.a)
	}
	return el, nil
}

func (q *Queue) Push(el int) error {
	// if there are no elements
	if q.t == -1 {
		q.a[0] = el
		q.s, q.t = 0, 0
		return nil
	}
	// we set the next element,
	// unless the next element is the oldest element
	// which would mean that the queue is full
	newIndex := (q.t + 1) % cap(q.a)
	if newIndex == q.s {
		return errors.New("we are full")
	}
	q.a[newIndex] = el
	q.t = newIndex
	return nil

}
