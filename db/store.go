package db

import "sync"

// DB is the main in-memory storage of the cache implementation.
var DB = sync.Map{}

// KeyTTL stores the key expiration timestamp for each key
// in EPOCH Unix milliseconds format.
// A nil value means the key has no expiration (persistent).
// A key absent from this map has never been SET.
var KeyTTL = sync.Map{}
