package db

import (
	"testing"
	"time"
)

func TestCleanExpiredKeys(t *testing.T) {
	// Clear state before each test run
	DB.Clear()
	KeyTTL.Clear()

	now := time.Now().UnixMilli()

	// Expired keys — TTL in the past
	for _, key := range []string{"exp1", "exp2", "exp3"} {
		DB.Store(key, "value")
		KeyTTL.Store(key, now-1000)
	}

	// Persistent key — nil TTL
	DB.Store("persistent", "value")
	KeyTTL.Store("persistent", nil)

	// Active key — TTL in the future
	DB.Store("active", "value")
	KeyTTL.Store("active", now+60_000)

	total, expired := cleanExpiredKeys()

	if expired != 3 {
		t.Errorf("expected 3 expired keys, got %d", expired)
	}
	if total != 5 {
		t.Errorf("expected 5 keys scanned, got %d", total)
	}

	// Expired keys must be gone from both maps
	for _, key := range []string{"exp1", "exp2", "exp3"} {
		if _, ok := DB.Load(key); ok {
			t.Errorf("expected %q to be deleted from DB", key)
		}
		if _, ok := KeyTTL.Load(key); ok {
			t.Errorf("expected %q to be deleted from KeyTTL", key)
		}
	}

	// Persistent and active keys must still exist
	for _, key := range []string{"persistent", "active"} {
		if _, ok := DB.Load(key); !ok {
			t.Errorf("expected %q to still be in DB", key)
		}
	}
}

func TestCleanExpiredKeys_ChunkLimit(t *testing.T) {
	DB.Clear()
	KeyTTL.Clear()

	now := time.Now().UnixMilli()

	// Insert more keys than chunkSize, all expired
	total := chunkSize + 5
	for i := 0; i < total; i++ {
		key := string(rune('a' + i))
		DB.Store(key, "value")
		KeyTTL.Store(key, now-1000)
	}

	scanned, _ := cleanExpiredKeys()

	if scanned > chunkSize {
		t.Errorf("cleanExpiredKeys scanned %d keys, expected at most %d", scanned, chunkSize)
	}
}
