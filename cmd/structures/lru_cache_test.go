package structures_test

import (
	"testing"

	"skabillium.io/kata-go/cmd/structures"
)

func TestLruCache(t *testing.T) {
	lru := structures.NewLruCache(3)
	if lru.Get("foo") != nil {
		t.Error("Expected Get('foo') to return nil")
	}

	lru.Update("foo", 69)
	if lru.Get("foo") != 69 {
		t.Error("Expected Get('foo') to return 69")
	}

	lru.Update("bar", 420)
	if lru.Get("bar") != 420 {
		t.Error("Expected Get('bar') to return 420")
	}

	lru.Update("baz", 1337)
	if lru.Get("baz") != 1337 {
		t.Error("Expected Get('baz') to return 1337")
	}

	lru.Update("ball", 69420)
	if lru.Get("ball") != 69420 {
		t.Error("Expected Get('ball') to return 1337")
	}
	if lru.Get("foo") != nil {
		t.Error("Expected Get('foo') to return nil")
	}
	if lru.Get("bar") != 420 {
		t.Error("Expected Get('bar') to return 420")
	}

	lru.Update("foo", 69)
	if lru.Get("bar") != 420 {
		t.Error("Expected Get('bar') to return 420")
	}
	if lru.Get("foo") != 69 {
		t.Error("Expected Get('foo') to return 69")
	}
}
