package pokecache_test

import (
	"testing"
	"time"

	"github.com/BambiCPT/pokedexcli/internal/pokecache"
)

func TestCache_Add(t *testing.T) {
	expiry := 2 * time.Second
	cache := pokecache.NewCache(expiry)

	key := "testKey"
	value := []byte("testValue")

	cache.Add(key, value)
	got, ok := cache.Get(key)
	if !ok {
		t.Fatalf("expected key %q to exist in cache", key)
	}
	if string(got) != string(value) {
		t.Errorf("expected %q, got %q", string(value), string(got))
	}
}

func TestCache_Expiration(t *testing.T) {
	expiry := 1 * time.Second
	cache := pokecache.NewCache(expiry)

	key := "expiringKey"
	value := []byte("tempValue")

	cache.Add(key, value)

	time.Sleep(2 * time.Second)

	_, ok := cache.Get(key)
	if ok {
		t.Errorf("expected key %q to be expired and removed from cache", key)
	}
}

func TestCache_NonExistentKey(t *testing.T) {
	expiry := 2 * time.Second
	cache := pokecache.NewCache(expiry)

	_, ok := cache.Get("nonExistent")
	if ok {
		t.Errorf("expected non-existent key to return false, but got true")
	}
}
