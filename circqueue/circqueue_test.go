package circqueue_test

import (
	"testing"

	"github.com/gadumitrachioaiei/algorythms/circqueue"
)

func TestQueue(t *testing.T) {
	q := circqueue.New(5)
	// pop from empty queue
	_, err := q.Pop()
	if err == nil {
		t.Fatal("Expected error")
	}
	// push and pop
	q.Push(0)
	el, err := q.Pop()
	if err != nil {
		t.Fatalf("Got error: %v, expected nothing", err)
	}
	if el != 0 {
		t.Fatalf("Got %d, expected 0", el)
	}
	// push 5 times
	for i := 0; i < 5; i++ {
		if err := q.Push(i); err != nil {
			t.Fatalf("Got error: %v, expected nothing", err)
		}
	}
	// push on full queue
	if err := q.Push(5); err == nil {
		t.Fatal("Got nothing, expected error")
	}
	// pop 5 elements
	for i := 0; i < 5; i++ {
		if el, err := q.Pop(); err != nil {
			t.Fatalf("Got error: %v, expected nothing", err)
		} else if el != i {
			t.Fatalf("Got element: %d, expected: %d", el, i)
		}
	}
	// pop empty queue
	if _, err := q.Pop(); err == nil {
		t.Fatal("Expected error")
	}
	// push and pop a lot of times
	for i := 0; i < 100; i++ {
		if err := q.Push(i); err != nil {
			t.Fatalf("Got error: %v, expected nothing", err)
		}
		if el, err := q.Pop(); err != nil {
			t.Fatalf("Got error: %v, expected nothing", err)
		} else if el != i {
			t.Fatalf("Got element: %d, expected: %d", el, i)
		}
	}
	// some pushs and pops
	// push 3 times pop two times (1 + 1)
	var countPops int
	for i := 0; i < 5*3; i++ {
		if err := q.Push(i); err != nil {
			t.Fatalf("Got error: %v, expected nothing", err)
		}
		if i%3 == 2 {
			continue
		}
		if el, err := q.Pop(); err != nil {
			t.Fatalf("Got error: %v, expected nothing", err)
		} else if el != countPops {
			t.Fatalf("Got element: %d, expected: %d", el, countPops)
		}
		countPops++
	}
	// push on full queue
	if err := q.Push(5); err == nil {
		t.Fatal("Got nothing, expected error")
	}
}
