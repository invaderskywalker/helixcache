package cache

import (
	"sync"
	"testing"
)

func TestCacheSetGetDelete(t *testing.T) {
	cache := Init()
	key := "test-key"
	val := []byte("test-value")

	// Test Set
	if err := cache.Set(key, val); err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	// Test Get
	got, ok := cache.Get(key)
	if !ok {
		t.Fatalf("Get failed: key not found")
	}
	if string(got) != string(val) {
		t.Errorf("expected %q, got %q", val, got)
	}

	// Test Delete
	if err := cache.Delete(key); err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	// Test Get after Delete (should return ok==false)
	got, ok = cache.Get(key)
	if ok {
		t.Errorf("expected key to be deleted, but found value: %v", got)
	}
}

func TestCacheConcurrentAccess(t *testing.T) {
	cache := Init()
	key := "concurrent-key"
	valBase := byte('A')
	WG := sync.WaitGroup{}
	setCount := 100

	// Multiple sets concurrently
	for i := 0; i < setCount; i++ {
		WG.Add(1)
		go func(i int) {
			defer WG.Done()
			cache.Set(key, []byte{valBase + byte(i%26)})
		}(i)
	}
	WG.Wait()

	// Get after concurrent sets
	got, ok := cache.Get(key)
	if !ok {
		t.Error("Value missing after concurrent Set")
	} else if got == nil {
		t.Error("Get returned nil after concurrent Set, want value")
	}
}

func TestCacheGetMissingKey(t *testing.T) {
	cache := Init()
	val, ok := cache.Get("does-not-exist")
	if ok {
		t.Errorf("Expected missing key to return ok==false, got ok==true and val=%v", val)
	}
	if val != nil && len(val) > 0 {
		t.Errorf("Expected nil/empty for missing key, got: %v", val)
	}
}
