package cache

import "testing"

func TestLRUHeap(t *testing.T) {
	c := NewLRU(2)
	c.Set("k1", "v1")
	if v := c.Get("k1"); v != "v1" {
		t.Fatalf("got value: %s, expected: %s", v, "v1")
	}
	c.Set("k2", "v2")
	c.Set("k3", "v3")
	if count := c.Len(); count != 2 {
		t.Fatalf("got number of values: %d, expected: %d", count, 2)
	}
	if v := c.Get("k3"); v != "v3" {
		t.Fatalf("got value: %s, expected: %s", v, "v3")
	}
	if v := c.Get("k2"); v != "v2" {
		t.Fatalf("got value: %s, expected: %s", v, "v2")
	}
	c.Set("k4", "v4")
	// k3 should have been evicted
	if v := c.Get("k3"); v != "" {
		t.Fatalf("the value for k3 is still in the cache: %d", c.Len())
	}
}
