package db

import (
	"encoding/json"
	"strconv"
)

// ValueType represents the detected type of a stored value.
type ValueType string

const (
	TypeString  ValueType = "string"
	TypeInteger ValueType = "integer"
	TypeFloat   ValueType = "float"
	TypeBoolean ValueType = "boolean"
	TypeJSON    ValueType = "json"
)

// Value wraps a stored string with its detected type.
// Raw always holds the original string representation so GET can return it
// unchanged, regardless of the detected type.
type Value struct {
	Type ValueType
	Raw  string
}

// NewValue detects the type of raw and returns a Value.
// Detection order: integer → float → boolean → JSON → string.
func NewValue(raw string) Value {
	if _, err := strconv.ParseInt(raw, 10, 64); err == nil {
		return Value{Type: TypeInteger, Raw: raw}
	}
	if _, err := strconv.ParseFloat(raw, 64); err == nil {
		return Value{Type: TypeFloat, Raw: raw}
	}
	if _, err := strconv.ParseBool(raw); err == nil {
		return Value{Type: TypeBoolean, Raw: raw}
	}
	var js json.RawMessage
	if err := json.Unmarshal([]byte(raw), &js); err == nil {
		return Value{Type: TypeJSON, Raw: raw}
	}
	return Value{Type: TypeString, Raw: raw}
}
