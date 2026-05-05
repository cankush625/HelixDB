package db

import (
	"sync"
	"time"
)

// Config holds runtime configuration for HelixDB.
// All fields are safe for concurrent access via the embedded mutex.
type Config struct {
	mu sync.RWMutex

	// Hz controls how many times per second the active expiry job runs.
	// Maps to the cleanup interval: interval = 1s / Hz.
	Hz int

	// ActiveExpireEnabled controls whether the active expiry background job runs.
	ActiveExpireEnabled bool

	// MaxMemory is the maximum memory limit in bytes. 0 means unlimited.
	MaxMemory int64
}

// ServerConfig is the package-level config instance used across the server.
var ServerConfig = &Config{
	Hz:                  1,
	ActiveExpireEnabled: true,
	MaxMemory:           0,
}

// CleanupInterval returns the active expiry ticker interval derived from Hz.
func (c *Config) CleanupInterval() time.Duration {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.Hz <= 0 {
		return time.Second
	}
	return time.Duration(float64(time.Second) / float64(c.Hz))
}

// Get returns the string value of the named config parameter.
// Returns ("", false) if the parameter name is unknown.
func (c *Config) Get(param string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	switch param {
	case "hz":
		return itoa(c.Hz), true
	case "active-expire-enabled":
		return boolToYesNo(c.ActiveExpireEnabled), true
	case "maxmemory":
		return itoa64(c.MaxMemory), true
	}
	return "", false
}

// Set updates the named config parameter from a string value.
// Returns false if the parameter name is unknown or the value is invalid.
func (c *Config) Set(param, value string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	switch param {
	case "hz":
		n, ok := parseInt(value)
		if !ok || n <= 0 {
			return false
		}
		c.Hz = n
		return true
	case "active-expire-enabled":
		b, ok := yesNoToBool(value)
		if !ok {
			return false
		}
		c.ActiveExpireEnabled = b
		return true
	case "maxmemory":
		n, ok := parseInt64(value)
		if !ok || n < 0 {
			return false
		}
		c.MaxMemory = n
		return true
	}
	return false
}
